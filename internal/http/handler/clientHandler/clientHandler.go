package clientHandler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"

	"sixTask/internal/database"
	"sixTask/internal/http/request/clientRequest"
	"sixTask/internal/http/validator"
	"sixTask/internal/repository/clientRepository"
)

// GetClients retorna todos os clientes com paginação
func GetClients(c *gin.Context) {
	ctx := context.Background()

	// Parâmetros de paginação
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	// Buscar clientes com paginação via repositório
	result, err := clientRepository.GetClientsWithPagination(ctx, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar clientes: " + err.Error()})
		return
	}

	// Retornar o resultado
	c.JSON(http.StatusOK, result)
}

// GetClient retorna um cliente pelo ID
func GetClient(c *gin.Context) {
	ctx := context.Background()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	client, err := clientRepository.GetClient(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
		return
	}

	c.JSON(http.StatusOK, client)
}

// CreateClient cria um novo cliente
func CreateClient(c *gin.Context) {
	ctx := context.Background()

	var request clientRequest.CreateClientRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + validator.Translate(err)})
		return
	}

	// Converte a request para o formato esperado pelo sqlc
	params := request.ToCreateClientParams().(struct {
		Name    string      `json:"name"`
		Email   string      `json:"email"`
		Phone   pgtype.Text `json:"phone"`
		Address pgtype.Text `json:"address"`
	})

	// Cria o cliente usando o repositório
	client, err := clientRepository.CreateClient(ctx, database.CreateClientParams{
		Name:    params.Name,
		Email:   params.Email,
		Phone:   params.Phone,
		Address: params.Address,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar cliente: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, client)
}

// UpdateClient atualiza um cliente existente
func UpdateClient(c *gin.Context) {
	ctx := context.Background()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var request clientRequest.UpdateClientRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + validator.Translate(err)})
		return
	}

	// Converte a request para o formato esperado pelo sqlc
	params := request.ToUpdateClientParams(id).(struct {
		Name    string      `json:"name"`
		Email   string      `json:"email"`
		Phone   pgtype.Text `json:"phone"`
		Address pgtype.Text `json:"address"`
		ID      int64       `json:"id"`
	})

	// Atualiza o cliente usando o repositório
	client, err := clientRepository.UpdateClient(ctx, database.UpdateClientParams{
		Name:    params.Name,
		Email:   params.Email,
		Phone:   params.Phone,
		Address: params.Address,
		ID:      params.ID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar cliente: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, client)
}

// DeleteClient remove um cliente
func DeleteClient(c *gin.Context) {
	ctx := context.Background()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = clientRepository.DeleteClient(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao remover cliente: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cliente removido com sucesso"})
}
