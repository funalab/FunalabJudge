package users

import (
	"context"
	"go-test/db"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func SearchUser(client *mongo.Client, searchField User) User {
	dbName := os.Getenv("DB_NAME")
	usrCol := os.Getenv("USERS_COLLECTION")
	collection := client.Database(dbName).Collection(usrCol)

	filter := db.MakeFilter(searchField)

	var u User
	err := collection.FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		log.Printf("Failed to find single result from DB: %v\n", err.Error())
		return User{}
	}
	return u
}
