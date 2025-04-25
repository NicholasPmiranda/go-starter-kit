package projectHandler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"

	"sixTask/internal/database"
)

// GetProjects retorna todos os projetos
func GetProjects(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	queries := database.New(conn)
	projects, err := queries.FindManyProjects(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar projetos: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, projects)
}

// GetProject retorna um projeto pelo ID
func GetProject(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	queries := database.New(conn)
	project, err := queries.FindProjectById(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Projeto não encontrado"})
		return
	}

	c.JSON(http.StatusOK, project)
}

// GetProjectsByClient retorna projetos pelo ID do cliente
func GetProjectsByClient(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	clientId, err := strconv.ParseInt(c.Param("client_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do cliente inválido"})
		return
	}

	// Converter int64 para pgtype.Int8
	clientIdPg := pgtype.Int8{Int64: clientId, Valid: true}

	queries := database.New(conn)
	projects, err := queries.FindProjectsByClientId(ctx, clientIdPg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar projetos: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, projects)
}

// GetProjectsByUser retorna projetos pelo ID do usuário
func GetProjectsByUser(c *gin.Context) {
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
	projects, err := queries.FindProjectsByUserId(ctx, userIdPg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar projetos: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, projects)
}

// CreateProject cria um novo projeto
func CreateProject(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	var params database.CreateProjectParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	queries := database.New(conn)
	project, err := queries.CreateProject(ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar projeto: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, project)
}

// UpdateProject atualiza um projeto existente
func UpdateProject(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var params database.UpdateProjectParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}
	params.ID = id

	queries := database.New(conn)
	project, err := queries.UpdateProject(ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar projeto: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, project)
}

// DeleteProject remove um projeto
func DeleteProject(c *gin.Context) {
	conn, ctx := database.ConnectDB()
	defer conn.Close(context.Background())

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	queries := database.New(conn)
	err = queries.DeleteProject(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao remover projeto: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Projeto removido com sucesso"})
}
