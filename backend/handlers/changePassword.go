package handlers

import (
	"errors"
	"fmt"
	"go-test/auth"
	"go-test/db/users"
	"go-test/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChangePassRequest struct {
	ExPass  string `json:"exPass"`
	NewPass string `json:"newPass"`
}

func ChangePasswordHandler(c *gin.Context) {
	client_, exists := c.Get("mongoClient")
	if !exists {
		util.ResponseDBNotFoundError(c)
	}
	client := client_.(*mongo.Client)

	var jsonRequest ChangePassRequest
	if err := c.Bind(&jsonRequest); err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to handle form content : %s", err.Error()))
		return
	}

	u_ := auth.GetUserFromJwt(c).(*users.User)
	u, err := users.SearchOneUserWithUserName(client, u_.UserName)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to find single result : %s", err.Error()))
		return
	}

	if !auth.CheckPasswordHash(jsonRequest.ExPass, u.Password) {
		c.AbortWithError(http.StatusBadRequest, errors.New("password did not match"))
		return
	}

	hash, _ := auth.HashPassword(jsonRequest.NewPass)
	updateField := users.User{Password: hash}

	err = users.UpdateUserWithUserName(client, u.UserName, updateField)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to update password : %s", err.Error()))
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "Password updated successfully!"})
	}
}
