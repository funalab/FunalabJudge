package submission

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetSubmissionsFromUserId(c *gin.Context, userId string) (*[]Submission, error) {
	client, exists := c.Get("mongoClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "DB client is not available."})
		return &[]Submission{}, errors.New(fmt.Sprint("Error: NotExist\n"))
	}

	dbName := os.Getenv("DB_NAME")
	submitCol := os.Getenv("SUBMISSION_COLLECTION")
	collection := (client.(*mongo.Client)).Database(dbName).Collection(submitCol)

	id, err := strconv.Atoi(userId)

	filter := bson.D{{Key: "userId", Value: id}}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Printf("Failed to find single result from DB: %v\n", err.Error())
		return &[]Submission{}, err
	}
	var submissions []Submission
	if err = cursor.All(context.TODO(), &submissions); err != nil {
		log.Printf("Failed to find single result from DB: %v\n", err.Error())
		return &[]Submission{}, err
	}
	for i, submission := range submissions {
		status := "AC"
		for _, object := range submission.Results {
			if object.Status == "WA" {
				status = "WA"
				break
			} else if object.Status == "TLE" {
				status = "TLE"
				break
			}
		}
		submissions[i].Status = status
	}
	return &submissions, nil
}
