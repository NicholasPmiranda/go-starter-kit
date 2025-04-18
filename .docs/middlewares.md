# Como Usar Middlewares

## Visão Geral

Os middlewares são componentes que interceptam requisições HTTP antes que elas cheguem aos handlers finais. Eles são úteis para executar operações comuns como autenticação, logging, tratamento de erros, entre outros. No Go Starter Kit, os middlewares são implementados usando o framework [Gin](https://github.com/gin-gonic/gin) e são organizados no diretório `internal/middleware/`.

## Estrutura de Middlewares

Os middlewares são organizados por funcionalidade:

```
internal/middleware/
└── authMiddleware/       # Middleware de autenticação
```

## Como Criar um Novo Middleware

### 1. Criar um Novo Pacote de Middleware

Para criar um novo middleware, crie um novo pacote dentro do diretório `internal/middleware/`:

```go
// internal/middleware/logMiddleware/logMiddleware.go
package logmiddleware

import (
	"github.com/gin-gonic/gin"
	"time"
)

// Funções de middleware aqui...
```

### 2. Implementar a Função de Middleware

Um middleware no Gin é uma função que recebe um ponteiro para `gin.Context` e chama `c.Next()` para continuar a execução da cadeia de middlewares:

```go
// LogMiddleware registra informações sobre a requisição
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Antes da requisição
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// Processa a requisição
		c.Next()

		// Após a requisição
		latency := time.Since(start)
		status := c.Writer.Status()
		
		// Log da requisição
		log.Printf("[%s] %s %s %d %s", method, path, status, latency)
	}
}
```

### 3. Registrar o Middleware

Após criar o middleware, você pode registrá-lo de diferentes formas:

#### Middleware Global (aplicado a todas as rotas)

```go
// routes/api.go
func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Aplicar middleware global
	router.Use(logmiddleware.LogMiddleware())

	// Resto da configuração de rotas
	// ...

	return router
}
```

#### Middleware para um Grupo de Rotas

```go
// routes/api.go
func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Grupo de rotas da API
	api := router.Group("/api")
	api.Use(logmiddleware.LogMiddleware()) // Aplicar middleware ao grupo
	{
		// Rotas do grupo
		// ...
	}

	return router
}
```

#### Middleware para uma Rota Específica

```go
// routes/api.go
func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Aplicar middleware a uma rota específica
	router.GET("/rota", logmiddleware.LogMiddleware(), handler.RoteHandler)

	return router
}
```

## Exemplo de Middleware de Autenticação

O Go Starter Kit inclui um middleware de autenticação que verifica tokens JWT:

```go
package authmiddleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	authhelper "boilerPlate/helpers/authHelper"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := authhelper.GetSecret()

	return func(c *gin.Context) {
		// Obter o token do cabeçalho Authorization
		bearer := c.GetHeader("Authorization")

		if bearer == "" || !strings.HasPrefix(bearer, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			return
		}

		// Extrair o token
		tokenStr := strings.Split(bearer, " ")[1]
		claims := &authhelper.Claims{}

		// Validar o token
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
			}
			return secretKey, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido: " + err.Error()})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expirado"})
			return
		}

		// Armazenar o ID do usuário no contexto
		c.Set("authUser", claims.UserID)
		
		// Continuar a execução
		c.Next()
	}
}
```

## Tipos Comuns de Middlewares

### Middleware de CORS

```go
package corsmiddleware

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
```

### Middleware de Recuperação de Pânico

```go
package recoverymiddleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recuperado: %v", err)
				
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "Ocorreu um erro interno no servidor",
				})
			}
		}()
		
		c.Next()
	}
}
```

### Middleware de Rate Limiting

```go
package ratelimitmiddleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
	"time"
)

type client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	clients = make(map[string]*client)
	mu      sync.Mutex
)

// Limpa clientes antigos periodicamente
func init() {
	go func() {
		for {
			time.Sleep(time.Minute)
			
			mu.Lock()
			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}
			mu.Unlock()
		}
	}()
}

func RateLimitMiddleware(rps float64, burst int) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		
		mu.Lock()
		if _, exists := clients[ip]; !exists {
			clients[ip] = &client{
				limiter: rate.NewLimiter(rate.Limit(rps), burst),
			}
		}
		
		clients[ip].lastSeen = time.Now()
		
		if !clients[ip].limiter.Allow() {
			mu.Unlock()
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Limite de requisições excedido",
			})
			return
		}
		
		mu.Unlock()
		
		c.Next()
	}
}
```

## Encadeamento de Middlewares

Você pode aplicar múltiplos middlewares a uma rota ou grupo de rotas:

```go
// Aplicar múltiplos middlewares a um grupo
api := router.Group("/api")
api.Use(
	logmiddleware.LogMiddleware(),
	corsmiddleware.CORSMiddleware(),
	recoverymiddleware.RecoveryMiddleware(),
)
```

A ordem de aplicação dos middlewares é importante. Os middlewares são executados na ordem em que são registrados.

## Acessando Dados do Middleware no Handler

Os middlewares podem armazenar dados no contexto do Gin para que os handlers possam acessá-los:

```go
// No middleware
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ...
		
		// Armazenar dados no contexto
		c.Set("authUser", userID)
		c.Set("userRole", role)
		
		c.Next()
	}
}

// No handler
func ProfileHandler(c *gin.Context) {
	// Acessar dados do contexto
	userID, _ := c.Get("authUser")
	role, _ := c.Get("userRole")
	
	// ...
}
```

## Interrompendo a Execução

Se um middleware precisar interromper a execução da cadeia (por exemplo, se a autenticação falhar), use `c.Abort()` ou `c.AbortWithStatusJSON()`:

```go
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ...
		
		if !autenticado {
			// Interrompe a execução e retorna uma resposta
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Não autorizado",
			})
			return
		}
		
		// Continua a execução
		c.Next()
	}
}
```

## Boas Práticas

1. **Separação de Responsabilidades**: Cada middleware deve ter uma única responsabilidade.
2. **Ordem de Execução**: Considere cuidadosamente a ordem de aplicação dos middlewares.
3. **Desempenho**: Evite operações pesadas em middlewares globais.
4. **Tratamento de Erros**: Implemente tratamento adequado de erros nos middlewares.
5. **Logging**: Adicione logs adequados para facilitar o diagnóstico de problemas.
6. **Segurança**: Esteja atento a questões de segurança ao implementar middlewares.
7. **Testes**: Escreva testes para seus middlewares para garantir que eles funcionem corretamente.
