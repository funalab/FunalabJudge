package handlers

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"go-test/assignment"
	"go-test/util"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetProblemHandler(c *gin.Context) {
	dbName := os.Getenv("DB_NAME")
	prbCol := os.Getenv("PROBLEMS_COLLECTION")
	client, exists := c.Get("mongoClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "DB client is not available."})
		return
	}
	collection := (client.(*mongo.Client)).Database(dbName).Collection(prbCol)
	pid, err := strconv.Atoi(c.Param("problemId"))
	if err != nil {
		log.Fatalf("Failed to parse problemId as a number: %v\n", pid)
		c.JSON(400, util.NewMongoConnectionErr(err.Error()))
	}
	resp := assignment.TranslatePathIntoProblemResp(collection, pid)
	if resp == nil {
		log.Fatalf("Failed to find single result from DB: %v\n", err)
		c.JSON(400, util.NewFindOneAssignmentErr(err.Error()))
	}
	c.JSON(200, resp)
}
