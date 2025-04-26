package notificationRepository

import (
	"context"

	"sixTask/internal/database"
)

// PaginationResult contém as notificações paginadas e os metadados de paginação
type PaginationResult struct {
	Data []database.Notification `json:"data"`
	Meta PaginationMeta          `json:"meta"`
}

// PaginationMeta contém os metadados de paginação
type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
	TotalPages  int   `json:"total_pages"`
}

// GetNotificationsWithPagination retorna as notificações paginadas e os metadados de paginação
func GetNotificationsWithPagination(ctx context.Context, page, limit int) (PaginationResult, error) {
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

	// Buscar o total de notificações para metadados de paginação
	total, err := queries.CountNotifications(ctx)
	if err != nil {
		return PaginationResult{}, err
	}

	// Calcular o total de páginas
	totalPages := (int(total) + limit - 1) / limit

	// Buscar notificações com paginação via SQL
	paginatedNotifications, err := queries.FindManyNotificationsWithPagination(ctx, database.FindManyNotificationsWithPaginationParams{
		Offset: int32(offset),
		Limit:  int32(limit),
	})
	if err != nil {
		return PaginationResult{}, err
	}

	// Montar resultado com metadados de paginação
	result := PaginationResult{
		Data: paginatedNotifications,
		Meta: PaginationMeta{
			CurrentPage: page,
			PerPage:     limit,
			Total:       total,
			TotalPages:  totalPages,
		},
	}

	return result, nil
}

// GetNotifications retorna todas as notificações com paginação (mantido para compatibilidade)
func GetNotifications(ctx context.Context, offset, limit int32) ([]database.Notification, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.FindManyNotificationsWithPagination(ctx, database.FindManyNotificationsWithPaginationParams{
		Offset: offset,
		Limit:  limit,
	})
}

// CountNotifications retorna o total de notificações (mantido para compatibilidade)
func CountNotifications(ctx context.Context) (int64, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.CountNotifications(ctx)
}

// GetNotification retorna uma notificação pelo ID
func GetNotification(ctx context.Context, id int64) (database.Notification, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.FindNotificationById(ctx, id)
}

// CreateNotification cria uma nova notificação
func CreateNotification(ctx context.Context, params database.CreateNotificationParams) (database.Notification, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.CreateNotification(ctx, params)
}

// UpdateNotification atualiza uma notificação existente
func UpdateNotification(ctx context.Context, params database.UpdateNotificationParams) (database.Notification, error) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.UpdateNotification(ctx, params)
}

// DeleteNotification remove uma notificação
func DeleteNotification(ctx context.Context, id int64) error {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	return queries.DeleteNotification(ctx, id)
}
