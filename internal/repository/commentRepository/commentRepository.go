package commentRepository

import (
	"context"

	"sixTask/internal/database"
)

// PaginationResult contém os comentários paginados e os metadados de paginação
type PaginationResult struct {
	Data []database.Comment `json:"data"`
	Meta PaginationMeta     `json:"meta"`
}

// PaginationMeta contém os metadados de paginação
type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
	TotalPages  int   `json:"total_pages"`
}

// GetCommentsWithPagination retorna os comentários paginados e os metadados de paginação
func GetCommentsWithPagination(ctx context.Context, page, limit int) (PaginationResult, error) {
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

	// Buscar o total de comentários para metadados de paginação
	total, err := queries.CountComments(ctx)
	if err != nil {
		return PaginationResult{}, err
	}

	// Calcular o total de páginas
	totalPages := (int(total) + limit - 1) / limit

	// Buscar comentários com paginação via SQL
	paginatedComments, err := queries.FindManyCommentsWithPagination(ctx, database.FindManyCommentsWithPaginationParams{
		Offset: int32(offset),
		Limit:  int32(limit),
	})
	if err != nil {
		return PaginationResult{}, err
	}

	// Montar resultado com metadados de paginação
	result := PaginationResult{
		Data: paginatedComments,
		Meta: PaginationMeta{
			CurrentPage: page,
			PerPage:     limit,
			Total:       total,
			TotalPages:  totalPages,
		},
	}

	return result, nil
}

// GetComments retorna todos os comentários com paginação (mantido para compatibilidade)
func GetComments(ctx context.Context, offset, limit int32) ([]database.Comment, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.FindManyCommentsWithPagination(ctx, database.FindManyCommentsWithPaginationParams{
		Offset: offset,
		Limit:  limit,
	})
}

// CountComments retorna o total de comentários (mantido para compatibilidade)
func CountComments(ctx context.Context) (int64, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.CountComments(ctx)
}

// GetComment retorna um comentário pelo ID
func GetComment(ctx context.Context, id int64) (database.Comment, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.FindCommentById(ctx, id)
}

// CreateComment cria um novo comentário
func CreateComment(ctx context.Context, params database.CreateCommentParams) (database.Comment, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.CreateComment(ctx, params)
}

// UpdateComment atualiza um comentário existente
func UpdateComment(ctx context.Context, params database.UpdateCommentParams) (database.Comment, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.UpdateComment(ctx, params)
}

// DeleteComment remove um comentário
func DeleteComment(ctx context.Context, id int64) error {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.DeleteComment(ctx, id)
}
