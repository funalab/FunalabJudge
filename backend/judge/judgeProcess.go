package judge

import (
	"fmt"
	"go-test/db/problems"
	"go-test/db/submission"
	"log"
	"os"
	"path/filepath"

	"go.mongodb.org/mongo-driver/mongo"
)

func JudgeProcess(client *mongo.Client, s submission.Submission) {
	ceFlag := false
	if !isHaveMakeFile(s.Id.Hex()) {
		err := writeMakeFile(s.Id.Hex())
		if err != nil {
			log.Println("Failed to write make file :", err.Error())
			updateSubmissionStatus(client, s.Id, "CE")
			ceFlag = true
		}
	}

	_, err := execCommand(s.Id, "make")
	if err != nil {
		log.Println("Failed to compile :", err.Error())
		updateSubmissionStatus(client, s.Id, "CE")
		ceFlag = true
	}

	execFile, err := searchExecutableFile(s.Id)
	if err != nil {
		log.Println("Failed to search executable file :", err.Error())
		updateSubmissionStatus(client, s.Id, "CE")
		ceFlag = true
	}

	p, err := problems.SearchProblemWithId(client, s.ProblemId)
	if err != nil {
		log.Fatalf("Failed to find single result from DB: %v\n", err)
	}
	if ceFlag {
		for _, t := range p.TestcaseWithPaths {
			updateSubmissionResult(client, s.Id, int(t.TestcaseId), "CE")
		}
		return
	}

	// exec with all test cases
	staticDir := os.Getenv("STATIC_DIR")
	tLen := len(p.TestcaseWithPaths)
	acNum := 0
	reFlag := false
	tleFlag := false
	for i, t := range p.TestcaseWithPaths {
		updateSubmissionStatus(client, s.Id, fmt.Sprintf("%d/%d", i, tLen))

		// exec test case
		input, err := os.ReadFile(filepath.Join(staticDir, t.InputFilePath))
		if err != nil {
			log.Println("Failed to read input of test case :", err.Error())
			updateSubmissionResult(client, s.Id, int(t.TestcaseId), "RE")
			reFlag = true
			continue
		}
		output, err := execCommandWithInput(s.Id, fmt.Sprintf("./%s", execFile), string(input))
		if err != nil {
			if err.Error() == "signal: killed" {
				log.Println("Failed to run the testcase. TLE is caused :", err.Error())
				updateSubmissionResult(client, s.Id, int(t.TestcaseId), "TLE")
				tleFlag = true
			} else {
				log.Println("Failed to run the testcase. RE is caused :", err.Error())
				updateSubmissionResult(client, s.Id, int(t.TestcaseId), "RE")
				reFlag = true
			}
			continue
		}

		// judge result
		answer, err := os.ReadFile(filepath.Join(staticDir, t.OutputFilePath))
		if err != nil {
			log.Println("Failed to read output of test case :", err.Error())
			updateSubmissionResult(client, s.Id, int(t.TestcaseId), "RE")
			reFlag = true
			continue
		}
		if compareWithAnswer(output, string(answer)) {
			acNum++
			updateSubmissionResult(client, s.Id, int(t.TestcaseId), "AC")
		} else {
			updateSubmissionResult(client, s.Id, int(t.TestcaseId), "WA")
		}
	}
	// judge pass/fail and update status
	if reFlag {
		updateSubmissionStatus(client, s.Id, "RE")
	} else if tleFlag {
		updateSubmissionStatus(client, s.Id, "TLE")
	} else if acNum < int(p.BorderScore) {
		updateSubmissionStatus(client, s.Id, "WA")
	} else {
		updateSubmissionStatus(client, s.Id, "AC")
	}

	_, err = execCommand(s.Id, "make clean")
	if err != nil {
		log.Println("Failed to exec make clean :", err.Error())
		updateSubmissionStatus(client, s.Id, "RE")
		return
	}
}
