package projectRepository

import (
	"context"

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
func CreateProject(ctx context.Context, params database.CreateProjectParams) (database.Project, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.CreateProject(ctx, params)
}

// UpdateProject atualiza um projeto existente
func UpdateProject(ctx context.Context, params database.UpdateProjectParams) (database.Project, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.UpdateProject(ctx, params)
}

// DeleteProject remove um projeto
func DeleteProject(ctx context.Context, id int64) error {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.DeleteProject(ctx, id)
}
