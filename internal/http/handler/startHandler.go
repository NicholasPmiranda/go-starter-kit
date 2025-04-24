package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	emailprovider "sixTask/config/emailProvider"
	"time"

	"sixTask/internal/database"
)

func StartHandler(c *gin.Context) {
	// Conecta ao banco de dados usando sqlx
	db := database.ConnectDBX()
	defer db.Close()

	// Gera um email único usando timestamp
	timestamp := time.Now().UnixNano()
	emailUnico := fmt.Sprintf("usuario_%d@exemplo.com", timestamp)

	// Cria um novo usuário usando sqlx
	log.Printf("Criando novo usuário com email: %s", emailUnico)

	var novoUsuario database.User
	err := db.QueryRowx(
		"INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING *",
		"Novo Usuário", emailUnico, "senha123",
	).StructScan(&novoUsuario)

	if err != nil {
		c.JSON(500, gin.H{
			"erro": "Erro ao criar usuário: " + err.Error(),
		})
		return
	}

	to := []string{
		"teste@teste.com",
		"teste2@testeco",
	}

	templateData := map[string]interface{}{
		"Nome":    novoUsuario.Name,
		"Empresa": "Sua Empresa",
	}

	emailprovider.SendMail(emailprovider.EmailMessage{
		To:           to,
		Subject:      "teste",
		Template:     "cadastro",
		TemplateData: templateData,
	})

	log.Print("email enviado")
	c.JSON(200, gin.H{
		"mensagem":      "Usuário criado com sucesso",
		"usuarioCriado": novoUsuario,
	})
}
