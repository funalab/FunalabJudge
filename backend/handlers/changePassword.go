package handlers

import (
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
		return
	}
	client := client_.(*mongo.Client)

	var jsonRequest ChangePassRequest
	if err := c.ShouldBind(&jsonRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	searchField := users.User{UserName: jsonRequest.UserName}
	u := users.SearchUser(client, searchField)
	println(u.Password)
	if !auth.CheckPasswordHash(jsonRequest.ExPass, u.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "Password did not match"})
		return
	}

	hash, _ := auth.HashPassword(jsonRequest.NewPass)
	updateField := users.User{Password: hash}

	err := users.UpdateUser(client, u, updateField)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "Password updated successfully!"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
	}
}
