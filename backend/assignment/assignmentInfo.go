package assignment

import (
	"context"
	"log"
	"os"
	"strconv"

	"go-test/db"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func AssignmentInfoHandler(c *gin.Context) {
	dbName := os.Getenv("DB_NAME")
	prbCol := os.Getenv("PROBLEMS_COLLECTION")

	err, client := db.Mongo_connectable()
	if err != nil {
		log.Fatalf("connection error :%v", err.Error())
	}

	collection := client.Database(dbName).Collection(prbCol)
	pid, err := strconv.Atoi(c.Param("id"))
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
	c.JSON(200, resp)
}
