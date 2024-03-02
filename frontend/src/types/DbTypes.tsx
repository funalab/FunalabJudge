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
	InputFilePath: string,
	OutputFilePath: string,
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
	InputFileContent: string,
	OutputFileContent: string,
}