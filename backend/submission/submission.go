package submission

import (
	"go-test/api"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Submission struct {
	Id            int64        `bson:"id"`
	UserId        int64        `bson:"userId"`
	ProblemId     int64        `bson:"problemId"`
	SubmittedDate time.Time    `bson:"submittedDate"`
	Results       []api.Result `bson:"results"`
	Status        string       `bson:"status"`
}
