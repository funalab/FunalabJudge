package submission

import (
	"context"
	"go-test/db"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
)

func GetSubmissionsFromUserId(userId string) (*[]Submission, error) {
	err, client := db.Mongo_connectable()
	if err != nil {
		log.Printf("connection error :%v\n", err.Error())
		return &[]Submission{}, err
	}
	dbName := os.Getenv("DB_NAME")
	collection := client.Database(dbName).Collection("USERS_COLLECTION")

	filter := bson.M{"userId": userId}
	var submissions []Submission
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Printf("Failed to find single result from DB: %v\n", err.Error())
		return &[]Submission{}, err
	}

	if err = cursor.All(context.TODO(), &submissions); err != nil {
		log.Printf("Failed to find single result from DB: %v\n", err.Error())
		return &[]Submission{}, err
	}
	return &submissions, nil
}
