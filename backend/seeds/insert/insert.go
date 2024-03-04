package main

import (
	"context"
	"encoding/json"
	"flag"
	"go-test/db"
	"go-test/db/problems"
	"go-test/db/submission"
	"go-test/db/users"
	"log"
	"os"
	"path/filepath"
	"slices"

	"github.com/joho/godotenv"
)

// insert document
func main() {
	if err := godotenv.Load("../frontend/.env"); err != nil {
		log.Fatal("Failed to load .env file.")
	}

	var (
		targCol  = flag.String("c", "", "target collection")
		targFile = flag.String("f", "", "target json file")
	)
	flag.Parse()

	client, err := db.Mongo_connectable()
	if err != nil {
		log.Fatalf("Connection err: %v\n", err.Error())
	}

	// validate args
	staticDir := os.Getenv("STATIC_DIR")
	seedDataDir := os.Getenv("SEED_DATA_DIR")
	usrCol := os.Getenv("USERS_COLLECTION")
	prbCol := os.Getenv("PROBLEMS_COLLECTION")
	subCol := os.Getenv("SUBMISSION_COLLECTION")
	if !slices.Contains([]string{usrCol, prbCol, subCol}, *targCol) {
		log.Fatalf("Collection not found: %s\n", *targCol)
	}
	jsonData, err := os.ReadFile(filepath.Join(staticDir, seedDataDir, *targFile))
	if err != nil {
		log.Fatalf("file not found in %s\n", filepath.Join(staticDir, seedDataDir))
	}
	// parse json and insert data
	dbName := os.Getenv("DB_NAME")
	var insertInterface []interface{}
	switch *targCol {
	case usrCol:
		var insertData []users.User
		err = json.Unmarshal(jsonData, &insertData)
		if err != nil {
			log.Fatalf("failed to decode json :%s\n", err.Error())
		}
		for _, v := range insertData {
			insertInterface = append(insertInterface, v)
		}
	case prbCol:
		var insertData []problems.Problem
		err = json.Unmarshal(jsonData, &insertData)
		if err != nil {
			log.Fatalf("failed to decode json :%s\n", err.Error())
		}
		for _, v := range insertData {
			insertInterface = append(insertInterface, v)
		}
	case subCol:
		var insertData []submission.Submission
		err = json.Unmarshal(jsonData, &insertData)
		if err != nil {
			log.Fatalf("failed to decode json :%s\n", err.Error())
		}
		for _, v := range insertData {
			insertInterface = append(insertInterface, v)
		}
	}
	insertResult, err := client.Database(dbName).Collection(*targCol).InsertMany(context.TODO(), insertInterface)
	if err != nil {
		log.Fatalf("failed to insert :%s\n", err.Error())
	}
	log.Printf("Inserted %v documents in %s collection!\n", len(insertResult.InsertedIDs), *targCol)
}
