package clientRepository

import (
	"context"

	"sixTask/internal/database"
)

// PaginationResult contém os clientes paginados e os metadados de paginação
type PaginationResult struct {
	Data []database.Client `json:"data"`
	Meta PaginationMeta    `json:"meta"`
}

// PaginationMeta contém os metadados de paginação
type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
	TotalPages  int   `json:"total_pages"`
}

// GetClientsWithPagination retorna os clientes paginados e os metadados de paginação
func GetClientsWithPagination(ctx context.Context, page, limit int) (PaginationResult, error) {
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

	// Buscar o total de clientes para metadados de paginação
	total, err := queries.CountClients(ctx)
	if err != nil {
		return PaginationResult{}, err
	}

	// Calcular o total de páginas
	totalPages := (int(total) + limit - 1) / limit

	// Buscar clientes com paginação via SQL
	paginatedClients, err := queries.FindManyClientsWithPagination(ctx, database.FindManyClientsWithPaginationParams{
		Offset: int32(offset),
		Limit:  int32(limit),
	})
	if err != nil {
		return PaginationResult{}, err
	}

	// Montar resultado com metadados de paginação
	result := PaginationResult{
		Data: paginatedClients,
		Meta: PaginationMeta{
			CurrentPage: page,
			PerPage:     limit,
			Total:       total,
			TotalPages:  totalPages,
		},
	}

	return result, nil
}

// GetClients retorna todos os clientes com paginação (mantido para compatibilidade)
func GetClients(ctx context.Context, offset, limit int32) ([]database.Client, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.FindManyClientsWithPagination(ctx, database.FindManyClientsWithPaginationParams{
		Offset: offset,
		Limit:  limit,
	})
}

// CountClients retorna o total de clientes (mantido para compatibilidade)
func CountClients(ctx context.Context) (int64, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.CountClients(ctx)
}

// GetClient retorna um cliente pelo ID
func GetClient(ctx context.Context, id int64) (database.Client, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.FindClientById(ctx, id)
}

// CreateClient cria um novo cliente
func CreateClient(ctx context.Context, params database.CreateClientParams) (database.Client, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.CreateClient(ctx, params)
}

// UpdateClient atualiza um cliente existente
func UpdateClient(ctx context.Context, params database.UpdateClientParams) (database.Client, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.UpdateClient(ctx, params)
}

// DeleteClient remove um cliente
func DeleteClient(ctx context.Context, id int64) error {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.DeleteClient(ctx, id)
}
