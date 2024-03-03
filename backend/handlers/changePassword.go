package handlers

import (
	"errors"
	"go-test/auth"
	"go-test/db/users"
	"go-test/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChangePassRequest struct {
	UserName string `json:"userName"`
	ExPass   string `json:"exPass"`
	NewPass  string `json:"newPass"`
}

func ChangePasswordHandler(c *gin.Context) {
	client_, exists := c.Get("mongoClient")
	if !exists {
		util.ResponseDBNotFoundError(c)
	}
	client := client_.(*mongo.Client)

	var jsonRequest ChangePassRequest
	if err := c.Bind(&jsonRequest); err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, errors.Join(errors.New("failed to handle form content"), err))
	}

	u, err := users.SearchOneUserWithUserName(client, jsonRequest.UserName)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.Join(errors.New("failed to find single result"), err))
	}

	if !auth.CheckPasswordHash(jsonRequest.ExPass, u.Password) {
		c.AbortWithError(http.StatusUnauthorized, errors.New("password did not match"))
	}

	hash, _ := auth.HashPassword(jsonRequest.NewPass)
	updateField := users.User{Password: hash}

	err = users.UpdateUserWithUserName(client, u.UserName, updateField)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.Join(errors.New("failed to update password"), err))
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "Password updated successfully!"})
	}
}
