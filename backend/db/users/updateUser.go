package users

import (
	"context"
	"go-test/db"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateUserWithUserName(client *mongo.Client, userName string, updateField User) error {
	dbName := os.Getenv("DB_NAME")
	usrCol := os.Getenv("USERS_COLLECTION")
	collection := client.Database(dbName).Collection(usrCol)

	sFilter := bson.M{"userName": userName}
	uFilter := bson.M{"$set": db.MakeFilter(updateField)}

	_, err := collection.UpdateOne(context.TODO(), sFilter, uFilter)
	return err
}
