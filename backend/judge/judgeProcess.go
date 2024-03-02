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
	p, err := problems.SearchOneProblemWithId(client, s.ProblemId)
	if err != nil {
		log.Fatalf("Failed to find single result from DB: %v\n", err)
	}

	ceFlag := false
	if !isHaveMakeFile(s.Id.Hex()) {
		err := writeMakeFile(s.Id.Hex())
		if err != nil {
			log.Println("Failed to write make file :", err.Error())
			submission.UpdateSubmissionStatus(client, s.Id, "CE")
			ceFlag = true
		}
	}

	_, err = execCommand(s.Id, "make")
	if err != nil {
		log.Println("Failed to compile :", err.Error())
		submission.UpdateSubmissionStatus(client, s.Id, "CE")
		ceFlag = true
	}

	execFile, err := searchExecutableFile(s.Id)
	if err != nil {
		log.Println("Failed to search executable file :", err.Error())
		submission.UpdateSubmissionStatus(client, s.Id, "CE")
		ceFlag = true
	}

	if ceFlag {
		for _, t := range p.TestcaseWithPaths {
			submission.UpdateSubmissionResult(client, s.Id, int(t.TestcaseId), "CE")
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
		submission.UpdateSubmissionStatus(client, s.Id, fmt.Sprintf("%d/%d", i, tLen))
		command := fmt.Sprintf("./%s", execFile)

		// exec test case
		if t.ArgsFilePath != nil {
			a, err := os.ReadFile(filepath.Join(staticDir, *t.ArgsFilePath))
			if err != nil {
				log.Println("Failed to read args of test case :", err.Error())
				submission.UpdateSubmissionResult(client, s.Id, int(t.TestcaseId), "RE")
				reFlag = true
				continue
			} else {
				command = command + " " + string(a)
			}
		}
		if t.InputFilePath != nil {
			// stdinのpipeを使うとバカ長いinputを入れるときに正常に動かなくなるので、リダイレクトする
			// TODO 相対パス使ってる応急処置でnot elegant、絶対パスにしたい
			command = command + " < ../" + filepath.Join(staticDir, *t.InputFilePath)
		}

		output, err := execCommand(s.Id, command)
		if err != nil {
			if err.Error() == "signal: killed" {
				log.Println("Failed to run the testcase. TLE is caused :", err.Error())
				submission.UpdateSubmissionResult(client, s.Id, int(t.TestcaseId), "TLE")
				tleFlag = true
			} else {
				log.Println("Failed to run the testcase. RE is caused :", err.Error())
				submission.UpdateSubmissionResult(client, s.Id, int(t.TestcaseId), "RE")
				reFlag = true
			}
			continue
		}

		// judge result
		answer, err := os.ReadFile(filepath.Join(staticDir, t.OutputFilePath))
		if err != nil {
			log.Println("Failed to read output of test case :", err.Error())
			submission.UpdateSubmissionResult(client, s.Id, int(t.TestcaseId), "RE")
			reFlag = true
			continue
		}
		if compareWithAnswer(output, string(answer)) {
			acNum++
			submission.UpdateSubmissionResult(client, s.Id, int(t.TestcaseId), "AC")
		} else {
			submission.UpdateSubmissionResult(client, s.Id, int(t.TestcaseId), "WA")
		}
	}
	// judge pass/fail and update status
	if reFlag {
		submission.UpdateSubmissionStatus(client, s.Id, "RE")
	} else if tleFlag {
		submission.UpdateSubmissionStatus(client, s.Id, "TLE")
	} else if acNum < int(p.BorderScore) {
		submission.UpdateSubmissionStatus(client, s.Id, "WA")
	} else {
		submission.UpdateSubmissionStatus(client, s.Id, "AC")
	}

	_, err = execCommand(s.Id, "make clean")
	if err != nil {
		log.Println("Failed to exec make clean :", err.Error())
		submission.UpdateSubmissionStatus(client, s.Id, "RE")
		return
	}
}
