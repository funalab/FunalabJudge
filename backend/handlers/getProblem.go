package handlers

import (
	"log"
	"strconv"

	"go-test/db/problems"
	"go-test/util"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetProblemHandler(c *gin.Context) {
	client_, exists := c.Get("mongoClient")
	if !exists {
		util.ResponseDBNotFoundError(c)
		return
	}
	client := client_.(*mongo.Client)
	pId, err := strconv.Atoi(c.Param("problemId"))
	if err != nil {
		log.Fatalf("Failed to parse problemId as a number: %v\n", pId)
		c.JSON(400, util.NewMongoConnectionErr(err.Error()))
	}
	p, _ := problems.SearchOneProblemWithId(client, int32(pId))
	pwt := problems.ReadTestcaseContent(p)
	c.JSON(200, pwt)
}
