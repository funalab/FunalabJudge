package submission

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SearchSubmissionWithId(client *mongo.Client, sId primitive.ObjectID) (Submission, error) {
	dbName := os.Getenv("DB_NAME")
	subCol := os.Getenv("SUBMISSION_COLLECTION")
	collection := client.Database(dbName).Collection(subCol)

	filter := bson.M{"_id": sId}

	var s Submission
	err := collection.FindOne(context.TODO(), filter).Decode(&s)
	return s, err
}

func SearchSubmissionsWithUserName(client *mongo.Client, userName string) ([]Submission, error) {
	dbName := os.Getenv("DB_NAME")
	subCol := os.Getenv("SUBMISSION_COLLECTION")
	collection := client.Database(dbName).Collection(subCol)

	filter := bson.M{"userName": userName}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return []Submission{}, err
	}
	var submissions []Submission
	if err = cursor.All(context.TODO(), &submissions); err != nil {
		return []Submission{}, err
	}
	return submissions, nil
}
