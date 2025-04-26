package attachmentHandler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"

	"sixTask/internal/database"
	"sixTask/internal/http/request/attachmentRequest"
)

// GetAttachments retorna todos os anexos
func GetAttachments(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	attachments, err := queries.FindManyAttachments(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar anexos: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, attachments)
}

// GetAttachment retorna um anexo pelo ID
func GetAttachment(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	queries := database.New(conn)
	attachment, err := queries.FindAttachmentById(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Anexo não encontrado"})
		return
	}

	c.JSON(http.StatusOK, attachment)
}

// GetAttachmentsByUser retorna anexos pelo ID do usuário
func GetAttachmentsByUser(c *gin.Context) {
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
	attachments, err := queries.FindAttachmentsByUserId(ctx, userIdPg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar anexos: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, attachments)
}

// GetAttachmentsByAttachable retorna anexos pelo tipo e ID do objeto anexável
func GetAttachmentsByAttachable(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	attachableType := c.Param("attachable_type")
	if attachableType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tipo de objeto anexável inválido"})
		return
	}

	attachableId, err := strconv.ParseInt(c.Param("attachable_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do objeto anexável inválido"})
		return
	}

	params := database.FindAttachmentsByAttachableParams{
		AttachableType: attachableType,
		AttachableID:   attachableId,
	}

	queries := database.New(conn)
	attachments, err := queries.FindAttachmentsByAttachable(ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar anexos: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, attachments)
}

// CreateAttachment cria um novo anexo
func CreateAttachment(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	var request attachmentRequest.CreateAttachmentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	// Converte a request para o formato esperado pelo sqlc
	params := request.ToCreateAttachmentParams().(database.CreateAttachmentParams)

	queries := database.New(conn)
	attachment, err := queries.CreateAttachment(ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar anexo: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, attachment)
}

// UpdateAttachment atualiza um anexo existente
func UpdateAttachment(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var request attachmentRequest.UpdateAttachmentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	// Converte a request para o formato esperado pelo sqlc
	params := request.ToUpdateAttachmentParams(id).(database.UpdateAttachmentParams)

	queries := database.New(conn)
	attachment, err := queries.UpdateAttachment(ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar anexo: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, attachment)
}

// DeleteAttachment remove um anexo
func DeleteAttachment(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	queries := database.New(conn)
	err = queries.DeleteAttachment(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao remover anexo: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Anexo removido com sucesso"})
}
