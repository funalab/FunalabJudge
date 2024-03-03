package submission

import (
	"context"
	"go-test/db"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func UpdateSubmissionStatus(client *mongo.Client, sId primitive.ObjectID, status string) error {
	dbName := os.Getenv("DB_NAME")
	usrCol := os.Getenv("SUBMISSION_COLLECTION")
	collection := client.Database(dbName).Collection(usrCol)
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": sId}, bson.M{"$set": bson.M{"status": status}})
	return err
}

func UpdateSubmissionResult(client *mongo.Client, sId primitive.ObjectID, tId int, status string) error {
	dbName := os.Getenv("DB_NAME")
	subCol := os.Getenv("SUBMISSION_COLLECTION")
	collection := client.Database(dbName).Collection(subCol)

	filter := bson.M{"_id": sId}
	update := bson.M{
		"$set": bson.M{
			"results.$[elem].status": status,
		},
	}
	arrayFilters := options.ArrayFilters{
		Filters: []interface{}{bson.M{"elem.testCaseId": tId}},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update, &options.UpdateOptions{
		ArrayFilters: &arrayFilters,
	})
	return err
}
