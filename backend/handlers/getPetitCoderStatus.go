package handlers

import (
	"fmt"
	"go-test/db/problems"
	"go-test/db/submission"
	"go-test/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type petitCoderStatus struct {
	ProblemId    int
	ProblemName  string
	PCSubmission []PCSubmission
}

type PCSubmission struct {
	UserName      string
	SubmittedDate time.Time
}

func GetPetitCoderStatusHandler(c *gin.Context) {
	client_, exists := c.Get("mongoClient")
	if !exists {
		util.ResponseDBNotFoundError(c)
	}
	client := client_.(*mongo.Client)

	// isPetitCoderのproblemIdかつACなsubmissionを、userNameでdistinctし、submittedDateで並び替えて返す
	tmpBool := true
	pList, err := problems.SearchProblems(client, problems.Problem{IsPetitCoder: &tmpBool})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to search problem info : %s", err.Error()))
		return
	}
	problemMap := make(map[int32]string)
	for _, problem := range pList {
		problemMap[problem.Id] = problem.Name
	}
	pcs := []petitCoderStatus{}
	for key, value := range problemMap {
		sList, err := submission.SearchPetitCoderSubmissions(client, key)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to search submission info : %s", err.Error()))
			return
		}
		pcsub := []PCSubmission{}
		for _, s := range sList {
			pcsub = append(pcsub, PCSubmission{
				UserName:      s.UserName,
				SubmittedDate: s.SubmittedDate,
			})
		}
		pcs = append(pcs, petitCoderStatus{
			ProblemId:    int(key),
			ProblemName:  value,
			PCSubmission: pcsub,
		})
	}
	c.JSON(http.StatusOK, pcs)
}
