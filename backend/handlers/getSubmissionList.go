package handlers

import (
	"fmt"
	"go-test/db/problems"
	"go-test/db/submission"
	"go-test/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetSubmissionListHandler(c *gin.Context) {
	client_, exists := c.Get("mongoClient")
	if !exists {
		util.ResponseDBNotFoundError(c)
	}
	client := client_.(*mongo.Client)

	userName := c.Param("userName")
	sList, err := submission.SearchSubmissions(client, submission.Submission{UserName: userName})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to find submission list : %s", err.Error()))
		return
	}
	var ssWithPname []submission.SubmissionWithProblemName
	for _, s := range sList {
		pList, _ := problems.SearchProblems(client, problems.Problem{Id: s.ProblemId})
		ssWithPname = append(ssWithPname, submission.SubmissionWithProblemName{Submission: s, ProblemName: pList[0].Name})
	}
	c.JSON(http.StatusOK, ssWithPname)
}
