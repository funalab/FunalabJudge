package submission

import (
	"context"
	"go-test/db"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func SearchSubmission(client *mongo.Client, searchField Submission) (Submission, error) {
	dbName := os.Getenv("DB_NAME")
	subCol := os.Getenv("SUBMISSION_COLLECTION")
	collection := client.Database(dbName).Collection(subCol)

	filter := db.MakeFilter(searchField)

	var s Submission
	err := collection.FindOne(context.TODO(), filter).Decode(&s)
	return s, err
}
