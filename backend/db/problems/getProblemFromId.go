package problems

import (
	"context"
	"go-test/myTypes"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetProblemFromId(c *gin.Context, problemId int32) myTypes.Problem {
	client, exists := c.Get("mongoClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "DB client is not available."})
		return myTypes.Problem{}
	}
	dbName := os.Getenv("DB_NAME")
	prbCol := os.Getenv("PROBLEMS_COLLECTION")
	collection := (client.(*mongo.Client)).Database(dbName).Collection(prbCol)

	filter := bson.M{"problemId": problemId}

	var problem myTypes.Problem
	err := collection.FindOne(context.TODO(), filter).Decode(&problem)
	if err != nil {
		log.Printf("Failed to find single result from DB: %v\n", err.Error())
		return myTypes.Problem{}
	}
	return problem
}
