package handlers

import (
	"go-test/db/problems"
	"go-test/db/submission"
	"go-test/judge"
	"go-test/util"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type submissionRequest struct {
	ProblemId int32 `form:"problemId"`
}

func AddSubmissionHandler(c *gin.Context) {
	client_, exists := c.Get("mongoClient")
	if !exists {
		util.ResponseDBNotFoundError(c)
		return
	}
	client := client_.(*mongo.Client)

	var sr submissionRequest
	if err := c.Bind(&sr); err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{"err": err.Error()})
	}
	userName := c.Param("userName")
	p, err := problems.SearchProblemWithId(client, sr.ProblemId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to find single result from DB:" + err.Error()})
		return
	}
	s, err := submission.InsertNewSubmission(client, userName, p)
	if err != nil {
		println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to insert:" + err.Error()})
		return
	}

	// save posted files
	form, _ := c.MultipartForm()
	files := form.File["files[]"]
	for _, file := range files {
		c.SaveUploadedFile(file, filepath.Join(os.Getenv("EXEC_DIR"), s.Id.Hex(), file.Filename))
	}
	// コンパイル&実行プロセスのマルチスレッド予約
	go judge.JudgeProcess(client, s)
	c.JSON(200, nil)
}
