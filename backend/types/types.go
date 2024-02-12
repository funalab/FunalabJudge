package types

import "time"

type Data struct {
	Message string `json:"message"`
}

type SubmissionWithStatus struct {
	Submission Submission `bson:",inline"`
	Status     string     `bson:"status"`
}

type ProblemWithStatus struct {
	ProblemId    int       `bson:"problemId"`
	ProblemPath  string    `bson:"problemPath"`
	InputFormat  string    `bson:"inputFormat"`
	OutputFormat string    `bson:"outputFormat"`
	OpenDate     time.Time `bson:"openDate"`
	CloseDate    time.Time `bson:"closeDate"`
	BorderScore  int       `bson:"borderScore"`
	Status       bool      `bson:"status"`
}

type ProblemResp struct {
	Pid       int
	Name      string
	ExTime    int32
	MemLim    int32
	Statement string
	PrbConst  string
	InputFmt  string
	OutputFmt string
	Testcases []Testcase
}

type ProblemRespWithDateInfo struct {
	ProblemResp ProblemResp
	OpenDate    time.Time
	CloseDate   time.Time
	Status      bool
}

type ProblemContainPath struct {
	Pid               int                `bson:"problemId"`
	ProblemPath       string             `bson:"problemPath"`
	InputFmt          string             `bson:"inputFormat"`
	OutputFmt         string             `bson:"outputFormat"`
	TestcaseWithPaths []TestcaseWithPath `bson:"testCases"`
}

type ProblemJSON struct {
	Name          string `json:"name"`
	Statement     string `json:"statement"`
	Constraints   string `json:"constraints"`
	ExecutionTime int32  `json:"executionTime"`
	MemoryLimit   int32  `json:"memoryLimit"`
}

type SampleJSON struct {
	Content string `json:"content"`
}
