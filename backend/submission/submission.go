package submission

import "go-test/assignment"

type Submission struct {
	userId        int64
	problemId     int64
	submittedDate string
	results       []assignment.Testcase
}
