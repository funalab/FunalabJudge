package users

import (
	"context"
	"go-test/db"
	"go-test/myTypes"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserFromUserName(c *gin.Context, userName string) *myTypes.User {
	client, exists := c.Get("mongoClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "DB client is not available."})
		return &myTypes.User{}
	}
	dbName := os.Getenv("DB_NAME")
	usrCol := os.Getenv("USERS_COLLECTION")
	collection := (client.(*mongo.Client)).Database(dbName).Collection(usrCol)

	filter := bson.M{"userName": userName}

	var user myTypes.User
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Printf("Failed to find single result from DB: %v\n", err.Error())
		return &myTypes.User{}
	}
	return &user
}

func GetUserFromUserId(c *gin.Context, userId int32) *myTypes.User {
	client, exists := c.Get("mongoClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "DB client is not available."})
		return &myTypes.User{}
	}
	dbName := os.Getenv("DB_NAME")
	usrCol := os.Getenv("USERS_COLLECTION")
	collection := (client.(*mongo.Client)).Database(dbName).Collection(usrCol)

	filter := bson.M{"userId": userId}

	var user myTypes.User
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Printf("Failed to find single result from DB: %v\n", err.Error())
		return &myTypes.User{}
	}
	return &user
}

// MakeFilterWithNonnilField関数が未完成なので動かない
func SearchUser(c *gin.Context, userInfo myTypes.User) *myTypes.User {
	client, exists := c.Get("mongoClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "DB client is not available."})
		return &myTypes.User{}
	}
	dbName := os.Getenv("DB_NAME")
	usrCol := os.Getenv("USERS_COLLECTION")
	collection := (client.(*mongo.Client)).Database(dbName).Collection(usrCol)

	filter := db.MakeFilterWithNonnilField(userInfo)

	var user myTypes.User
	// 基本的にusers collectionからは1つのdocumentを問い合わせる想定
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Printf("Failed to find single result from DB: %v\n", err.Error())
		return &myTypes.User{}
	}
	return &user
}
