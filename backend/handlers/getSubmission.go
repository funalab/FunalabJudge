package handlers

import (
	"errors"
	"net/http"

	"go-test/db/submission"
	"go-test/util"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetSubmissionHandler(c *gin.Context) {
	client_, exists := c.Get("mongoClient")
	if !exists {
		util.ResponseDBNotFoundError(c)
	}
	client := client_.(*mongo.Client)

	submissionId := c.Param("submissionId")
	sId, err := primitive.ObjectIDFromHex(submissionId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.Join(errors.New("failed to parse objectId from hex"), err))
	}
	s, err := submission.SearchOneSubmissionWithId(client, sId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.Join(errors.New("failed to find single result"), err))
	}
	c.JSON(http.StatusOK, s)
}
