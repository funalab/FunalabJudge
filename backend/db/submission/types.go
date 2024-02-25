package submission

import "time"

type Submission struct {
	Id            int32     `bson:"id"`
	UserId        int32     `bson:"userId"`
	ProblemId     int32     `bson:"problemId"`
	SubmittedDate time.Time `bson:"submittedDate"`
	Results       []Result  `bson:"results"`
	Status        string    `bson:"status"`
}

type Result struct {
	TestId int    `bson:"testCaseId"`
	Status string `bson:"status"`
}

type SubmissionWithStatus struct {
	Submission Submission `bson:",inline"`
	Status     string     `bson:"status"`
}
