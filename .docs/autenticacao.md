# Sistema de Autenticação

## Visão Geral

O sistema de autenticação do Go Starter Kit utiliza tokens JWT (JSON Web Tokens) para autenticar usuários e proteger rotas. Ele permite que você controle o acesso a determinadas partes da aplicação, garantindo que apenas usuários autorizados possam acessá-las.

## Componentes

O sistema de autenticação é composto por:

1. **Helpers de Autenticação**: Funções auxiliares para geração e validação de tokens JWT
2. **Middleware de Autenticação**: Middleware para proteger rotas
3. **Handlers de Autenticação**: Endpoints para login e acesso a informações do usuário autenticado

## Estrutura

```
helpers/
└── authHelper/            # Funções auxiliares para autenticação
    └── authHelper.go

internal/
└── middleware/
    └── authMiddleware/    # Middleware de autenticação
        └── authMidleware.go

internal/http/handler/
└── authHandler/          # Handlers de autenticação
    └── authHandler.go
```

## Como Funciona

### 1. Login e Geração de Token

Quando um usuário faz login, o sistema:
1. Verifica as credenciais do usuário (email/senha)
2. Se as credenciais forem válidas, gera um token JWT contendo o ID do usuário e outras informações relevantes
3. Retorna o token para o cliente

### 2. Autenticação de Requisições

Para requisições a rotas protegidas:
1. O cliente envia o token JWT no cabeçalho `Authorization` (formato: `Bearer <token>`)
2. O middleware de autenticação valida o token
3. Se o token for válido, a requisição continua e o ID do usuário é disponibilizado para os handlers
4. Se o token for inválido ou estiver ausente, a requisição é rejeitada com um erro 401 (Não Autorizado)

## Como Usar

### 1. Implementar Login

```go
package authhandler

import (
	"sixTask/helpers/authHelper"
	"sixTask/internal/database"
	"context"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Buscar usuário pelo email
	db := c.MustGet("db").(*database.Queries)
	user, err := db.FindByEmail(context.Background(), request.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Credenciais inválidas",
		})
		return
	}

	// Verificar senha
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Credenciais inválidas",
		})
		return
	}

	// Gerar token JWT
	token, err := authHelper.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao gerar token",
		})
		return
	}

	// Retornar token
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
```

### 2. Proteger Rotas

Para proteger rotas, use o middleware de autenticação:

```go
// routes/api.go
func SetupRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		// Rotas públicas
		api.POST("/login", authhandler.Login)
		api.GET("/public", publicHandler)

		// Rotas protegidas
		authenticated := api.Group("/")
		authenticated.Use(authmiddleware.AuthMiddleware())
		{
			authenticated.GET("/profile", authhandler.Profile)
			authenticated.POST("/logout", authhandler.Logout)
			authenticated.GET("/protected", protectedHandler)
		}
	}

	return router
}
```

### 3. Acessar Dados do Usuário Autenticado

Nos handlers de rotas protegidas, você pode acessar o ID do usuário autenticado:

```go
func Profile(c *gin.Context) {
	// Obter ID do usuário do contexto
	userID, _ := c.Get("authUser")

	// Buscar dados do usuário
	db := c.MustGet("db").(*database.Queries)
	user, err := db.FindById(context.Background(), userID.(int64))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Usuário não encontrado",
		})
		return
	}

	// Retornar dados do usuário
	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
```

## Configuração do JWT

O sistema JWT é configurado no arquivo `helpers/authHelper/authHelper.go`:

```go
package authhelper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims define as informações armazenadas no token JWT
type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

// GetSecret retorna a chave secreta para assinatura de tokens
func GetSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

// GenerateToken gera um novo token JWT para o usuário
func GenerateToken(userID int64) (string, error) {
	// Definir tempo de expiração
	expirationTime := time.Now().Add(24 * time.Hour)

	// Criar claims
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	// Criar token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Assinar token
	tokenString, err := token.SignedString(GetSecret())
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
```

## Segurança

Para garantir a segurança do sistema de autenticação:

1. **Chave Secreta**: Use uma chave secreta forte e armazene-a em variáveis de ambiente
2. **HTTPS**: Use HTTPS em produção para proteger os tokens durante a transmissão
3. **Tempo de Expiração**: Configure um tempo de expiração adequado para os tokens
4. **Renovação de Tokens**: Implemente um mecanismo de renovação de tokens para sessões longas
5. **Revogação de Tokens**: Considere implementar um mecanismo de revogação de tokens para casos de logout ou comprometimento

## Exemplo de Uso no Cliente

### Armazenar o Token

```javascript
// Exemplo em JavaScript
async function login(email, password) {
  const response = await fetch('/api/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ email, password }),
  });

  const data = await response.json();
  
  if (response.ok) {
    // Armazenar o token no localStorage
    localStorage.setItem('token', data.token);
    return data.user;
  } else {
    throw new Error(data.error);
  }
}
```

### Usar o Token em Requisições

```javascript
// Exemplo em JavaScript
async function fetchProtectedData() {
  const token = localStorage.getItem('token');
  
  if (!token) {
    throw new Error('Não autenticado');
  }

  const response = await fetch('/api/protected', {
    headers: {
      'Authorization': `Bearer ${token}`,
    },
  });

  if (response.status === 401) {
    // Token inválido ou expirado
    localStorage.removeItem('token');
    throw new Error('Sessão expirada');
  }

  return await response.json();
}
```

## Boas Práticas

1. **Hash de Senhas**: Sempre armazene senhas com hash (usando bcrypt ou similar)
2. **Validação de Entrada**: Valide cuidadosamente os dados de entrada nos endpoints de autenticação
3. **Limites de Tentativas**: Implemente limites de tentativas de login para prevenir ataques de força bruta
4. **Logs de Autenticação**: Registre tentativas de login (bem-sucedidas e falhas) para auditoria
5. **Tokens Curtos**: Use tokens com tempo de vida curto para minimizar riscos
6. **Refresh Tokens**: Para sessões longas, considere usar refresh tokens
7. **Segurança em Camadas**: Não confie apenas na autenticação JWT; implemente verificações de autorização em cada endpoint
8. **Testes**: Escreva testes para garantir que o sistema de autenticação funcione corretamente
