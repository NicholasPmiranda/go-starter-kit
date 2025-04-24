package authhandler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	authhelper "sixTask/helpers/authHelper"
	"sixTask/internal/database"
)

// make the Login with Password and email
func Login(c *gin.Context) {

	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbConn, ctx := database.ConnectDB()
	defer dbConn.Close(ctx)

	query := database.New(dbConn)

	user, err := query.FindByEmail(ctx, input.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passwordCheck := authhelper.CheckPasswordHash(input.Password, user.Password)

	if !passwordCheck {
		c.JSON(http.StatusBadRequest, gin.H{"error": "usuario e ou senha incorretos"})
		return
	}

	claims := &authhelper.Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := authhelper.GetSecret()
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"token": tokenString,
	})
}

func Profile(c *gin.Context) {

	user, _ := c.Get("authUser")

	c.JSON(200, gin.H{
		"user": user,
	})
}
