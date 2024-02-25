package users

import (
	"context"
	"go-test/db"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

// 基本的にusers collectionからは1つのdocumentを問い合わせる想定
func SearchUser(client *mongo.Client, userInfo User) User {
	dbName := os.Getenv("DB_NAME")
	usrCol := os.Getenv("USERS_COLLECTION")
	collection := client.Database(dbName).Collection(usrCol)

	filter := db.MakeFilter(userInfo)

	var user User
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Printf("Failed to find single result from DB: %v\n", err.Error())
		return User{}
	}
	return user
}
