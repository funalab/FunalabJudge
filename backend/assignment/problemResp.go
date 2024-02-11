package assignment

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
