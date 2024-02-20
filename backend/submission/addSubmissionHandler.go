package submission

import (
	"context"
	"go-test/types"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TmpSubmission struct {
	SubmissionId  int32
	UserName      string
	ProblemId     int32
	SubmittedDate time.Time
}

func AddSubmissionHandler(c *gin.Context) {
	client, exists := c.Get("mongoClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "DB client is not available."})
	}

	dbName := os.Getenv("DB_NAME")
	submitCol := os.Getenv("SUBMISSION_COLLECTION")
	col := (client.(*mongo.Client)).Database(dbName).Collection(submitCol)

	/*Bind Submission and push into db*/
	// submissionId: submissionId,
	// userName: userName,
	// problemId: problemId,
	// submittedDate: new Date(),

	var ts TmpSubmission
	if err := c.BindJSON(&ts); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
	}
	/*Map to Submission Struct*/
	s := mapToSubmission(c, &ts)
	col.InsertOne(context.TODO(), *s)
	c.JSON(200, nil)
}

func mapToSubmission(c *gin.Context, ts *TmpSubmission) *types.Submission {
	un := ts.UserName
	client, exists := c.Get("mongoClient")
	if !exists {
		return nil
	}
	col := (client.(*mongo.Client)).Database(os.Getenv("DB_NAME")).Collection(os.Getenv("USERS_COLLECTION"))
	filter := bson.M{"userName": un}
	var u types.User
	err := col.FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		log.Println("Failed to find such user.")
		return nil
	}
	var s types.Submission
	s.Id = ts.SubmissionId
	s.UserId = u.UserId
	s.ProblemId = ts.ProblemId
	s.SubmittedDate = ts.SubmittedDate

	/*Map Results*/
	col = (client.(*mongo.Client)).Database(os.Getenv("DB_NAME")).Collection(os.Getenv("PROBLEMS_COLLECTION"))
	filter = bson.M{"problemId": s.ProblemId}
	var p types.Problem
	err = col.FindOne(context.TODO(), filter).Decode(&p)
	if err != nil {
		log.Println("Failed to find such problem.")
		return nil
	}
	nt := len(p.TestcaseWithPaths)
	s.Results = make([]types.Result, nt)

	for index, ele := range s.Results {
		ele.TestId = index + 1
		ele.Status = "WJ"
		s.Results[index] = ele
	}
	return &s
}
