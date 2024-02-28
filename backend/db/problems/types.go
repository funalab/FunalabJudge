package problems

import "time"

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
