package main

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type JsonRequest struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
}

type JsonReturn struct {
	Authorized bool   `json:"authorized"`
	UserName   string `json:"userName"`
	Role       string `json:"role"`
}

type User struct {
	UserId      int64
	Email       string
	Password    string
	CreatedDate string
	Role        string
}

func extractUsername(email string) string {
	parts := strings.Split(email, "@")
	username := parts[0]
	return username
}

func authorizeUser(user User, form JsonRequest) bool {
	if user.Password == form.Password {
		return true
	} else {
		return false
	}
}

func authUser(c *gin.Context) {
	var jsonRequest JsonRequest
	if err := c.ShouldBindJSON(&jsonRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var result User
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	client, _ := mongo.Connect(
		ctx,
		options.Client().ApplyURI("mongodb://localhost:27017/"),
	)
	err := client.Database("dev").Collection("users").FindOne(context.TODO(), bson.D{{"email", jsonRequest.UserId}}).Decode(&result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	jsonReturn := JsonReturn{
		Authorized: authorizeUser(result, jsonRequest),
		UserName:   extractUsername(result.Email),
		Role:       result.Role,
	}
	c.JSON(http.StatusOK, jsonReturn)
}
