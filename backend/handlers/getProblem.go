package handlers

import (
	"errors"
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
		c.AbortWithError(http.StatusInternalServerError, errors.Join(errors.New("failed to parse problemId as a numbe"), err))
	}
	p, err := problems.SearchOneProblemWithId(client, int32(pId))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.Join(errors.New("failed to find single result"), err))
	}
	pwt, err := problems.ReadTestcaseContent(p)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.Join(errors.New("failed to read testcase file"), err))
	}
	c.JSON(http.StatusOK, pwt)
}
