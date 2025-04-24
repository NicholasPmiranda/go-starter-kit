# Como Usar Handlers

## Visão Geral

Os handlers são componentes responsáveis por processar requisições HTTP no Go Starter Kit. Eles são implementados usando o framework [Gin](https://github.com/gin-gonic/gin) e são organizados em pacotes dentro do diretório `internal/http/handler/`.

## Estrutura de Handlers

Os handlers são organizados por domínio ou funcionalidade:

```
internal/http/handler/
├── JobHandler/           # Handlers relacionados a jobs
├── authHandler/          # Handlers relacionados a autenticação
├── fileHandler/          # Handlers relacionados a arquivos
└── startHandler.go       # Handler inicial
```

## Como Criar um Novo Handler

### 1. Criar um Novo Pacote de Handler

Para funcionalidades relacionadas, crie um novo pacote dentro do diretório `internal/http/handler/`:

```go
// internal/http/handler/produtoHandler/produtoHandler.go
package produtohandler

import (
	"github.com/gin-gonic/gin"
)

// Funções de handler aqui...
```

### 2. Implementar Funções de Handler

Cada função de handler deve receber um ponteiro para `gin.Context` e processar a requisição:

```go
// ListarProdutos retorna a lista de produtos
func ListarProdutos(c *gin.Context) {
	// Lógica para listar produtos
	// ...

	c.JSON(200, gin.H{
		"produtos": produtos,
	})
}

// ObterProduto retorna um produto específico
func ObterProduto(c *gin.Context) {
	// Obter o ID do produto da URL
	id := c.Param("id")

	// Lógica para buscar o produto
	// ...

	c.JSON(200, gin.H{
		"produto": produto,
	})
}

// CriarProduto cria um novo produto
func CriarProduto(c *gin.Context) {
	// Receber dados do produto
	var produto Produto
	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Lógica para criar o produto
	// ...

	c.JSON(201, gin.H{
		"message": "Produto criado com sucesso",
		"produto": produto,
	})
}
```

### 3. Registrar as Rotas

Após criar as funções de handler, você deve registrá-las nas rotas da aplicação no arquivo `routes/api.go`:

```go
// routes/api.go
import (
	// ...
	produtohandler "sixTask/internal/http/handler/produtoHandler"
)

func SetupRoutes() *gin.Engine {
	// ...
	
	api := router.Group("/api")
	{
		// ...
		
		// Rotas de produtos
		api.GET("/produtos", produtohandler.ListarProdutos)
		api.GET("/produtos/:id", produtohandler.ObterProduto)
		api.POST("/produtos", produtohandler.CriarProduto)
	}
	
	// ...
}
```

## Estrutura de um Handler

Um handler típico segue esta estrutura:

1. **Validação de Entrada**: Validar os dados recebidos na requisição
2. **Processamento**: Executar a lógica de negócio
3. **Resposta**: Retornar uma resposta apropriada

### Exemplo Completo

```go
package userhandler

import (
	"sixTask/internal/database"
	"sixTask/internal/http/request/RequestModel"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetUser retorna um usuário pelo ID
func GetUser(c *gin.Context) {
	// 1. Validação de Entrada
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	// 2. Processamento
	db := c.MustGet("db").(*database.Queries)
	user, err := db.FindById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Usuário não encontrado",
		})
		return
	}

	// 3. Resposta
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// CreateUser cria um novo usuário
func CreateUser(c *gin.Context) {
	// 1. Validação de Entrada
	var request RequestModel.CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Validações adicionais
	if len(request.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "A senha deve ter pelo menos 6 caracteres",
		})
		return
	}

	// 2. Processamento
	db := c.MustGet("db").(*database.Queries)
	
	// Verificar se o email já existe
	_, err := db.FindByEmail(context.Background(), request.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email já cadastrado",
		})
		return
	}

	// Criar o usuário
	user, err := db.CreateUser(context.Background(), database.CreateUserParams{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password, // Na prática, deve-se fazer hash da senha
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao criar usuário",
		})
		return
	}

	// 3. Resposta
	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuário criado com sucesso",
		"user": user,
	})
}
```

## Recebendo Dados da Requisição

### Parâmetros de URL

```go
// /api/produtos/123
func ObterProduto(c *gin.Context) {
	id := c.Param("id") // "123"
	// ...
}
```

### Parâmetros de Query

```go
// /api/produtos?pagina=2&limite=10
func ListarProdutos(c *gin.Context) {
	pagina := c.DefaultQuery("pagina", "1")     // "2"
	limite := c.DefaultQuery("limite", "20")    // "10"
	// ...
}
```

### Corpo da Requisição (JSON)

```go
func CriarProduto(c *gin.Context) {
	var produto Produto
	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	// ...
}
```

### Formulários

```go
func ProcessarFormulario(c *gin.Context) {
	nome := c.PostForm("nome")
	email := c.PostForm("email")
	// ...
}
```

### Upload de Arquivos

```go
func UploadArquivo(c *gin.Context) {
	file, err := c.FormFile("arquivo")
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// Salvar o arquivo
	dst := "storage/uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// ...
}
```

## Retornando Respostas

### Resposta JSON

```go
c.JSON(200, gin.H{
	"message": "Operação realizada com sucesso",
	"data": resultado,
})
```

### Resposta com Status HTTP

```go
c.JSON(http.StatusOK, gin.H{
	"message": "Operação realizada com sucesso",
})

c.JSON(http.StatusBadRequest, gin.H{
	"error": "Dados inválidos",
})

c.JSON(http.StatusNotFound, gin.H{
	"error": "Recurso não encontrado",
})

c.JSON(http.StatusInternalServerError, gin.H{
	"error": "Erro interno do servidor",
})
```

### Resposta HTML

```go
c.HTML(http.StatusOK, "template.sql.html", gin.H{
	"titulo": "Minha Página",
	"conteudo": "Conteúdo da página",
})
```

### Resposta de Arquivo

```go
c.File("caminho/para/arquivo.pdf")
```

### Resposta de Redirecionamento

```go
c.Redirect(http.StatusFound, "/nova-url")
```

## Acessando Serviços e Dependências

Os handlers podem acessar serviços e dependências através do contexto do Gin:

```go
// Acessar o banco de dados
db := c.MustGet("db").(*database.Queries)

// Acessar o logger
logger := c.MustGet("logger").(*zap.Logger)

// Acessar o usuário autenticado
userID, _ := c.Get("authUser")
```

## Tratamento de Erros

É importante tratar erros adequadamente nos handlers:

```go
func ExemploHandler(c *gin.Context) {
	// ...
	
	resultado, err := algumaOperacao()
	if err != nil {
		// Log do erro
		logger := c.MustGet("logger").(*zap.Logger)
		logger.Error("Erro ao executar operação", zap.Error(err))
		
		// Resposta de erro
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ocorreu um erro ao processar sua solicitação",
		})
		return
	}
	
	// ...
}
```

## Boas Práticas

1. **Separação de Responsabilidades**: Os handlers devem se concentrar apenas na manipulação da requisição HTTP. A lógica de negócio deve ser delegada para serviços.
2. **Validação de Entrada**: Sempre valide os dados de entrada antes de processá-los.
3. **Tratamento de Erros**: Trate todos os erros possíveis e forneça mensagens de erro claras.
4. **Respostas Consistentes**: Mantenha um formato consistente para suas respostas.
5. **Logging**: Adicione logs adequados para facilitar o diagnóstico de problemas.
6. **Segurança**: Esteja atento a questões de segurança, como injeção de SQL, XSS, CSRF, etc.
7. **Testes**: Escreva testes para seus handlers para garantir que eles funcionem corretamente.
