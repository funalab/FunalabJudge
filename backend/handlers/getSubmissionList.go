package handlers

import (
	"go-test/db/submission"
	"go-test/util"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetSubmissionListHandler(c *gin.Context) {
	client_, exists := c.Get("mongoClient")
	if !exists {
		util.ResponseDBNotFoundError(c)
		return
	}
	client := client_.(*mongo.Client)

	userName := c.Param("userName")
	submissions, err := submission.SearchSubmissions(client, submission.Submission{UserName: userName})
	if err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, submissions)
}
