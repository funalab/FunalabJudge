package main

import (
	"context"
	"flag"
	"go-test/db"
	"log"
	"os"
	"slices"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

// delete db
func main() {
	if err := godotenv.Load("../frontend/.env"); err != nil {
		log.Fatal("Failed to load .env file.")
	}
	var (
		targCol = flag.String("c", "", "target collection")
	)
	flag.Parse()

	client, err := db.Mongo_connectable()
	if err != nil {
		log.Fatalf("Connection err: %v\n", err.Error())
	}
	usrCol := os.Getenv("USERS_COLLECTION")
	prbCol := os.Getenv("PROBLEMS_COLLECTION")
	subCol := os.Getenv("SUBMISSION_COLLECTION")
	if !slices.Contains([]string{usrCol, prbCol, subCol}, *targCol) {
		log.Fatalf("Collection not found: %s\n", *targCol)
	}
	dbName := os.Getenv("DB_NAME")
	deleteResult, err := client.Database(dbName).Collection(*targCol).DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		log.Fatalf("failed to delete :%s\n", err.Error())
	}
	log.Printf("Deleted %v documents in %s collection!\n", deleteResult.DeletedCount, *targCol)
}
