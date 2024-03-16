package handlers

import (
	"fmt"
	"go-test/auth"
	"go-test/db/problems"
	"go-test/db/submission"
	"go-test/db/users"
	"go-test/judge"
	"go-test/util"
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
	}
	client := client_.(*mongo.Client)

	var sr submissionRequest
	if err := c.Bind(&sr); err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, fmt.Errorf("failed to handle form content : %s", err.Error()))
		return
	}
	u := auth.GetUserFromJwt(c).(*users.User)
	userName := u.UserName
	p, err := problems.SearchOneProblemWithId(client, sr.ProblemId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to find single result : %s", err.Error()))
		return
	}
	s, err := submission.InsertNewSubmission(client, userName, p)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to insert new submission : %s", err.Error()))
		return
	}

	// save posted files
	form, _ := c.MultipartForm()
	files := form.File["files[]"]
	for _, file := range files {
		c.SaveUploadedFile(file, filepath.Join(os.Getenv("EXEC_DIR"), s.Id.Hex(), file.Filename))
	}
	// コンパイル&実行プロセスのマルチスレッド予約
	go judge.JudgeProcess(c, s)
	c.JSON(http.StatusOK, nil)
}
