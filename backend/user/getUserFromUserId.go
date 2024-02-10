package user

import (
	"context"
	"go-test/db"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
)

func GetUserFromUserId(userId string) *User {
	err, client := db.Mongo_connectable()
	if err != nil {
		log.Printf("connection error :%v\n", err.Error())
		return &User{}
	}
	dbName := os.Getenv("DB_NAME")
	collection := client.Database(dbName).Collection("USERS_COLLECTION")

	filter := bson.M{"userId": userId}

	var user User
	err = collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Printf("Failed to find single result from DB: %v\n", err.Error())
		return &User{}
	}
	return &user
}
