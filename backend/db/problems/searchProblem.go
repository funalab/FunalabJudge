package problems

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SearchProblemWithId(client *mongo.Client, problemId int32) (Problem, error) {
	dbName := os.Getenv("DB_NAME")
	prbCol := os.Getenv("PROBLEMS_COLLECTION")
	collection := client.Database(dbName).Collection(prbCol)

	filter := bson.M{"problemId": problemId}

	var p Problem
	err := collection.FindOne(context.TODO(), filter).Decode(&p)
	return p, err
}
