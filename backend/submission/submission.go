package submission

import (
	"go-test/assignment"
	"time"
)

type Submission struct {
	userId        int64
	problemId     int64
	submittedDate time.Time
	results       []assignment.Testcase
}
