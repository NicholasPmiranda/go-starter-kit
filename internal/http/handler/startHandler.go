package handler

import (
	"github.com/gin-gonic/gin"

	"boilerPlate/internal/database"
)

func StartHandler(c *gin.Context) {

	dbConn, ctx := database.ConnectDB()
	defer dbConn.Close(ctx)

	queries := database.New(dbConn)

	users, _ := queries.FindMany(ctx)

	c.JSON(200, gin.H{
		"users": users,
	})
}
