package problems

import "time"

type Problem struct {
	Id                int32              `bson:"problemId"`
	Name              string             `bson:"name"`
	Statement         string             `bson:"statement"`
	Constraints       string             `bson:"constraints"`
	ExecutionTime     int32              `bson:"executionTime"`
	MemoryLimit       int32              `bson:"memoryLimit"`
	InputFmt          string             `bson:"inputFormat"`
	OutputFmt         string             `bson:"outputFormat"`
	OpenDate          time.Time          `bson:"openDate"`
	CloseDate         time.Time          `bson:"closeDate"`
	BorderScore       int32              `bson:"borderScore"`
	TestcaseWithPaths []TestcaseWithPath `bson:"testcases"`
}

type TestcaseWithPath struct {
	TestcaseId        int32    `bson:"testCaseId"`
	ArgsFilePath      string   `bson:"argsFilePath"`
	StdinFilePath     string   `bson:"stdinFilePath"`
	InputFilePathList []string `bson:"inputFilePathList"`
	OutputFileName    string   `bson:"outputFileName"`
	AnswerFilePath    string   `bson:"answerFilePath"`
}

type ProblemWithTestcase struct {
	Id            int32      `bson:"problemId"`
	Name          string     `bson:"name"`
	Statement     string     `bson:"statement"`
	Constraints   string     `bson:"constraints"`
	ExecutionTime int32      `bson:"executionTime"`
	MemoryLimit   int32      `bson:"memoryLimit"`
	InputFmt      string     `bson:"inputFormat"`
	OutputFmt     string     `bson:"outputFormat"`
	OpenDate      time.Time  `bson:"openDate"`
	CloseDate     time.Time  `bson:"closeDate"`
	BorderScore   int32      `bson:"borderScore"`
	Testcases     []Testcase `bson:"testcases"`
}

type Testcase struct {
	TestcaseId        int32
	ArgsFileContent   string
	StdinFileContent  string
	InputFileList     []InputFileContent
	AnswerFileContent string
}

type InputFileContent struct {
	FileName string
	Content  string
}

type ProblemWithStatus struct {
	ProblemId   int
	ProblemName string
	Status      string
}
