package commentHandler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"

	"sixTask/internal/database"
)

// GetComments retorna todos os comentários
func GetComments(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	comments, err := queries.FindManyComments(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar comentários: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// GetComment retorna um comentário pelo ID
func GetComment(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	queries := database.New(conn)
	comment, err := queries.FindCommentById(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comentário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, comment)
}

// GetCommentsByUser retorna comentários pelo ID do usuário
func GetCommentsByUser(c *gin.Context) {
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
	comments, err := queries.FindCommentsByUserId(ctx, userIdPg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar comentários: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// GetCommentsByCommentable retorna comentários pelo tipo e ID do objeto comentável
func GetCommentsByCommentable(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	commentableType := c.Param("commentable_type")
	if commentableType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tipo de objeto comentável inválido"})
		return
	}

	commentableId, err := strconv.ParseInt(c.Param("commentable_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do objeto comentável inválido"})
		return
	}

	params := database.FindCommentsByCommentableParams{
		CommentableType: commentableType,
		CommentableID:   commentableId,
	}

	queries := database.New(conn)
	comments, err := queries.FindCommentsByCommentable(ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar comentários: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// CreateComment cria um novo comentário
func CreateComment(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	var params database.CreateCommentParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	queries := database.New(conn)
	comment, err := queries.CreateComment(ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar comentário: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

// UpdateComment atualiza um comentário existente
func UpdateComment(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var params database.UpdateCommentParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}
	params.ID = id

	queries := database.New(conn)
	comment, err := queries.UpdateComment(ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar comentário: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, comment)
}

// DeleteComment remove um comentário
func DeleteComment(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	queries := database.New(conn)
	err = queries.DeleteComment(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao remover comentário: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comentário removido com sucesso"})
}
