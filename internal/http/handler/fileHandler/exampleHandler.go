package fileHandler

import (
	"boilerPlate/config/storageProvider"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadFileExample demonstrates how to use the SaveFile function
func UploadFileExample(c *gin.Context) {

	subPath := "pdfs"

	filename, relativePath, err := storageProvider.SaveFile(c, "file", subPath)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"filename":     filename,
		"path":         relativePath,
		"download_url": "/storage/" + relativePath,
	})
}
