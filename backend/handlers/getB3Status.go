package handlers

import (
	"errors"
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
	fmt.Printf("handler\n")
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
		c.AbortWithError(http.StatusInternalServerError, errors.Join(errors.New("failed to find single result"), err))
	}
	ps, err := problems.SearchProblems(client, problems.Problem{})
	var usst []users.UserStatus
	for _, u := range us {
		var rs []submission.Result
		for _, p := range ps {
			us := submission.Submission{UserName: u, ProblemId: p.Id}
			ss, err := submission.SearchSubmissions(client, us)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, errors.Join(errors.New("failed to find single result"), err))
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
			rs = append(rs, submission.Result{TestId: int(p.Id), Status: pst})

		}
		usst = append(usst, users.UserStatus{UserName: u, ProblemsStatus: rs})
	}

	c.JSON(http.StatusOK, usst)
}
