export interface Problem {
	Id: number,
	Name: string,
	Statement: string,
	Constraints: string,
	ExecutionTime: number,
	MemoryLimit: number,
	InputFmt: string,
	OutputFmt: string,
	OpenDate: string,
	CloseDate: string,
	BorderScore: number,
	TestcaseWithPaths: TestcaseWithPath[],
}

export interface TestcaseWithPath {
	TestcaseId: number,
	ArgsFilePath: string,
	StdinFilePath: string,
	InputFilePathList: string[],
	OutputFileName: string,
	AnswerFilePath: string,
}

export interface ProblemWithTestcase {
	Id: number,
	Name: string,
	Statement: string,
	Constraints: string,
	ExecutionTime: number,
	MemoryLimit: number,
	InputFmt: string,
	OutputFmt: string,
	OpenDate: string,
	CloseDate: string,
	BorderScore: number,
	Testcases: Testcase[],
}

export interface Testcase {
	TestcaseId: number,
	ArgsFileContent: string,
	StdinFileContent: string
	InputFileList: InputFileContent
	AnswerFileContent: string,
}

export interface InputFileContent {
	FileName: string
	Content: string
}