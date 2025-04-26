package projectHandler

import (
	"context"
	"net/http"
	"sixTask/internal/entity/projectEntity"
	"sixTask/internal/repository/projectRepository"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"

	"sixTask/internal/database"
	"sixTask/internal/http/request/projectRequest"
	"sixTask/internal/http/validator"
)

// GetProjects retorna todos os projetos
func GetProjects(c *gin.Context) {
	// Obter parâmetros de paginação da query
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))

	// Chamar repositório para buscar projetos com paginação
	projects, total, err := projectRepository.GetProjectsWithUsersAndPagination(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar projetos: " + err.Error()})
		return
	}

	// Calcular total de páginas
	totalPages := (int(total) + limit - 1) / limit

	// Montar resposta com metadados de paginação
	c.JSON(http.StatusOK, gin.H{
		"data": projects,
		"meta": gin.H{
			"current_page": page,
			"per_page":     limit,
			"total":        total,
			"total_pages":  totalPages,
		},
	})
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
	project, err := queries.FindProjectWithUsers(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Projeto não encontrado"})
		return
	}

	c.JSON(http.StatusOK, project)
}

// GetProjectsByClient retorna projetos pelo ID do cliente
func GetProjectsByClient(c *gin.Context) {
	// Obter parâmetros de paginação da query
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))

	clientId, err := strconv.ParseInt(c.Param("client_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do cliente inválido"})
		return
	}

	// Converter int64 para pgtype.Int8
	clientIdPg := pgtype.Int8{Int64: clientId, Valid: true}

	// Chamar repositório para buscar projetos com paginação
	projects, total, err := projectRepository.GetProjectsByClientIdAndPagination(clientIdPg, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar projetos: " + err.Error()})
		return
	}

	// Calcular total de páginas
	totalPages := (int(total) + limit - 1) / limit

	// Montar resposta com metadados de paginação
	c.JSON(http.StatusOK, gin.H{
		"data": projects,
		"meta": gin.H{
			"current_page": page,
			"per_page":     limit,
			"total":        total,
			"total_pages":  totalPages,
		},
	})
}

// GetProjectsByUser retorna projetos pelo ID do usuário
func GetProjectsByUser(c *gin.Context) {
	// Obter parâmetros de paginação da query
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	// Chamar repositório para buscar projetos com paginação
	projects, total, err := projectRepository.GetProjectsWithUsersAndPagination(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar projetos: " + err.Error()})
		return
	}

	// Calcular total de páginas
	totalPages := (int(total) + limit - 1) / limit

	// Montar resposta com metadados de paginação
	c.JSON(http.StatusOK, gin.H{
		"data": projects,
		"meta": gin.H{
			"current_page": page,
			"per_page":     limit,
			"total":        total,
			"total_pages":  totalPages,
		},
	})
}

// CreateProject cria um novo projeto
func CreateProject(c *gin.Context) {

	var request projectRequest.CreateProjectRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + validator.Translate(err)})
		return
	}

	project, users, err := projectRepository.CreateProject(request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar projeto: " + err.Error()})
	}

	response := projectEntity.GetProjectEntity(project, users)

	c.JSON(http.StatusCreated, response)
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

	var request projectRequest.UpdateProjectRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + validator.Translate(err)})
		return
	}

	// Converte a request para o formato esperado pelo sqlc
	params := request.ToUpdateProjectParams(id).(database.UpdateProjectParams)

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
