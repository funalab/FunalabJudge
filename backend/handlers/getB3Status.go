package handlers

import (
	"fmt"
	"net/http"
	"time"

	"go-test/db/problems"
	"go-test/db/submission"
	"go-test/db/users"
	"go-test/util"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetB3StatusHandler(c *gin.Context) {
	client_, exists := c.Get("mongoClient")
	if !exists {
		util.ResponseDBNotFoundError(c)
	}
	client := client_.(*mongo.Client)

	now := time.Now()
	year := now.Year()
	stofYear := time.Date(year, 1, 1, 0, 0, 0, 0, now.Location())
	edofYear := time.Date(year, 12, 31, 23, 59, 59, 999, now.Location())
	us, err := users.SearchUsersWithJoinnedDate(client, stofYear, edofYear)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to find single result : %s", err.Error()))
		return
	}
	ps, _ := problems.SearchProblems(client, problems.Problem{})
	var usst []users.UserStatus
	for _, u := range us {
		var rs []problems.ProblemWithStatus
		for _, p := range ps {
			ups := submission.Submission{UserName: u, ProblemId: p.Id}
			ss, err := submission.SearchSubmissions(client, ups)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to find single result : %s", err.Error()))
				return
			}
			pst := "NS"
			for _, s := range ss {
				st := s.Status
				if st == "AC" {
					pst = "AC"
					break
				} else if st == "WA" {
					pst = "WA"
				} else if st == "TLE" && (pst != "WA") {
					pst = "TLE"
				} else if st == "RE" && (pst != "WA" && pst != "TLE") {
					pst = "RE"
				} else if st == "CE" && (pst != "WA" && pst != "TLE" && pst != "RE") {
					pst = "CE"
				}
			}
			pn, _ := problems.SearchProblems(client, problems.Problem{Id: p.Id})
			rs = append(rs, problems.ProblemWithStatus{ProblemId: int(p.Id), ProblemName: pn[0].Name, Status: pst})

		}
		usst = append(usst, users.UserStatus{UserName: u, ProblemsStatus: rs})
	}

	c.JSON(http.StatusOK, usst)
}
