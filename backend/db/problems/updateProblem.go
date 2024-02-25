package problems

import (
	"context"
	"go-test/db"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateProblem(client *mongo.Client, searchField Problem, updateField Problem) error {
	dbName := os.Getenv("DB_NAME")
	prbCol := os.Getenv("PROBLEMS_COLLECTION")
	collection := client.Database(dbName).Collection(prbCol)

	sFilter := db.MakeFilter(searchField)
	uFilter := bson.M{"$set": db.MakeFilter(updateField)}

	_, err := collection.UpdateOne(context.TODO(), sFilter, uFilter)
	return err
}
