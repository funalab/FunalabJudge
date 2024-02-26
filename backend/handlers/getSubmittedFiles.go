package handlers

import (
	"go-test/db/submission"
	"go-test/util"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func GetSubmittedFilesHandler(c *gin.Context) {
	projectRoot := util.GetProjectRoot()
	sid := c.Param("submissionId")
	compileResourcePath := filepath.Join(projectRoot, "compile_resource", sid)

	fs, err := os.ReadDir(compileResourcePath)
	if err != nil {
		log.Printf("Failed to get files from directory path: %v\n", err.Error())
		c.JSON(http.StatusBadRequest, err)
	}

	files := make([]submission.SubmittedFile, 0)

	for _, f := range fs {
		content, err := os.ReadFile(filepath.Join(compileResourcePath, f.Name()))
		if err != nil {
			log.Printf("Failed to read file content: %v\n", err.Error())
			c.JSON(http.StatusBadRequest, err)
		}
		var sf submission.SubmittedFile
		sf.Name = f.Name()
		sf.Content = string(content)
		files = append(files, sf)
	}

	c.JSON(http.StatusOK, files)
}
