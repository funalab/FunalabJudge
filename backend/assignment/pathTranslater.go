package assignment

import (
	"context"
	"encoding/json"
	"go-test/types"
	"go-test/util"
	"io"
	"log"
	"os"
	"path/filepath"

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

	backendProjectRootPath = util.GetBackendProjectRoot()
	if err != nil {
		log.Fatalf("Failed to get current working directory path: %v\n", err.Error())
		return nil
	}

	parsedPath := parseIntoProjectPath(p.ProblemPath)
	pf, err := os.Open(parsedPath)
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
	byteValue, err := io.ReadAll(pf)
	if err != nil {
		log.Fatalf("Failed to read problem file as byte: %v\n", err.Error())
		return &types.ProblemJSON{}, err
	}
	err = json.Unmarshal(byteValue, &pj)
	if err != nil {
		log.Fatalf("Failed to unmarshal json file: %v\n", err.Error())
		return &types.ProblemJSON{}, err
	}
	return &pj, nil
}

func parseIntoProjectPath(path string) string {
	return filepath.Join(filepath.Dir(backendProjectRootPath), path)
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
	pr.Testcases = p.Testcases
	return pr
}
