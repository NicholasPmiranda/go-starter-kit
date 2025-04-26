package attachmentRepository

import (
	"context"

	"sixTask/internal/database"
)

// PaginationResult contém os anexos paginados e os metadados de paginação
type PaginationResult struct {
	Data []database.Attachment `json:"data"`
	Meta PaginationMeta        `json:"meta"`
}

// PaginationMeta contém os metadados de paginação
type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
	TotalPages  int   `json:"total_pages"`
}

// GetAttachmentsWithPagination retorna os anexos paginados e os metadados de paginação
func GetAttachmentsWithPagination(ctx context.Context, page, limit int) (PaginationResult, error) {
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

	// Buscar o total de anexos para metadados de paginação
	total, err := queries.CountAttachments(ctx)
	if err != nil {
		return PaginationResult{}, err
	}

	// Calcular o total de páginas
	totalPages := (int(total) + limit - 1) / limit

	// Buscar anexos com paginação via SQL
	paginatedAttachments, err := queries.FindManyAttachmentsWithPagination(ctx, database.FindManyAttachmentsWithPaginationParams{
		Offset: int32(offset),
		Limit:  int32(limit),
	})
	if err != nil {
		return PaginationResult{}, err
	}

	// Montar resultado com metadados de paginação
	result := PaginationResult{
		Data: paginatedAttachments,
		Meta: PaginationMeta{
			CurrentPage: page,
			PerPage:     limit,
			Total:       total,
			TotalPages:  totalPages,
		},
	}

	return result, nil
}

// GetAttachments retorna todos os anexos com paginação (mantido para compatibilidade)
func GetAttachments(ctx context.Context, offset, limit int32) ([]database.Attachment, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.FindManyAttachmentsWithPagination(ctx, database.FindManyAttachmentsWithPaginationParams{
		Offset: offset,
		Limit:  limit,
	})
}

// CountAttachments retorna o total de anexos (mantido para compatibilidade)
func CountAttachments(ctx context.Context) (int64, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.CountAttachments(ctx)
}

// GetAttachment retorna um anexo pelo ID
func GetAttachment(ctx context.Context, id int64) (database.Attachment, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.FindAttachmentById(ctx, id)
}

// CreateAttachment cria um novo anexo
func CreateAttachment(ctx context.Context, params database.CreateAttachmentParams) (database.Attachment, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.CreateAttachment(ctx, params)
}

// UpdateAttachment atualiza um anexo existente
func UpdateAttachment(ctx context.Context, params database.UpdateAttachmentParams) (database.Attachment, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.UpdateAttachment(ctx, params)
}

// DeleteAttachment remove um anexo
func DeleteAttachment(ctx context.Context, id int64) error {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.DeleteAttachment(ctx, id)
}
