package submission

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

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
	collection := (client.(*mongo.Client)).Database(dbName).Collection("USERS_COLLECTION")

	filter := bson.M{"userId": userId}
	var submissions []Submission
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Printf("Failed to find single result from DB: %v\n", err.Error())
		return &[]Submission{}, err
	}

	if err = cursor.All(context.TODO(), &submissions); err != nil {
		log.Printf("Failed to find single result from DB: %v\n", err.Error())
		return &[]Submission{}, err
	}
	return &submissions, nil
}
