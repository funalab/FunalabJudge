package handlers

import (
	"log"

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
		return
	}
	client := client_.(*mongo.Client)

	submissionId := c.Param("submissionId")
	sId, err := primitive.ObjectIDFromHex(submissionId)
	if err != nil {
		log.Printf("Failed to parse objectId from hex: %v \n", err.Error())
		c.JSON(400, err.Error())
	}
	s, err := submission.SearchOneSubmissionWithId(client, sId)
	if err != nil {
		log.Printf("Failed to find single result from DB: %v \n", err.Error())
		c.JSON(400, err.Error())
	}
	c.JSON(200, s)
}
