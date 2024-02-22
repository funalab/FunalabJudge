package auth

import (
	"go-test/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChangePassRequest struct {
	UserName string `json:"userName"`
	ExPass   string `json:"exPass"`
	NewPass  string `json:"newPass"`
}

func ChangeUserPass(c *gin.Context) {
	var jsonRequest ChangePassRequest

	if err := c.ShouldBind(&jsonRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// targUser := myTypes.User{UserName: jsonRequest.UserName}
	// userInfo := user.SearchUser(c, targUser)
	userInfo := user.GetUserFromUserName(c, jsonRequest.UserName)
	if !CheckPasswordHash(jsonRequest.ExPass, userInfo.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "Password did not match"})
		return
	}

	hash, _ := HashPassword(jsonRequest.NewPass)
	updated := user.UpdateUserPass(c, userInfo.UserName, hash)
	if updated {
		c.JSON(http.StatusOK, gin.H{"status": "Password updated successfully!"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
	}
}
