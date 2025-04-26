package taskRepository

import (
	"context"

	"sixTask/internal/database"
)

// PaginationResult contém as tarefas paginadas e os metadados de paginação
type PaginationResult struct {
	Data []database.Task `json:"data"`
	Meta PaginationMeta  `json:"meta"`
}

// PaginationMeta contém os metadados de paginação
type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
	TotalPages  int   `json:"total_pages"`
}

// GetTasksWithPagination retorna as tarefas paginadas e os metadados de paginação
func GetTasksWithPagination(ctx context.Context, page, limit int) (PaginationResult, error) {
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

	// Buscar o total de tarefas para metadados de paginação
	total, err := queries.CountTasks(ctx)
	if err != nil {
		return PaginationResult{}, err
	}

	// Calcular o total de páginas
	totalPages := (int(total) + limit - 1) / limit

	// Buscar tarefas com paginação via SQL
	paginatedTasks, err := queries.FindManyTasksWithPagination(ctx, database.FindManyTasksWithPaginationParams{
		Offset: int32(offset),
		Limit:  int32(limit),
	})
	if err != nil {
		return PaginationResult{}, err
	}

	// Montar resultado com metadados de paginação
	result := PaginationResult{
		Data: paginatedTasks,
		Meta: PaginationMeta{
			CurrentPage: page,
			PerPage:     limit,
			Total:       total,
			TotalPages:  totalPages,
		},
	}

	return result, nil
}

// GetTasks retorna todas as tarefas com paginação (mantido para compatibilidade)
func GetTasks(ctx context.Context, offset, limit int32) ([]database.Task, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.FindManyTasksWithPagination(ctx, database.FindManyTasksWithPaginationParams{
		Offset: offset,
		Limit:  limit,
	})
}

// CountTasks retorna o total de tarefas (mantido para compatibilidade)
func CountTasks(ctx context.Context) (int64, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.CountTasks(ctx)
}

// GetTask retorna uma tarefa pelo ID
func GetTask(ctx context.Context, id int64) (database.Task, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.FindTaskById(ctx, id)
}

// CreateTask cria uma nova tarefa
func CreateTask(ctx context.Context, params database.CreateTaskParams) (database.Task, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.CreateTask(ctx, params)
}

// UpdateTask atualiza uma tarefa existente
func UpdateTask(ctx context.Context, params database.UpdateTaskParams) (database.Task, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.UpdateTask(ctx, params)
}

// DeleteTask remove uma tarefa
func DeleteTask(ctx context.Context, id int64) error {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.DeleteTask(ctx, id)
}
