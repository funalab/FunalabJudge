package judge

import (
	"errors"
	"fmt"
	"go-test/db/problems"
	"go-test/db/submission"
	"go-test/util"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

const DEFAULT_OUTPUT_FILENAME = "output"

func JudgeProcess(c *gin.Context, s submission.Submission) {
	client_, exists := c.Get("mongoClient")
	if !exists {
		util.ResponseDBNotFoundError(c)
	}
	client := client_.(*mongo.Client)

	p, err := problems.SearchOneProblemWithId(client, s.ProblemId)
	if err != nil {
		c.Error(fmt.Errorf("[%s] failed to find single result: %s", s.Id.Hex(), err.Error()))
		submission.UpdateSubmissionStatus(client, s.Id, "RE")
		return
	}

	ceFlag := false
	r, err := isHaveMakeFile(s.Id.Hex())
	if err != nil {
		c.Error(errors.Join(fmt.Errorf("[%s] failed to find exec dir", s.Id.Hex()), err))
	}
	if !r {
		err := writeMakeFile(s.Id.Hex())
		if err != nil {
			c.Error(errors.Join(fmt.Errorf("[%s] failed to write make file", s.Id.Hex()), err))
			submission.UpdateSubmissionStatus(client, s.Id, "CE")
			ceFlag = true
		}
	}

	_, err = execCommand(s.Id, "make", 2)
	if err != nil {
		c.Error(errors.Join(fmt.Errorf("[%s] failed to compile", s.Id.Hex()), err))
		submission.UpdateSubmissionStatus(client, s.Id, "CE")
		ceFlag = true
	}

	execFile, err := searchExecutableFile(s.Id)
	if err != nil {
		c.Error(errors.Join(fmt.Errorf("[%s] failed to find executable file", s.Id.Hex()), err))
		submission.UpdateSubmissionStatus(client, s.Id, "CE")
		ceFlag = true
	}

	if ceFlag {
		for _, t := range p.TestcaseWithPaths {
			submission.UpdateSubmissionResult(client, s.Id, int(t.TestcaseId), "CE")
		}
		_, err = execCommand(s.Id, "make clean", 2)
		if err != nil {
			c.Error(errors.Join(fmt.Errorf("[%s] failed to exec make clean", s.Id.Hex()), err))
			submission.UpdateSubmissionStatus(client, s.Id, "RE")
			return
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
		if t.ArgsFilePath != "" {
			a, err := os.ReadFile(filepath.Join(staticDir, t.ArgsFilePath))
			if err != nil {
				c.Error(errors.Join(fmt.Errorf("[%s] failed to read args of test case", s.Id.Hex()), err))
				submission.UpdateSubmissionResult(client, s.Id, int(t.TestcaseId), "RE")
				reFlag = true
				continue
			} else {
				command = command + " " + string(a)
			}
		}
		if t.InputFilePath != "" {
			// stdinのpipeを使うとバカ長いinputを入れるときに正常に動かなくなるので、リダイレクトする
			// TODO 相対パス使ってる応急処置でnot elegant、絶対パスにしたい
			command = command + " < ../" + filepath.Join(staticDir, t.InputFilePath)
		}

		output, err := execCommand(s.Id, command, int(p.ExecutionTime))
		if err != nil {
			if err.Error() == "signal: killed" {
				c.Error(errors.Join(fmt.Errorf("[%s] failed to run the testcase, TLE is caused", s.Id.Hex()), err))
				submission.UpdateSubmissionResult(client, s.Id, int(t.TestcaseId), "TLE")
				tleFlag = true
			} else {
				c.Error(errors.Join(fmt.Errorf("[%s] failed to run the testcase, RE is caused", s.Id.Hex()), err))
				submission.UpdateSubmissionResult(client, s.Id, int(t.TestcaseId), "RE")
				reFlag = true
			}
			continue
		}

		// file.cのために用意したフィールド
		if t.OutputFilePath != "" {
			output, err = os.ReadFile(filepath.Join(os.Getenv("EXEC_DIR"), s.Id.Hex(), t.OutputFilePath))
			if err != nil {
				c.Error(errors.Join(fmt.Errorf("[%s] failed to read output of test case", s.Id.Hex()), err))
				submission.UpdateSubmissionResult(client, s.Id, int(t.TestcaseId), "RE")
				reFlag = true
				continue
			}
			// outputファイルを削除
			if err := os.Remove(filepath.Join(os.Getenv("EXEC_DIR"), s.Id.Hex(), t.OutputFilePath)); err != nil {
				c.Error(errors.Join(fmt.Errorf("[%s] failed to remove output file", s.Id.Hex()), err))
				submission.UpdateSubmissionResult(client, s.Id, int(t.TestcaseId), "RE")
				reFlag = true
				continue
			}
		}

		// judge result
		answer, err := os.ReadFile(filepath.Join(staticDir, t.AnswerFilePath))
		if err != nil {
			c.Error(errors.Join(fmt.Errorf("[%s] failed to read answer of test case", s.Id.Hex()), err))
			submission.UpdateSubmissionResult(client, s.Id, int(t.TestcaseId), "RE")
			reFlag = true
			continue
		}
		if compareWithAnswer(string(output), string(answer)) {
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

	_, err = execCommand(s.Id, "make clean", 2)
	if err != nil {
		c.Error(errors.Join(fmt.Errorf("[%s] failed to exec make clean", s.Id.Hex()), err))
		submission.UpdateSubmissionStatus(client, s.Id, "RE")
		return
	}
}
