package userRepository

import (
	"context"

	"sixTask/internal/database"
)

// PaginationResult contém os usuários paginados e os metadados de paginação
type PaginationResult struct {
	Data []database.User `json:"data"`
	Meta PaginationMeta  `json:"meta"`
}

// PaginationMeta contém os metadados de paginação
type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
	TotalPages  int   `json:"total_pages"`
}

// GetUsersWithPagination retorna os usuários paginados e os metadados de paginação
func GetUsersWithPagination(ctx context.Context, page, limit int) (PaginationResult, error) {
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

	// Buscar o total de usuários para metadados de paginação
	total, err := queries.CountUsers(ctx)
	if err != nil {
		return PaginationResult{}, err
	}

	// Calcular o total de páginas
	totalPages := (int(total) + limit - 1) / limit

	// Buscar usuários com paginação via SQL
	paginatedUsers, err := queries.FindManyWithPagination(ctx, database.FindManyWithPaginationParams{
		Offset: int32(offset),
		Limit:  int32(limit),
	})
	if err != nil {
		return PaginationResult{}, err
	}

	// Montar resultado com metadados de paginação
	result := PaginationResult{
		Data: paginatedUsers,
		Meta: PaginationMeta{
			CurrentPage: page,
			PerPage:     limit,
			Total:       total,
			TotalPages:  totalPages,
		},
	}

	return result, nil
}

// GetUsers retorna todos os usuários com paginação (mantido para compatibilidade)
func GetUsers(ctx context.Context, offset, limit int32) ([]database.User, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.FindManyWithPagination(ctx, database.FindManyWithPaginationParams{
		Offset: offset,
		Limit:  limit,
	})
}

// CountUsers retorna o total de usuários (mantido para compatibilidade)
func CountUsers(ctx context.Context) (int64, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.CountUsers(ctx)
}

// GetUser retorna um usuário pelo ID
func GetUser(ctx context.Context, id int64) (database.User, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.FindById(ctx, id)
}

// GetUserByEmail retorna um usuário pelo email
func GetUserByEmail(ctx context.Context, email string) (database.User, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.FindByEmail(ctx, email)
}

// CreateUser cria um novo usuário
func CreateUser(ctx context.Context, params database.CreateUserParams) (database.User, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.CreateUser(ctx, params)
}

// UpdateUser atualiza um usuário existente
func UpdateUser(ctx context.Context, params database.UpdateUserParams) (database.User, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.UpdateUser(ctx, params)
}

// DeleteUser remove um usuário
func DeleteUser(ctx context.Context, id int64) error {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.DeleteUser(ctx, id)
}
