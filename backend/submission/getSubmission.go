package submission

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"

	"go-test/myTypes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetSubmissionHandler(c *gin.Context) {
	submissionId := c.Param("submissionId")
	client, exists := c.Get("mongoClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "DB client is not availale."})
	}
	dbName := os.Getenv("DB_NAME")
	submitCol := os.Getenv("SUBMISSION_COLLECTION")
	collection := (client.(*mongo.Client)).Database(dbName).Collection(submitCol)

	id, err := strconv.Atoi(submissionId)
	filter := bson.D{{Key: "id", Value: id}}

	var submission myTypes.Submission
	err = collection.FindOne(context.TODO(), filter).Decode(&submission)
	if err != nil {
		log.Printf("Failed to find single result from DB: %v \n", err.Error())
		c.JSON(400, err.Error())
	}

	if err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, submission)
}
