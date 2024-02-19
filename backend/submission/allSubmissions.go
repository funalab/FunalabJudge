package submission

import (
	"context"
	"go-test/math"
	"go-test/types"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func MaxSubmissionIdHandler(c *gin.Context) {
	dbName := os.Getenv("DB_NAME")
	subCol := os.Getenv("SUBMISSION_COLLECTION")
	client, exists := c.Get("mongoClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "DB client is not available."})
		return
	}
	collection := (client.(*mongo.Client)).Database(dbName).Collection(subCol)
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("Failed to find all submission information.")
		c.JSON(400, gin.H{"err": err.Error()})
	}
	var submissions []types.Submission
	if err = cur.All(context.TODO(), &submissions); err != nil {
		log.Println("Failed to fetch all submission information.")
		c.JSON(400, gin.H{"err": err.Error()})
	}
	maxSubmissionId := -1
	for _, submission := range submissions {
		submissionId := int(submission.Id)
		maxSubmissionId = math.Max(submissionId, maxSubmissionId)
	}
	c.JSON(200, gin.H{"maxSubmissionId": maxSubmissionId})
}
