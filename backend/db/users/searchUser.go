package users

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SearchOneUserWithUserName(client *mongo.Client, userName string) (User, error) {
	dbName := os.Getenv("DB_NAME")
	usrCol := os.Getenv("USERS_COLLECTION")
	collection := client.Database(dbName).Collection(usrCol)

	filter := bson.M{"userName": userName}

	var u User
	err := collection.FindOne(context.TODO(), filter).Decode(&u)
	return u, err
}
