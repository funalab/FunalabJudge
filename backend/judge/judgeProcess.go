package judge

import (
	"fmt"
	"go-test/assignment"
	"go-test/myTypes"
	"go-test/problems"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
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
		// return myTypes.NewMakeFailErr(fmt.Sprintf("Failed to execute make command: %v\n", err.Error()))
		log.Println("Failed to compile")
		updateSubmissionStatus(c, int(s.Id), "CE")
		return
	}

	execFile, err := searchExecutableFile(int(s.Id))
	if err != nil {
		log.Println(err.Error())
		updateSubmissionStatus(c, int(s.Id), "CE")
		return
	}
	// 全テストケースをジャッジする
	// 以下はassignment.TranslatePathIntoProblemRespを使うために流用, 整えたい
	dbName := os.Getenv("DB_NAME")
	prbCol := os.Getenv("PROBLEMS_COLLECTION")
	client, exists := c.Get("mongoClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "DB client is not available."})
		return
	}
	collection := (client.(*mongo.Client)).Database(dbName).Collection(prbCol)
	p := assignment.TranslatePathIntoProblemResp(collection, int(s.ProblemId))
	tLen := len(p.Testcases)
	acNum := 0
	for i, t := range p.Testcases {
		updateSubmissionStatus(c, int(s.Id), fmt.Sprintf("%d/%d", i, tLen))

		output, err := execCommandWithInput(int(s.Id), fmt.Sprintf("./%s", execFile), t.InputFileContent)
		if err != nil {
			log.Println("Failed to run the testcase. RE is caused.")
			log.Println(err.Error())
			updateSubmissionResult(c, int(s.Id), int(t.TestcaseId), "RE")
			continue
		}

		// 実行結果をジャッジする
		if compareWithAnswer(output, t.OutputFileContent) {
			acNum++
			updateSubmissionResult(c, int(s.Id), int(t.TestcaseId), "AC")
		} else {
			updateSubmissionResult(c, int(s.Id), int(t.TestcaseId), "WA")
		}
	}
	// 合否を判定して更新
	if acNum >= int(problems.GetProblemFromId(c, s.ProblemId).BorderScore) {
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
