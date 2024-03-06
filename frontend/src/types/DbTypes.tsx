export interface Problem {
	Id: number,
	Name: string,
	Statement: string,
	Constraints: string,
	ExecutionTime: number,
	MemoryLimit: number,
	OpenDate: string,
	CloseDate: string,
	BorderScore: number,
	TestcaseWithPaths: TestcaseWithPath[],
}

export interface TestcaseWithPath {
	TestcaseId: number,
	ArgsFilePath: string,
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
	OpenDate: string,
	CloseDate: string,
	BorderScore: number,
	Testcases: Testcase[],
}

export interface Testcase {
	TestcaseId: number,
	ArgsFileContent: string,
	InputFileContent: string,
	OutputFileContent: string,
}