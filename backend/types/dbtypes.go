package types

import "time"

type Submission struct {
	Id            int32     `bson:"id"`
	UserId        int32     `bson:"userId"`
	ProblemId     int32     `bson:"problemId"`
	SubmittedDate time.Time `bson:"submittedDate"`
	Results       []Result  `bson:"results"`
}

type Result struct {
	TestId int    `bson:"testCaseId"`
	Status string `bson:"status"`
}

type Problem struct {
	Id          int32
	Path        string
	InputFmt    string
	OutputFmt   string
	OpenDate    time.Time
	CloseDate   time.Time
	BorderScore int32
	Testcases   []Testcase
}

type Testcase struct {
	TestcaseId     int32  `bson:"testCaseId"`
	InputFilePath  string `bson:"inputFilePath"`
	OutputFilePath string `bson:"outputFilePath"`
}

type User struct {
	UserId      int64
	Email       string
	Password    string
	CreatedDate time.Time
	Role        string
}
