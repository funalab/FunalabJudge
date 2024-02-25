package handlers

import (
	"go-test/db/submission"
	"go-test/db/users"

	"github.com/gin-gonic/gin"
)

func GetSubmissionListHandler(c *gin.Context) {
	userName := c.Param("userName")
	userData := users.GetUserFromUserName(c, userName)
	submissions, err := submission.GetSubmissionsFromUser(c, *userData)
	if err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, submissions)
}
