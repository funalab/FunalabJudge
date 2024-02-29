package submission

import (
	"context"
	"go-test/db"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SearchOneSubmissionWithId(client *mongo.Client, sId primitive.ObjectID) (Submission, error) {
	dbName := os.Getenv("DB_NAME")
	subCol := os.Getenv("SUBMISSION_COLLECTION")
	collection := client.Database(dbName).Collection(subCol)

	filter := bson.M{"_id": sId}

	var s Submission
	err := collection.FindOne(context.TODO(), filter).Decode(&s)
	return s, err
}

func SearchSubmissions(client *mongo.Client, searchField Submission) ([]Submission, error) {
	dbName := os.Getenv("DB_NAME")
	subCol := os.Getenv("SUBMISSION_COLLECTION")
	collection := client.Database(dbName).Collection(subCol)

	sFilter := db.MakeFilter(searchField)

	cursor, err := collection.Find(context.TODO(), sFilter)
	if err != nil {
		return []Submission{}, err
	}
	var submissions []Submission
	if err = cursor.All(context.TODO(), &submissions); err != nil {
		return []Submission{}, err
	}
	return submissions, nil
}
