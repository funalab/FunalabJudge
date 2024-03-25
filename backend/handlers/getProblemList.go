package handlers

import (
	"fmt"
	"net/http"

	"go-test/db/problems"
	"go-test/db/submission"
	"go-test/util"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type problemWithStatusForDashboard struct {
	Problem problems.Problem
	Status  bool
}

func GetProblemListHandler(c *gin.Context) {
	client_, exists := c.Get("mongoClient")
	if !exists {
		util.ResponseDBNotFoundError(c)
	}
	client := client_.(*mongo.Client)

	pList, err := problems.SearchProblems(client, problems.Problem{})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to find problem list : %s", err.Error()))
		return
	}

	var ps []problemWithStatusForDashboard
	userName := c.Param("userName")
	for _, p := range pList {
		sList, err := submission.SearchSubmissions(client, submission.Submission{UserName: userName, ProblemId: p.Id})
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to find submission list : %s", err.Error()))
			return
		}
		statusFlag := false
		for _, s := range sList {
			if s.Status == "AC" {
				statusFlag = true
				break
			}
		}
		ps = append(ps, problemWithStatusForDashboard{Problem: p, Status: statusFlag})
	}

	c.JSON(http.StatusOK, ps)
}
