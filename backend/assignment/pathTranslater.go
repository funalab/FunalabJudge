package assignment

import (
	"context"
	"encoding/json"
	"go-test/db/problems"
	"go-test/myTypes"
	"go-test/util"
	"io"
	"log"
	"os"
	"path/filepath"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProblemContainPath struct {
	Pid               int                         `bson:"problemId"`
	ProblemPath       string                      `bson:"problemPath"`
	InputFmt          string                      `bson:"inputFormat"`
	OutputFmt         string                      `bson:"outputFormat"`
	TestcaseWithPaths []problems.TestcaseWithPath `bson:"testCases"`
}

func TranslatePathIntoProblemResp(coll *mongo.Collection, pid int) *myTypes.ProblemResp {
	var p ProblemContainPath

	err := coll.FindOne(context.TODO(), bson.M{"problemId": pid}).Decode(&p)
	if err != nil {
		log.Fatalf("Failed to parse problemId as a number: %v\n", err.Error())
		return nil
	}
	pf, err := util.OpenFileFromDB(p.ProblemPath)
	if err != nil {
		log.Fatalf("Failed to open problem file: %v\n", err.Error())
		return nil
	}
	pj, err := parseProblemJSON(pf)
	if err != nil {
		log.Fatalf("Failed to parase problem json: %v\n", err.Error())
		return nil
	}
	return mapToProblemResp(&p, pj)
}

func parseProblemJSON(pf *os.File) (*myTypes.ProblemJSON, error) {
	var pj myTypes.ProblemJSON
	b, err := io.ReadAll(pf)
	if err != nil {
		log.Fatalf("Failed to read problem file as byte: %v\n", err.Error())
		return &myTypes.ProblemJSON{}, err
	}
	err = json.Unmarshal(b, &pj)
	if err != nil {
		log.Fatalf("Failed to unmarshal json file: %v\n", err.Error())
		return &myTypes.ProblemJSON{}, err
	}
	return &pj, nil
}

/*TODO: Check whether parsing would be done correctly.*/
func parseTestcaseWithPathIntoTestcase(tws []problems.TestcaseWithPath) *[]myTypes.Testcase {
	ts := make([]myTypes.Testcase, 0)
	for _, tw := range tws {
		sIn, err := os.ReadFile(filepath.Join("..", tw.InputFilePath))
		if err != nil {
			log.Fatalf("Failed to parse into sample json: %v\n", err.Error())
		}
		sOut, err := os.ReadFile(filepath.Join("..", tw.OutputFilePath))
		if err != nil {
			log.Fatalf("Failed to parse into sample json: %v\n", err.Error())
		}
		t := mapToTestcase(tw, string(sIn), string(sOut))
		//After parsing both input and output file, append to slice.
		ts = append(ts, *t)
	}
	return &ts
}

func mapToProblemResp(p *ProblemContainPath, pj *myTypes.ProblemJSON) *myTypes.ProblemResp {
	pr := new(myTypes.ProblemResp)
	pr.Pid = p.Pid
	pr.Name = pj.Name
	pr.ExTime = pj.ExecutionTime
	pr.MemLim = pj.MemoryLimit
	pr.Statement = pj.Statement
	pr.PrbConst = pj.Constraints
	pr.InputFmt = p.InputFmt
	pr.OutputFmt = p.OutputFmt
	pr.Testcases = *parseTestcaseWithPathIntoTestcase(p.TestcaseWithPaths)
	return pr
}
func mapToTestcase(tw problems.TestcaseWithPath, sIn string, sOut string) *myTypes.Testcase {
	t := new(myTypes.Testcase)
	t.TestcaseId = tw.TestcaseId
	t.InputFileContent = sIn
	t.OutputFileContent = sOut
	return t
}
