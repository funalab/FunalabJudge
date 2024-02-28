package submission

import (
	"context"
	"go-test/db/problems"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertNewSubmission(client *mongo.Client, userName string, p problems.Problem) (Submission, error) {

	dbName := os.Getenv("DB_NAME")
	subCol := os.Getenv("SUBMISSION_COLLECTION")
	collection := client.Database(dbName).Collection(subCol)

	s := initialSubmissionDocument(userName, p)
	_, err := collection.InsertOne(context.TODO(), s)
	return s, err
}

func initialSubmissionDocument(userName string, p problems.Problem) Submission {
	s := Submission{
		Id:            primitive.NewObjectID(),
		UserName:      userName,
		ProblemId:     p.Id,
		SubmittedDate: time.Now(),
		Status:        "WJ",
	}
	nt := len(p.TestcaseWithPaths)
	s.Results = make([]Result, nt)
	for index, ele := range s.Results {
		ele.TestId = index + 1
		ele.Status = "WJ"
		s.Results[index] = ele
	}
	return s
}
