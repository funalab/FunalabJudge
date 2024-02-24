package judge

import (
	"fmt"
	"go-test/myTypes"
	"go-test/problems"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func JudgeProcess(c *gin.Context, s myTypes.Submission) {
	if !isHaveMakeFile(int(s.Id)) {
		err := writeMakeFile(int(s.Id))
		if err != nil {
			log.Println("Failed to write make file")
			updateSubmissionStatus(c, int(s.Id), "CE")
			return
		}
	}

	_, err := execCommand(int(s.Id), "make")
	if err != nil {
		log.Println("Failed to compile", err.Error())
		updateSubmissionStatus(c, int(s.Id), "CE")
		return
	}

	execFile, err := searchExecutableFile(int(s.Id))
	if err != nil {
		log.Println("Failed to search executable file", err.Error())
		updateSubmissionStatus(c, int(s.Id), "CE")
		return
	}
	// 全テストケースをジャッジする
	p := problems.GetProblemFromId(c, s.ProblemId)
	tLen := len(p.TestcaseWithPaths)
	acNum := 0
	for i, t := range p.TestcaseWithPaths {
		updateSubmissionStatus(c, int(s.Id), fmt.Sprintf("%d/%d", i, tLen))

		input, err := os.ReadFile(filepath.Join("..", t.InputFilePath))
		if err != nil {
			log.Println("Failed to read input of test case", err.Error())
			updateSubmissionResult(c, int(s.Id), int(t.TestcaseId), "RE")
			continue
		}
		output, err := execCommandWithInput(int(s.Id), fmt.Sprintf("./%s", execFile), string(input))
		if err != nil {
			log.Println("Failed to run the testcase. RE is caused.", err.Error())
			updateSubmissionResult(c, int(s.Id), int(t.TestcaseId), "RE")
			continue
		}

		// 実行結果をジャッジする
		answer, err := os.ReadFile(filepath.Join("..", t.OutputFilePath))
		if err != nil {
			log.Println("Failed to read output of test case", err.Error())
			updateSubmissionResult(c, int(s.Id), int(t.TestcaseId), "RE")
			continue
		}
		if compareWithAnswer(output, string(answer)) {
			acNum++
			updateSubmissionResult(c, int(s.Id), int(t.TestcaseId), "AC")
		} else {
			updateSubmissionResult(c, int(s.Id), int(t.TestcaseId), "WA")
		}
	}
	// 合否を判定して更新
	if acNum >= int(p.BorderScore) {
		updateSubmissionStatus(c, int(s.Id), "AC")
	} else {
		updateSubmissionStatus(c, int(s.Id), "WA")
	}

	// make cleanする
	_, err = execCommand(int(s.Id), "make clean")
	if err != nil {
		log.Println("Failed to exec make clean", err.Error())
		updateSubmissionStatus(c, int(s.Id), "RE")
		return
	}
}
