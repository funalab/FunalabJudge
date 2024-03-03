package handlers

import (
	"errors"
	"net/http"

	"go-test/db/problems"
	"go-test/db/submission"
	"go-test/util"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type problemWithStatus struct {
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
		c.AbortWithError(http.StatusInternalServerError, errors.Join(errors.New("failed to find problem list"), err))
	}

	var ps []problemWithStatus
	userName := c.Param("userName")
	for _, p := range pList {
		sList, err := submission.SearchSubmissions(client, submission.Submission{UserName: userName, ProblemId: p.Id})
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, errors.Join(errors.New("failed to find submission list"), err))
		}
		statusFlag := false
		for _, s := range sList {
			if s.Status == "AC" {
				statusFlag = true
				break
			}
		}
		ps = append(ps, problemWithStatus{Problem: p, Status: statusFlag})
	}

	c.JSON(http.StatusOK, ps)
}
