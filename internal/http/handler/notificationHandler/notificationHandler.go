package notificationHandler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"

	"sixTask/internal/database"
)

// GetNotifications retorna todas as notificações
func GetNotifications(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	notifications, err := queries.FindManyNotifications(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar notificações: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, notifications)
}

// GetNotification retorna uma notificação pelo ID
func GetNotification(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	queries := database.New(conn)
	notification, err := queries.FindNotificationById(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Notificação não encontrada"})
		return
	}

	c.JSON(http.StatusOK, notification)
}

// GetNotificationsByUser retorna notificações pelo ID do usuário
func GetNotificationsByUser(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do usuário inválido"})
		return
	}

	// Converter int64 para pgtype.Int8
	userIdPg := pgtype.Int8{Int64: userId, Valid: true}

	queries := database.New(conn)
	notifications, err := queries.FindNotificationsByUserId(ctx, userIdPg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar notificações: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, notifications)
}

// GetUnreadNotificationsByUser retorna notificações não lidas pelo ID do usuário
func GetUnreadNotificationsByUser(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do usuário inválido"})
		return
	}

	// Converter int64 para pgtype.Int8
	userIdPg := pgtype.Int8{Int64: userId, Valid: true}

	queries := database.New(conn)
	notifications, err := queries.FindUnreadNotificationsByUserId(ctx, userIdPg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar notificações: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, notifications)
}

// GetNotificationsByNotifiable retorna notificações pelo tipo e ID do objeto notificável
func GetNotificationsByNotifiable(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	notifiableType := c.Param("notifiable_type")
	if notifiableType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tipo de objeto notificável inválido"})
		return
	}

	notifiableId, err := strconv.ParseInt(c.Param("notifiable_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do objeto notificável inválido"})
		return
	}

	params := database.FindNotificationsByNotifiableParams{
		NotifiableType: notifiableType,
		NotifiableID:   notifiableId,
	}

	queries := database.New(conn)
	notifications, err := queries.FindNotificationsByNotifiable(ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar notificações: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, notifications)
}

// CreateNotification cria uma nova notificação
func CreateNotification(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	var params database.CreateNotificationParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	queries := database.New(conn)
	notification, err := queries.CreateNotification(ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar notificação: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, notification)
}

// UpdateNotification atualiza uma notificação existente
func UpdateNotification(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var params database.UpdateNotificationParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}
	params.ID = id

	queries := database.New(conn)
	notification, err := queries.UpdateNotification(ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar notificação: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, notification)
}

// MarkNotificationAsRead marca uma notificação como lida
func MarkNotificationAsRead(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	queries := database.New(conn)
	notification, err := queries.MarkNotificationAsRead(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao marcar notificação como lida: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, notification)
}

// MarkAllNotificationsAsRead marca todas as notificações de um usuário como lidas
func MarkAllNotificationsAsRead(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do usuário inválido"})
		return
	}

	// Converter int64 para pgtype.Int8
	userIdPg := pgtype.Int8{Int64: userId, Valid: true}

	queries := database.New(conn)
	notifications, err := queries.MarkAllNotificationsAsRead(ctx, userIdPg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao marcar notificações como lidas: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, notifications)
}

// DeleteNotification remove uma notificação
func DeleteNotification(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	queries := database.New(conn)
	err = queries.DeleteNotification(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao remover notificação: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notificação removida com sucesso"})
}
