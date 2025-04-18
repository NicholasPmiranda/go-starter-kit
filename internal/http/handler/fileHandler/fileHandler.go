package fileHandler

import (
	"boilerPlate/config/storageProvider"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// GetFileHandler returns a handler function to serve files from storage
func GetFileHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		filePath := c.Param("filepath")
		basePath := storageProvider.StorageBasePath
		fullPath := filepath.Join(basePath, filePath)

		// Check if file exists
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		// Serve the file
		c.File(fullPath)
	}
}
