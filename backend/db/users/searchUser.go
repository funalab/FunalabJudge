package users

import (
	"context"
	"go-test/db"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// MakeFilterWithNonnilField関数が未完成なので動かない
func SearchUser(c *gin.Context, userInfo User) *User {
	client, exists := c.Get("mongoClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "DB client is not available."})
		return &User{}
	}
	dbName := os.Getenv("DB_NAME")
	usrCol := os.Getenv("USERS_COLLECTION")
	collection := (client.(*mongo.Client)).Database(dbName).Collection(usrCol)

	filter := db.MakeFilterWithNonnilField(userInfo)

	var user User
	// 基本的にusers collectionからは1つのdocumentを問い合わせる想定
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Printf("Failed to find single result from DB: %v\n", err.Error())
		return &User{}
	}
	return &user
}
