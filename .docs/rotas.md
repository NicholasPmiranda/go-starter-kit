# Como Criar Rotas

## Visão Geral

O Go Starter Kit utiliza o framework [Gin](https://github.com/gin-gonic/gin) para gerenciar rotas HTTP. As rotas são definidas no arquivo `routes/api.go` e são organizadas em grupos para melhor estruturação e aplicação de middlewares.

## Estrutura de Rotas

As rotas são organizadas da seguinte forma:

1. **Rotas Públicas**: Acessíveis sem autenticação
2. **Rotas Autenticadas**: Protegidas pelo middleware de autenticação
3. **Rotas de Monitoramento**: Para monitoramento de jobs e outros recursos

## Como Definir Novas Rotas

### 1. Editar o Arquivo de Rotas

Para adicionar novas rotas, você deve editar o arquivo `routes/api.go`:

```go
package routes

import (
	"boilerPlate/internal/http/handler"
	"boilerPlate/internal/http/handler/JobHandler"
	authhandler "boilerPlate/internal/http/handler/authHandler"
	filehandler "boilerPlate/internal/http/handler/fileHandler"
	authmiddleware "boilerPlate/internal/middleware/authMiddleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Configuração do monitor de jobs
	// ...

	// Rota para servir arquivos estáticos
	router.GET("/storage/*filepath", filehandler.GetFileHandler())

	// Grupo de rotas da API
	api := router.Group("/api")
	{
		// Rotas públicas
		api.POST("/", handler.StartHandler)
		api.POST("/login", authhandler.Login)
		api.POST("/upload", filehandler.UploadFileExample)
		
		// Sua nova rota pública
		api.GET("/produtos", produtohandler.ListarProdutos)

		// Grupo de rotas autenticadas
		authenticated := api.Group("/")
		authenticated.Use(authmiddleware.AuthMiddleware())
		{
			authenticated.GET("/profile", authhandler.Profile)
			
			// Sua nova rota autenticada
			authenticated.POST("/produtos", produtohandler.CriarProduto)
		}
	}

	return router
}
```

### 2. Criar um Handler para a Rota

Cada rota deve ter um handler correspondente. Os handlers são organizados em pacotes dentro do diretório `internal/http/handler/`:

```go
package produtohandler

import (
	"github.com/gin-gonic/gin"
)

// ListarProdutos retorna a lista de produtos
func ListarProdutos(c *gin.Context) {
	// Lógica para listar produtos
	// ...

	c.JSON(200, gin.H{
		"produtos": produtos,
	})
}

// CriarProduto cria um novo produto
func CriarProduto(c *gin.Context) {
	// Obter o ID do usuário autenticado
	userID, _ := c.Get("authUser")

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

## Tipos de Rotas

O Gin suporta todos os métodos HTTP padrão:

```go
// GET
router.GET("/caminho", handler)

// POST
router.POST("/caminho", handler)

// PUT
router.PUT("/caminho", handler)

// DELETE
router.DELETE("/caminho", handler)

// PATCH
router.PATCH("/caminho", handler)

// OPTIONS
router.OPTIONS("/caminho", handler)

// HEAD
router.HEAD("/caminho", handler)

// Qualquer método
router.Any("/caminho", handler)
```

## Parâmetros de Rota

### Parâmetros na URL

```go
// /usuario/123
router.GET("/usuario/:id", func(c *gin.Context) {
	id := c.Param("id") // "123"
	// ...
})

// /arquivos/pasta/arquivo.txt
router.GET("/arquivos/*caminho", func(c *gin.Context) {
	caminho := c.Param("caminho") // "/pasta/arquivo.txt"
	// ...
})
```

### Parâmetros de Query

```go
// /produtos?pagina=2&limite=10
router.GET("/produtos", func(c *gin.Context) {
	pagina := c.DefaultQuery("pagina", "1")     // "2"
	limite := c.DefaultQuery("limite", "20")    // "10"
	// ...
})
```

## Grupos de Rotas

Os grupos de rotas permitem organizar rotas relacionadas e aplicar middlewares a um conjunto de rotas:

```go
// Grupo de rotas para API v1
v1 := router.Group("/api/v1")
{
	v1.GET("/usuarios", listarUsuarios)
	v1.POST("/usuarios", criarUsuario)
}

// Grupo de rotas para API v2
v2 := router.Group("/api/v2")
{
	v2.GET("/usuarios", listarUsuariosV2)
	v2.POST("/usuarios", criarUsuarioV2)
}

// Grupo de rotas com middleware
admin := router.Group("/admin")
admin.Use(authmiddleware.AdminMiddleware())
{
	admin.GET("/dashboard", adminDashboard)
	admin.GET("/usuarios", adminListarUsuarios)
}
```

## Middlewares

Os middlewares são funções que são executadas antes do handler da rota. Eles podem ser usados para autenticação, logging, tratamento de erros, etc.

```go
// Aplicar middleware a uma rota específica
router.GET("/rota", middleware1, middleware2, handler)

// Aplicar middleware a um grupo de rotas
grupo := router.Group("/grupo")
grupo.Use(middleware1, middleware2)
{
	grupo.GET("/rota1", handler1)
	grupo.GET("/rota2", handler2)
}
```

## Boas Práticas

1. **Organização**: Mantenha as rotas organizadas em grupos lógicos.
2. **Nomenclatura**: Use nomes descritivos para as rotas e siga padrões RESTful quando apropriado.
3. **Versionamento**: Considere versionar sua API (ex: `/api/v1/recurso`).
4. **Validação**: Valide os dados de entrada antes de processá-los.
5. **Respostas**: Padronize as respostas da API para facilitar o consumo por clientes.
6. **Documentação**: Mantenha a documentação da API atualizada.
7. **Tratamento de Erros**: Implemente tratamento adequado de erros para fornecer mensagens úteis aos clientes.
