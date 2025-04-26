package subtaskHandler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"

	"sixTask/internal/database"
	"sixTask/internal/http/request/subtaskRequest"
)

// GetSubtasks retorna todas as subtarefas
func GetSubtasks(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	subtasks, err := queries.FindManySubtasks(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar subtarefas: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, subtasks)
}

// GetSubtask retorna uma subtarefa pelo ID
func GetSubtask(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	queries := database.New(conn)
	subtask, err := queries.FindSubtaskById(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subtarefa não encontrada"})
		return
	}

	c.JSON(http.StatusOK, subtask)
}

// GetSubtasksByTask retorna subtarefas pelo ID da tarefa
func GetSubtasksByTask(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	taskId, err := strconv.ParseInt(c.Param("task_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID da tarefa inválido"})
		return
	}

	// Converter int64 para pgtype.Int8
	taskIdPg := pgtype.Int8{Int64: taskId, Valid: true}

	queries := database.New(conn)
	subtasks, err := queries.FindSubtasksByTaskId(ctx, taskIdPg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar subtarefas: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, subtasks)
}

// GetSubtasksByAssignedTo retorna subtarefas pelo ID do usuário atribuído
func GetSubtasksByAssignedTo(c *gin.Context) {
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
	subtasks, err := queries.FindSubtasksByAssignedTo(ctx, userIdPg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar subtarefas: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, subtasks)
}

// GetSubtasksByStatus retorna subtarefas pelo status
func GetSubtasksByStatus(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	status := c.Param("status")
	if status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status inválido"})
		return
	}

	queries := database.New(conn)
	subtasks, err := queries.FindSubtasksByStatus(ctx, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar subtarefas: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, subtasks)
}

// CreateSubtask cria uma nova subtarefa
func CreateSubtask(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	var request subtaskRequest.CreateSubtaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	// Converte a request para o formato esperado pelo sqlc
	params := request.ToCreateSubtaskParams().(database.CreateSubtaskParams)

	queries := database.New(conn)
	subtask, err := queries.CreateSubtask(ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar subtarefa: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, subtask)
}

// UpdateSubtask atualiza uma subtarefa existente
func UpdateSubtask(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var request subtaskRequest.UpdateSubtaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	// Converte a request para o formato esperado pelo sqlc
	params := request.ToUpdateSubtaskParams(id).(database.UpdateSubtaskParams)

	queries := database.New(conn)
	subtask, err := queries.UpdateSubtask(ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar subtarefa: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, subtask)
}

// CompleteSubtask marca uma subtarefa como concluída
func CompleteSubtask(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	queries := database.New(conn)
	subtask, err := queries.CompleteSubtask(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao concluir subtarefa: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, subtask)
}

// DeleteSubtask remove uma subtarefa
func DeleteSubtask(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	queries := database.New(conn)
	err = queries.DeleteSubtask(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao remover subtarefa: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Subtarefa removida com sucesso"})
}
