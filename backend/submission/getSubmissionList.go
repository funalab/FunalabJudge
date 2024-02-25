package submission

import (
	"go-test/user"

	"github.com/gin-gonic/gin"
)

func GetSubmissionListHandler(c *gin.Context) {
	userName := c.Param("userName")
	userData := user.GetUserFromUserName(c, userName)
	submissions, err := GetSubmissionsFromUser(c, *userData)
	if err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, submissions)
}
