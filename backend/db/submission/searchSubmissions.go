package submission

import (
	"context"
	"go-test/db"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SearchOneSubmissionWithId(client *mongo.Client, sId primitive.ObjectID) (Submission, error) {
	dbName := os.Getenv("DB_NAME")
	subCol := os.Getenv("SUBMISSION_COLLECTION")
	collection := client.Database(dbName).Collection(subCol)

	filter := bson.M{"_id": sId}

	var s Submission
	err := collection.FindOne(context.TODO(), filter).Decode(&s)
	return s, err
}

func SearchSubmissions(client *mongo.Client, searchField Submission) ([]Submission, error) {
	dbName := os.Getenv("DB_NAME")
	subCol := os.Getenv("SUBMISSION_COLLECTION")
	collection := client.Database(dbName).Collection(subCol)

	sFilter := db.MakeFilter(searchField)

	cursor, err := collection.Find(context.TODO(), sFilter)
	if err != nil {
		return []Submission{}, err
	}
	defer cursor.Close(context.Background())
	var submissions []Submission
	if err = cursor.All(context.TODO(), &submissions); err != nil {
		return []Submission{}, err
	}
	return submissions, nil
}

// isPetitCoderのproblemIdかつACなsubmissionを、userNameでdistinctし、submittedDateで並び替えて返す
func SearchPetitCoderSubmissions(client *mongo.Client, problemId int32) ([]Submission, error) {
	dbName := os.Getenv("DB_NAME")
	subCol := os.Getenv("SUBMISSION_COLLECTION")
	collection := client.Database(dbName).Collection(subCol)

	pipeline := []bson.M{
		{"$match": bson.M{
			"problemId": problemId,
			"status":    "AC",
		}},
		{"$sort": bson.M{"userName": 1, "submittedDate": 1}},
		{"$group": bson.M{
			"_id":        "$userName", // userNameでグループ化
			"firstEntry": bson.M{"$first": "$$ROOT"},
		}},
		{"$replaceRoot": bson.M{"newRoot": "$firstEntry"}}, // グループごとの最初のエントリを新しいルートに置き換える
	}

	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return []Submission{}, err
	}
	defer cursor.Close(context.Background())

	var submissions []Submission
	if err = cursor.All(context.TODO(), &submissions); err != nil {
		return []Submission{}, err
	}
	return submissions, nil
}
