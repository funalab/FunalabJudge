package submission

import (
	"github.com/gin-gonic/gin"
)

func SubmissionQueueHandler(c *gin.Context) {
	userId := c.Param("userId")
	submisions, err := GetSubmissionsFromUserId(c, userId)
	if err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, submisions)
}
