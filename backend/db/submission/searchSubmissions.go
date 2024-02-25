package submission

import (
	"context"
	"go-test/db"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func SearchSubmission(client *mongo.Client, searchField Submission) Submission {
	dbName := os.Getenv("DB_NAME")
	subCol := os.Getenv("SUBMISSION_COLLECTION")
	collection := client.Database(dbName).Collection(subCol)

	filter := db.MakeFilter(searchField)

	var s Submission
	err := collection.FindOne(context.TODO(), filter).Decode(&s)
	if err != nil {
		log.Printf("Failed to find single result from DB: %v\n", err.Error())
		return Submission{}
	}
	return s
}
