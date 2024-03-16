package handlers

import (
	"fmt"
	"go-test/db/submission"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func GetSubmittedFilesHandler(c *gin.Context) {
	sid := c.Param("submissionId")
	compileResourcePath := filepath.Join(os.Getenv("EXEC_DIR"), sid)

	fs, err := os.ReadDir(compileResourcePath)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get files from directory path : %s", err.Error()))
		return
	}

	files := make([]submission.SubmittedFile, 0)

	for _, f := range fs {
		content, err := os.ReadFile(filepath.Join(compileResourcePath, f.Name()))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to read file content : %s", err.Error()))
			return
		}
		var sf submission.SubmittedFile
		sf.Name = f.Name()
		sf.Content = string(content)
		files = append(files, sf)
	}

	c.JSON(http.StatusOK, files)
}
