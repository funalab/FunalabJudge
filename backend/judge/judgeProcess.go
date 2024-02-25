package judge

import (
	"fmt"
	"go-test/db/problems"
	"go-test/myTypes"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func JudgeProcess(c *gin.Context, s myTypes.Submission) {
	p := problems.GetProblemFromId(c, s.ProblemId)

	ceFlag := false
	if !isHaveMakeFile(int(s.Id)) {
		err := writeMakeFile(int(s.Id))
		if err != nil {
			log.Println("Failed to write make file :", err.Error())
			updateSubmissionStatus(c, int(s.Id), "CE")
			ceFlag = true
		}
	}

	_, err := execCommand(int(s.Id), "make")
	if err != nil {
		log.Println("Failed to compile :", err.Error())
		updateSubmissionStatus(c, int(s.Id), "CE")
		ceFlag = true
	}

	execFile, err := searchExecutableFile(int(s.Id))
	if err != nil {
		log.Println("Failed to search executable file :", err.Error())
		updateSubmissionStatus(c, int(s.Id), "CE")
		ceFlag = true
	}

	if ceFlag {
		for _, t := range p.TestcaseWithPaths {
			updateSubmissionResult(c, int(s.Id), int(t.TestcaseId), "CE")
		}
		return
	}

	// exec with all test cases
	tLen := len(p.TestcaseWithPaths)
	acNum := 0
	reFlag := false
	tleFlag := false
	for i, t := range p.TestcaseWithPaths {
		updateSubmissionStatus(c, int(s.Id), fmt.Sprintf("%d/%d", i, tLen))

		// exec test case
		input, err := os.ReadFile(filepath.Join("..", t.InputFilePath))
		if err != nil {
			log.Println("Failed to read input of test case :", err.Error())
			updateSubmissionResult(c, int(s.Id), int(t.TestcaseId), "RE")
			reFlag = true
			continue
		}
		output, err := execCommandWithInput(int(s.Id), fmt.Sprintf("./%s", execFile), string(input))
		if err != nil {
			if err.Error() == "signal: killed" {
				log.Println("Failed to run the testcase. TLE is caused :", err.Error())
				updateSubmissionResult(c, int(s.Id), int(t.TestcaseId), "TLE")
				tleFlag = true
			} else {
				log.Println("Failed to run the testcase. RE is caused :", err.Error())
				updateSubmissionResult(c, int(s.Id), int(t.TestcaseId), "RE")
				reFlag = true
			}
			continue
		}

		// judge result
		answer, err := os.ReadFile(filepath.Join("..", t.OutputFilePath))
		if err != nil {
			log.Println("Failed to read output of test case :", err.Error())
			updateSubmissionResult(c, int(s.Id), int(t.TestcaseId), "RE")
			reFlag = true
			continue
		}
		if compareWithAnswer(output, string(answer)) {
			acNum++
			updateSubmissionResult(c, int(s.Id), int(t.TestcaseId), "AC")
		} else {
			updateSubmissionResult(c, int(s.Id), int(t.TestcaseId), "WA")
		}
	}
	// judge pass/fail and update status
	if reFlag {
		updateSubmissionStatus(c, int(s.Id), "RE")
	} else if tleFlag {
		updateSubmissionStatus(c, int(s.Id), "TLE")
	} else if acNum < int(p.BorderScore) {
		updateSubmissionStatus(c, int(s.Id), "WA")
	} else {
		updateSubmissionStatus(c, int(s.Id), "AC")
	}

	_, err = execCommand(int(s.Id), "make clean")
	if err != nil {
		log.Println("Failed to exec make clean :", err.Error())
		updateSubmissionStatus(c, int(s.Id), "RE")
		return
	}
}
