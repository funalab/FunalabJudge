package api

import (
	"time"
)

type Result struct {
	TestId int    `bson:"testId"`
	Status string `bson:"status"`
}

type Submission struct {
	Id            int       `bson:"id"`
	UserId        int       `bson:"userId"`
	ProblemId     int       `bson:"problemId"`
	SubmittedDate time.Time `bson:"submittedDate"`
	Results       []Result  `bson:"results"`
}
