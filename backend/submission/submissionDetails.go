package submission

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SubmissionHandler(c *gin.Context) {
	submitId := c.Param("submitId")
	client, exists := c.Get("mongoClient")
	fmt.Printf("%v\n", submitId)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "DB client is not availale."})
	}
	dbName := os.Getenv("DB_NAME")
	submitCol := os.Getenv("SUBMISSION_COLLECTION")
	collection := (client.(*mongo.Client)).Database(dbName).Collection(submitCol)

	id, err := strconv.Atoi(submitId)
	filter := bson.D{{Key: "id", Value: id}}

	var submission Submission
	err = collection.FindOne(context.TODO(), filter).Decode(&submission)
	if err != nil {
		log.Printf("Failed to find single result from DB: %v \n", err.Error())
		c.JSON(400, err.Error())
	}

	if err != nil {
		c.JSON(400, err.Error())
	}
	fmt.Printf("%v\n", submission)
	c.JSON(200, submission)
}