package taskHandler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"

	"sixTask/internal/database"
	"sixTask/internal/http/validator"
	"sixTask/internal/repository/taskRepository"
)

// GetTasks retorna todas as tarefas com paginação e informações de usuário
func GetTasks(c *gin.Context) {
	// Extrair parâmetros de paginação da query
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 || limit > 100 {
		limit = 10
	}

	// Usar o repository para buscar tarefas com paginação e informações de usuário
	tasks, err := taskRepository.GetTasksWithPaginationAndUsers(context.Background(), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar tarefas: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetTask retorna uma tarefa pelo ID
func GetTask(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	queries := database.New(conn)
	task, err := queries.FindTaskById(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não encontrada"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// GetTasksByProject retorna tarefas pelo ID do projeto
func GetTasksByProject(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	projectId, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do projeto inválido"})
		return
	}

	// Converter int64 para pgtype.Int8
	projectIdPg := pgtype.Int8{Int64: projectId, Valid: true}

	queries := database.New(conn)
	tasks, err := queries.FindTasksByProjectId(ctx, projectIdPg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar tarefas: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetTasksByAssignedTo retorna tarefas pelo ID do usuário atribuído
func GetTasksByAssignedTo(c *gin.Context) {
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
	tasks, err := queries.FindTasksByAssignedTo(ctx, userIdPg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar tarefas: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetTasksByStatus retorna tarefas pelo status
func GetTasksByStatus(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	status := c.Param("status")
	if status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status inválido"})
		return
	}

	queries := database.New(conn)
	tasks, err := queries.FindTasksByStatus(ctx, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar tarefas: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetTasksByPriority retorna tarefas pela prioridade
func GetTasksByPriority(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	priority := c.Param("priority")
	if priority == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Prioridade inválida"})
		return
	}

	queries := database.New(conn)
	tasks, err := queries.FindTasksByPriority(ctx, priority)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar tarefas: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// CreateTask cria uma nova tarefa
func CreateTask(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	var params database.CreateTaskParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + validator.Translate(err)})
		return
	}

	queries := database.New(conn)
	task, err := queries.CreateTask(ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar tarefa: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// UpdateTask atualiza uma tarefa existente
func UpdateTask(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var params database.UpdateTaskParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + validator.Translate(err)})
		return
	}
	params.ID = id

	queries := database.New(conn)
	task, err := queries.UpdateTask(ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar tarefa: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

// CompleteTask marca uma tarefa como concluída
func CompleteTask(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	queries := database.New(conn)
	task, err := queries.CompleteTask(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao concluir tarefa: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

// DeleteTask remove uma tarefa
func DeleteTask(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	queries := database.New(conn)
	err = queries.DeleteTask(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir tarefa: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tarefa excluída com sucesso"})
}
