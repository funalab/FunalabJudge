package assignment

import (
	"context"
	"encoding/json"
	"go-test/types"
	"go-test/util"
	"io"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var backendProjectRootPath string

func TranslatePathIntoProblemResp(coll *mongo.Collection, pid int) *types.ProblemResp {
	var p types.ProblemContainPath

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

func parseProblemJSON(pf *os.File) (*types.ProblemJSON, error) {
	var pj types.ProblemJSON
	b, err := io.ReadAll(pf)
	if err != nil {
		log.Fatalf("Failed to read problem file as byte: %v\n", err.Error())
		return &types.ProblemJSON{}, err
	}
	err = json.Unmarshal(b, &pj)
	if err != nil {
		log.Fatalf("Failed to unmarshal json file: %v\n", err.Error())
		return &types.ProblemJSON{}, err
	}
	return &pj, nil
}

/*TODO: Check whether parsing would be done correctly.*/
func parseTestcaseWithPathIntoTestcase(tws *[]types.TestcaseWithPath) *[]types.Testcase {
	ts := make([]types.Testcase, 0)
	for _, tw := range *tws {
		inPath := tw.InputFilePath
		inFile, err := util.OpenFileFromDB(inPath)
		if err != nil {
			log.Fatalf("Failed to open input file: %v\n", err.Error())
		}
		sIn, err := parseSampleJSON(inFile)
		if err != nil {
			log.Fatalf("Failed to parse into sample json: %v\n", err.Error())
		}
		outPath := tw.OutputFilePath
		outFile, err := util.OpenFileFromDB(outPath)
		if err != nil {
			log.Fatalf("Failed to open output file: %v\n", err.Error())
		}
		sOut, err := parseSampleJSON(outFile)
		if err != nil {
			log.Fatalf("Failed to parse into sample json: %v\n", err.Error())
		}
		t := mapToTestcase(tw, sIn, sOut)
		//After parsing both input and output file, append to slice.
		ts = append(ts, *t)
	}
	return &ts
}

func parseSampleJSON(f *os.File) (*types.SampleJSON, error) {
	var s types.SampleJSON
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatalf("Failed to read samples as byte: %v\n", err.Error())
		return &types.SampleJSON{}, err
	}
	err = json.Unmarshal(b, &s)
	if err != nil {
		log.Fatalf("Failed to unmarshal json file: %v\n", err.Error())
		return &types.SampleJSON{}, err
	}
	return &s, nil
}

func mapToProblemResp(p *types.ProblemContainPath, pj *types.ProblemJSON) *types.ProblemResp {
	pr := new(types.ProblemResp)
	pr.Pid = p.Pid
	pr.Name = pj.Name
	pr.ExTime = pj.ExecutionTime
	pr.MemLim = pj.MemoryLimit
	pr.Statement = pj.Statement
	pr.PrbConst = pj.Constraints
	pr.InputFmt = p.InputFmt
	pr.OutputFmt = p.OutputFmt
	pr.Testcases = *parseTestcaseWithPathIntoTestcase(&p.TestcaseWithPaths)
	return pr
}
func mapToTestcase(tw types.TestcaseWithPath, sIn *types.SampleJSON, sOut *types.SampleJSON) *types.Testcase {
	t := new(types.Testcase)
	t.TestcaseId = tw.TestcaseId
	t.InputFileContent = sIn.Content
	t.OutputFileContent = sOut.Content
	return t
}
