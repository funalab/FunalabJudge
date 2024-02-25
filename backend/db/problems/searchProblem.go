package problems

import (
	"context"
	"go-test/db"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func SearchProblem(client *mongo.Client, searchField Problem) (Problem, error) {
	dbName := os.Getenv("DB_NAME")
	prbCol := os.Getenv("PROBLEM_COLLECTION")
	collection := client.Database(dbName).Collection(prbCol)

	filter := db.MakeFilter(searchField)

	var p Problem
	err := collection.FindOne(context.TODO(), filter).Decode(&p)
	return p, err
}
