package projectRepository

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"sixTask/helpers/conversionTypes"
	"sixTask/internal/http/request/projectRequest"

	"sixTask/internal/database"
)

// PaginationResult contém os projetos paginados e os metadados de paginação
type PaginationResult struct {
	Data []database.Project `json:"data"`
	Meta PaginationMeta     `json:"meta"`
}

// PaginationMeta contém os metadados de paginação
type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
	TotalPages  int   `json:"total_pages"`
}

// GetProjectsWithPagination retorna os projetos paginados e os metadados de paginação
func GetProjectsWithPagination(ctx context.Context, page, limit int) (PaginationResult, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)

	// Validação dos parâmetros
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Cálculo do offset
	offset := (page - 1) * limit

	// Buscar o total de projetos para metadados de paginação
	total, err := queries.CountProjects(ctx)
	if err != nil {
		return PaginationResult{}, err
	}

	// Calcular o total de páginas
	totalPages := (int(total) + limit - 1) / limit

	// Buscar projetos com paginação via SQL
	paginatedProjects, err := queries.FindManyProjectsWithPagination(ctx, database.FindManyProjectsWithPaginationParams{
		Offset: int32(offset),
		Limit:  int32(limit),
	})
	if err != nil {
		return PaginationResult{}, err
	}

	// Montar resultado com metadados de paginação
	result := PaginationResult{
		Data: paginatedProjects,
		Meta: PaginationMeta{
			CurrentPage: page,
			PerPage:     limit,
			Total:       total,
			TotalPages:  totalPages,
		},
	}

	return result, nil
}

// GetProjects retorna todos os projetos com paginação (mantido para compatibilidade)
func GetProjects(ctx context.Context, offset, limit int32) ([]database.Project, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.FindManyProjectsWithPagination(ctx, database.FindManyProjectsWithPaginationParams{
		Offset: offset,
		Limit:  limit,
	})
}

// CountProjects retorna o total de projetos (mantido para compatibilidade)
func CountProjects(ctx context.Context) (int64, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.CountProjects(ctx)
}

// GetProject retorna um projeto pelo ID
func GetProject(ctx context.Context, id int64) (database.Project, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.FindProjectById(ctx, id)
}

// CreateProject cria um novo projeto
func CreateProject(request projectRequest.CreateProjectRequest) (database.Project, []database.User, error) {
	params := request.ToCreateProjectParams().(database.CreateProjectParams)
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	project, err := queries.CreateProject(ctx, params)

	if err != nil {
		return database.Project{}, []database.User{}, err
	}

	for _, user_id := range request.UsersId {

		queries.CreateUserProject(ctx, database.CreateUserProjectParams{
			UserID: user_id,
			ProjectID: pgtype.Int8{
				Int64: project.ID,
				Valid: true,
			},
		})

	}

	usersList := conversionTypes.ConvertPgInt8Slice(request.UsersId)

	users, err := queries.FindManyUserIds(ctx, usersList)

	return project, users, nil
}

// UpdateProject atualiza um projeto existente
func UpdateProject(ctx context.Context, params database.UpdateProjectParams) (database.Project, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.UpdateProject(ctx, params)
}

// UpdateProjectWithUsers atualiza um projeto existente e suas relações com usuários
func UpdateProjectWithUsers(request projectRequest.UpdateProjectRequest, id int64) (database.Project, []database.User, error) {
	// Converter a request para os parâmetros do projeto
	params := request.ToUpdateProjectParams(id).(database.UpdateProjectParams)

	// Conectar ao banco de dados
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)

	// Atualizar o projeto
	project, err := queries.UpdateProject(ctx, params)
	if err != nil {
		return database.Project{}, []database.User{}, err
	}

	// Excluir todas as relações existentes entre usuários e o projeto
	err = queries.DeleteUserProjectsByProjectId(ctx, pgtype.Int8{
		Int64: id,
		Valid: true,
	})
	if err != nil {
		return database.Project{}, []database.User{}, err
	}

	// Criar novas relações entre usuários e o projeto
	for _, user_id := range request.UsersId {
		err = queries.CreateUserProject(ctx, database.CreateUserProjectParams{
			UserID: user_id,
			ProjectID: pgtype.Int8{
				Int64: id,
				Valid: true,
			},
		})
		if err != nil {
			return database.Project{}, []database.User{}, err
		}
	}

	// Buscar os usuários para retornar junto com o projeto
	usersList := conversionTypes.ConvertPgInt8Slice(request.UsersId)
	users, err := queries.FindManyUserIds(ctx, usersList)
	if err != nil {
		return database.Project{}, []database.User{}, err
	}

	return project, users, nil
}

// DeleteProject remove um projeto
func DeleteProject(ctx context.Context, id int64) error {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.DeleteProject(ctx, id)
}

// GetProjectsWithUsersAndPagination retorna os projetos com usuários e paginação
func GetProjectsWithUsersAndPagination(page, limit int) ([]database.FindManyProjectsWithUsersWithPaginationRow, int64, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)

	// Validação dos parâmetros
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Cálculo do offset
	offset := (page - 1) * limit

	// Buscar o total de projetos para metadados de paginação
	total, err := queries.CountProjectsWithUsers(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Buscar projetos com usuários e paginação via SQL
	projects, err := queries.FindManyProjectsWithUsersWithPagination(ctx, database.FindManyProjectsWithUsersWithPaginationParams{
		Offset: int32(offset),
		Limit:  int32(limit),
	})
	if err != nil {
		return nil, 0, err
	}

	return projects, total, nil
}

// FindManyProjectsWithUsers retorna todos os projetos com usuários
func FindManyProjectsWithUsers() ([]database.FindManyProjectsWithUsersRow, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.FindManyProjectsWithUsers(ctx)
}

// GetProjectsByClientIdAndPagination retorna projetos pelo ID do cliente com paginação
func GetProjectsByClientIdAndPagination(clientId pgtype.Int8, page, limit int) ([]database.FindManyProjectsClientWithUsersWithPaginationRow, int64, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)

	// Validação dos parâmetros
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Cálculo do offset
	offset := (page - 1) * limit

	// Buscar o total de projetos para metadados de paginação
	total, err := queries.CountProjectsByClientId(ctx, clientId)
	if err != nil {
		return nil, 0, err
	}

	// Buscar projetos com paginação via SQL
	projects, err := queries.FindManyProjectsClientWithUsersWithPagination(ctx, database.FindManyProjectsClientWithUsersWithPaginationParams{
		ClientID: clientId,
		Offset:   int32(offset),
		Limit:    int32(limit),
	})
	if err != nil {
		return nil, 0, err
	}

	return projects, total, nil
}

// GetProjectsByUserIdAndPagination retorna projetos pelo ID do usuário com paginação
func GetProjectsByUserIdAndPagination(userId pgtype.Int8, page, limit int) ([]database.FindManyProjectsUserWithUsersWithPaginationRow, int64, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)

	// Validação dos parâmetros
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Cálculo do offset
	offset := (page - 1) * limit

	// Buscar o total de projetos para metadados de paginação
	total, err := queries.CountProjectsByUserId(ctx, userId)
	if err != nil {
		return nil, 0, err
	}

	// Buscar projetos com paginação via SQL
	projects, err := queries.FindManyProjectsUserWithUsersWithPagination(ctx, database.FindManyProjectsUserWithUsersWithPaginationParams{
		UserID: userId,
		Offset: int32(offset),
		Limit:  int32(limit),
	})
	if err != nil {
		return nil, 0, err
	}

	return projects, total, nil
}
