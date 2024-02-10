package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"go-test/db"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AssignmentInfoHandler(c *gin.Context) {
	dbName := os.Getenv("DB_NAME")
	prbCol := os.Getenv("PROBLEMS_COLLECTION")

	client, exists := c.Get("mongoClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database client not available"})
		return
	}
	dbClient := client.(*mongo.Client)

	collection := dbClient.Database(dbName).Collection(prbCol)
	id := c.Param("id")
	fmt.Printf("%v\n", id)
	pid, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalf("Failed to parse problemId as a number: %v\n", pid)
		c.JSON(400, db.NewConnectionErr(err.Error()))
	}

	filter := bson.M{"problemId": pid}

	var resp ProblemResp
	err = collection.FindOne(context.TODO(), filter).Decode(&resp)

	if err != nil {
		log.Fatalf("Failed to find single result from DB: %v\n", err)
		c.JSON(400, NewFindOneAssignmentErr(err.Error()))
	}
	fmt.Printf("%v\n", resp)

	c.JSON(200, resp)
}
