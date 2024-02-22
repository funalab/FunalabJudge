package submission

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"go-test/myTypes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetSubmissionsFromUser(c *gin.Context, user myTypes.User) (*[]myTypes.SubmissionWithStatus, error) {
	client, exists := c.Get("mongoClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "DB client is not available."})
		return &[]myTypes.SubmissionWithStatus{}, errors.New(fmt.Sprint("Error: NotExist\n"))
	}

	dbName := os.Getenv("DB_NAME")
	submitCol := os.Getenv("SUBMISSION_COLLECTION")
	collection := (client.(*mongo.Client)).Database(dbName).Collection(submitCol)

	filter := bson.M{"userId": user.UserId}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Printf("Failed to find single result from DB: %v\n", err.Error())
		return &[]myTypes.SubmissionWithStatus{}, err
	}
	var submissions []myTypes.SubmissionWithStatus
	if err = cursor.All(context.TODO(), &submissions); err != nil {
		log.Printf("Failed to find single result from DB: %v\n", err.Error())
		return &[]myTypes.SubmissionWithStatus{}, err
	}
	for i, submission := range submissions {
		status := "AC"
		for _, object := range submission.Submission.Results {
			if object.Status == "WA" {
				status = "WA"
				break
			} else if object.Status == "TLE" {
				status = "TLE"
				break
			} else if object.Status == "WJ" {
				status = "WJ"
				break
			}
		}
		submissions[i].Status = status
	}
	return &submissions, nil
}

func GetSubmissionsFromSubmissionId(c *gin.Context, submissionId int) *myTypes.Submission {
	client, exists := c.Get("mongoClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "DB client is not available."})
		return &myTypes.Submission{}
	}
	dbName := os.Getenv("DB_NAME")
	usrCol := os.Getenv("SUBMISSION_COLLECTION")
	collection := (client.(*mongo.Client)).Database(dbName).Collection(usrCol)

	filter := bson.M{"id": submissionId}

	var submission myTypes.Submission
	err := collection.FindOne(context.TODO(), filter).Decode(&submission)
	if err != nil {
		log.Printf("Failed to find single result from DB: %v\n", err.Error())
		return &myTypes.Submission{}
	}
	return &submission
}
