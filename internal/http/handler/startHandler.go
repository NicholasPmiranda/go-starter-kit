package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"

	"boilerPlate/internal/database"
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

	// Busca todos os usuários usando sqlx
	var usuarios []database.User
	err = db.Select(&usuarios, "SELECT id, name, email, password FROM users")
	if err != nil {
		log.Printf("Erro ao buscar usuários: %v", err)
		c.JSON(500, gin.H{
			"erro": "Erro ao buscar usuários: " + err.Error(),
		})
		return
	}

	log.Printf("Total de usuários encontrados: %d", len(usuarios))

	// Retorna o usuário criado e a lista completa de usuários
	log.Printf("Operação concluída com sucesso. Usuário criado com ID: %d", novoUsuario.ID)

	c.JSON(200, gin.H{
		"mensagem":      "Usuário criado com sucesso",
		"usuarioCriado": novoUsuario,
		"todosUsuarios": usuarios,
		"totalUsuarios": len(usuarios),
	})
}
