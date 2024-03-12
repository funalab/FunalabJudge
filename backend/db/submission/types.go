package submission

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Submission struct {
	Id            primitive.ObjectID `bson:"_id"`
	UserName      string             `bson:"userName"`
	ProblemId     int32              `bson:"problemId"`
	SubmittedDate time.Time          `bson:"submittedDate"`
	Results       []Result           `bson:"results"`
	Status        string             `bson:"status"`
}

type Result struct {
	TestId int    `bson:"testCaseId"`
	Status string `bson:"status"`
}

type SubmittedFile struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type SubmissionWithProblemName struct {
	Submission
	ProblemName string
}
