package storageProvider

import (
	"crypto/rand"
	"encoding/hex"

	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const (
	StorageBasePath = "storage/app"
)

// SaveFile saves a file to storage/app directory with subdirectory path if provided
// Returns the random filename and path where file is stored
func SaveFile(c *gin.Context, fileField string, subPath string) (string, string, error) {
	// Get file from form
	file, err := c.FormFile(fileField)
	if err != nil {
		return "", "", err
	}

	// Generate random filename
	randomName, err := generateRandomFilename(filepath.Ext(file.Filename))
	if err != nil {
		return "", "", err
	}

	// Create full path with subdirectories
	fullPath := filepath.Join(StorageBasePath, subPath)

	// Create directories if they don't exist
	if err := os.MkdirAll(fullPath, 0755); err != nil {
		return "", "", err
	}

	// Full path with filename
	fullFilePath := filepath.Join(fullPath, randomName)

	// Save the file
	if err := c.SaveUploadedFile(file, fullFilePath); err != nil {
		return "", "", err
	}

	// Return random filename and relative path
	relativePath := filepath.Join(subPath, randomName)
	return fullFilePath, relativePath, nil
}

// GenerateRandomFilename creates a random filename with the original extension
func generateRandomFilename(extension string) (string, error) {
	// Generate 16 random bytes
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	// Convert to hex string and append original extension
	return hex.EncodeToString(bytes) + extension, nil
}
