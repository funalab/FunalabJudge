package myTypes

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

type Problem struct {
	Id                int32              `bson:"problemId"`
	Path              string             `bson:"problemPath"`
	InputFmt          string             `bson:"inputFormat"`
	OutputFmt         string             `bson:"outputFormat"`
	OpenDate          time.Time          `bson:"openDate"`
	CloseDate         time.Time          `bson:"closeDate"`
	BorderScore       int32              `bson:"borderScore"`
	TestcaseWithPaths []TestcaseWithPath `bson:"testcases"`
}

type TestcaseWithPath struct {
	TestcaseId     int32  `bson:"testCaseId"`
	InputFilePath  string `bson:"inputFilePath"`
	OutputFilePath string `bson:"outputFilePath"`
}

type Testcase struct {
	TestcaseId        int32
	InputFileContent  string
	OutputFileContent string
}

type User struct {
	UserId      int32     `bson:"userId"`
	UserName    string    `bson:"userName"`
	Password    string    `bson:"password"`
	CreatedDate time.Time `bson:"createdDate"`
	Role        string    `bson:"role"`
}
