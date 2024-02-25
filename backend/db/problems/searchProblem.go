package problems

import (
	"context"
	"go-test/db"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func SearchProblem(client *mongo.Client, searchField Problem) Problem {
	dbName := os.Getenv("DB_NAME")
	prbCol := os.Getenv("PROBLEM_COLLECTION")
	collection := client.Database(dbName).Collection(prbCol)

	filter := db.MakeFilter(searchField)

	var p Problem
	err := collection.FindOne(context.TODO(), filter).Decode(&p)
	if err != nil {
		log.Printf("Failed to find single result from DB: %v\n", err.Error())
		return Problem{}
	}
	return p
}
