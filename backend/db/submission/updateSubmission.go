package submission

import (
	"context"
	"go-test/db"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateSubmission(client *mongo.Client, searchField Submission, updateField Submission) error {
	dbName := os.Getenv("DB_NAME")
	subCol := os.Getenv("SUBMISSION_COLLECTION")
	collection := client.Database(dbName).Collection(subCol)

	sFilter := db.MakeFilter(searchField)
	uFilter := bson.M{"$set": db.MakeFilter(updateField)}

	_, err := collection.UpdateOne(context.TODO(), sFilter, uFilter)
	return err
}
