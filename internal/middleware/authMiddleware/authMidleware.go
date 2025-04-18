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
	secretKey := authhelper.GetSecret() // Idealmente, isso deveria vir de uma variável de ambiente

	return func(c *gin.Context) {
		bearer := c.GetHeader("Authorization")

		if bearer == "" || !strings.HasPrefix(bearer, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			return
		}

		tokenStr := strings.Split(bearer, " ")[1]
		claims := &authhelper.Claims{}

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

		c.Set("authUser", claims.UserID)
		c.Next()
	}
}
