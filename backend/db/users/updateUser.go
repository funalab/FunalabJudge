package users

import (
	"context"
	"go-test/db"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateUser(client *mongo.Client, searchField User, updateField User) error {
	dbName := os.Getenv("DB_NAME")
	usrCol := os.Getenv("USERS_COLLECTION")
	collection := client.Database(dbName).Collection(usrCol)

	sFilter := db.MakeFilter(searchField)
	uFilter := db.MakeFilter(updateField)

	_, err := collection.UpdateOne(context.TODO(), sFilter, bson.M{"$set": uFilter})
	// c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
	return err
}
