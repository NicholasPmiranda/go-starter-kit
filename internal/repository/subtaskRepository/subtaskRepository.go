package subtaskRepository

import (
	"context"

	"sixTask/internal/database"
)

// PaginationResult contém as subtarefas paginadas e os metadados de paginação
type PaginationResult struct {
	Data []database.Subtask `json:"data"`
	Meta PaginationMeta     `json:"meta"`
}

// PaginationMeta contém os metadados de paginação
type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
	TotalPages  int   `json:"total_pages"`
}

// GetSubtasksWithPagination retorna as subtarefas paginadas e os metadados de paginação
func GetSubtasksWithPagination(ctx context.Context, page, limit int) (PaginationResult, error) {
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

	// Buscar o total de subtarefas para metadados de paginação
	total, err := queries.CountSubtasks(ctx)
	if err != nil {
		return PaginationResult{}, err
	}

	// Calcular o total de páginas
	totalPages := (int(total) + limit - 1) / limit

	// Buscar subtarefas com paginação via SQL
	paginatedSubtasks, err := queries.FindManySubtasksWithPagination(ctx, database.FindManySubtasksWithPaginationParams{
		Offset: int32(offset),
		Limit:  int32(limit),
	})
	if err != nil {
		return PaginationResult{}, err
	}

	// Montar resultado com metadados de paginação
	result := PaginationResult{
		Data: paginatedSubtasks,
		Meta: PaginationMeta{
			CurrentPage: page,
			PerPage:     limit,
			Total:       total,
			TotalPages:  totalPages,
		},
	}

	return result, nil
}

// GetSubtasks retorna todas as subtarefas com paginação (mantido para compatibilidade)
func GetSubtasks(ctx context.Context, offset, limit int32) ([]database.Subtask, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.FindManySubtasksWithPagination(ctx, database.FindManySubtasksWithPaginationParams{
		Offset: offset,
		Limit:  limit,
	})
}

// CountSubtasks retorna o total de subtarefas (mantido para compatibilidade)
func CountSubtasks(ctx context.Context) (int64, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.CountSubtasks(ctx)
}

// GetSubtask retorna uma subtarefa pelo ID
func GetSubtask(ctx context.Context, id int64) (database.Subtask, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.FindSubtaskById(ctx, id)
}

// CreateSubtask cria uma nova subtarefa
func CreateSubtask(ctx context.Context, params database.CreateSubtaskParams) (database.Subtask, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.CreateSubtask(ctx, params)
}

// UpdateSubtask atualiza uma subtarefa existente
func UpdateSubtask(ctx context.Context, params database.UpdateSubtaskParams) (database.Subtask, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.UpdateSubtask(ctx, params)
}

// DeleteSubtask remove uma subtarefa
func DeleteSubtask(ctx context.Context, id int64) error {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.DeleteSubtask(ctx, id)
}
