package handlers

import (
	"fmt"
	"net/http"
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
	}
	client := client_.(*mongo.Client)
	pId, err := strconv.Atoi(c.Param("problemId"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to parse problemId as a numbe : %s", err.Error()))
		return
	}
	p, err := problems.SearchOneProblemWithId(client, int32(pId))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to find single result : %s", err.Error()))
		return
	}
	pwt, err := problems.ReadTestcaseContent(p)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to read testcase file : %s", err.Error()))
		return
	}
	c.JSON(http.StatusOK, pwt)
}
