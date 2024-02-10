package user

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserFromUserId(c *gin.Context, userId string) *User {
	client, exists := c.Get("mongoClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "DB client is not available."})
		return &User{}
	}
	dbName := os.Getenv("DB_NAME")
	collection := (client.(*mongo.Client)).Database(dbName).Collection("USERS_COLLECTION")

	filter := bson.M{"userId": userId}

	var user User
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Printf("Failed to find single result from DB: %v\n", err.Error())
		return &User{}
	}
	return &user
}
