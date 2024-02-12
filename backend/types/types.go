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
	Pid       int        `bson:"problemId"`
	Name      string     `bson:"name"`
	ExTime    int32      `bson:"executionTime"`
	MemLim    int32      `bson:"memoryLimit"`
	Statement string     `bson:"statement"`
	PrbConst  string     `bson:"problemConstraints"`
	InputFmt  string     `bson:"inputFormat"`
	OutputFmt string     `bson:"outputFormat"`
	Testcases []Testcase `bson:"testCases"`
}

type ProblemContainPath struct {
	Pid         int        `bson:"problemId"`
	ProblemPath string     `bson:"problemPath"`
	InputFmt    string     `bson:"inputFormat"`
	OutputFmt   string     `bson:"outputFormat"`
	Testcases   []Testcase `bson:"testCases"`
}

type ProblemJSON struct {
	Name          string `json:"name"`
	Statement     string `json:"statement"`
	Constraints   string `json:"constraints"`
	ExecutionTime int32  `json:"executionTime"`
	MemoryLimit   int32  `json:"memoryLimit"`
}
