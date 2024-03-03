package handlers

import (
	"errors"
	"go-test/db/submission"
	"go-test/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetSubmissionListHandler(c *gin.Context) {
	client_, exists := c.Get("mongoClient")
	if !exists {
		util.ResponseDBNotFoundError(c)
	}
	client := client_.(*mongo.Client)

	userName := c.Param("userName")
	submissions, err := submission.SearchSubmissions(client, submission.Submission{UserName: userName})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.Join(errors.New("failed to find submission list"), err))
	}
	c.JSON(http.StatusOK, submissions)
}
