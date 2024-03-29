package problems

import (
	"context"
	"errors"
	"go-test/db"
	"os"
	"path/filepath"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SearchOneProblemWithId(client *mongo.Client, problemId int32) (Problem, error) {
	dbName := os.Getenv("DB_NAME")
	prbCol := os.Getenv("PROBLEMS_COLLECTION")
	collection := client.Database(dbName).Collection(prbCol)

	filter := bson.M{"problemId": problemId}

	var p Problem
	err := collection.FindOne(context.TODO(), filter).Decode(&p)
	return p, err
}

func SearchProblems(client *mongo.Client, searchField Problem) ([]Problem, error) {
	dbName := os.Getenv("DB_NAME")
	prbCol := os.Getenv("PROBLEMS_COLLECTION")
	collection := client.Database(dbName).Collection(prbCol)

	sFilter := db.MakeFilter(searchField)

	cursor, err := collection.Find(context.TODO(), sFilter)
	if err != nil {
		return []Problem{}, err
	}
	defer cursor.Close(context.TODO())

	var p []Problem
	if err := cursor.All(context.TODO(), &p); err != nil {
		return []Problem{}, err
	}

	return p, nil
}

func ReadTestcaseContent(p Problem) (ProblemWithTestcase, error) {
	t, err := parseTestcaseWithPathToTestcase(p.TestcaseWithPaths)
	return ProblemWithTestcase{
		Id:            p.Id,
		Name:          p.Name,
		Statement:     p.Statement,
		Constraints:   p.Constraints,
		ExecutionTime: p.ExecutionTime,
		MemoryLimit:   p.MemoryLimit,
		InputFmt:      p.InputFmt,
		OutputFmt:     p.OutputFmt,
		OpenDate:      p.OpenDate,
		CloseDate:     p.CloseDate,
		BorderScore:   p.BorderScore,
		Testcases:     t,
	}, err
}

func parseTestcaseWithPathToTestcase(tws []TestcaseWithPath) ([]Testcase, error) {
	staticDir := os.Getenv("STATIC_DIR")
	ts := make([]Testcase, 0)
	for _, tw := range tws {
		ts_ := Testcase{TestcaseId: tw.TestcaseId}
		if tw.ArgsFilePath != "" {
			sArgs, err := os.ReadFile(filepath.Join(staticDir, tw.ArgsFilePath))
			if err != nil {
				return []Testcase{}, errors.Join(errors.New("failed to read args file"), err)
			}
			ts_.ArgsFileContent = string(sArgs)
		}
		if tw.StdinFilePath != "" {
			sIn, err := os.ReadFile(filepath.Join(staticDir, tw.StdinFilePath))
			if err != nil {
				return []Testcase{}, errors.Join(errors.New("failed to read stdin file"), err)
			}
			ts_.StdinFileContent = string(sIn)
		}
		inputs := make([]InputFileContent, 0)
		for _, f := range tw.InputFilePathList {
			inpt, err := os.ReadFile(filepath.Join(staticDir, f))
			if err != nil {
				return []Testcase{}, errors.Join(errors.New("failed to read input file"), err)
			}
			inputs = append(inputs, InputFileContent{
				FileName: filepath.Base(f),
				Content:  string(inpt),
			})
		}
		ts_.InputFileList = inputs
		if tw.AnswerFilePath != "" {
			sOut, err := os.ReadFile(filepath.Join(staticDir, tw.AnswerFilePath))
			if err != nil {
				return []Testcase{}, errors.Join(errors.New("failed to read output file"), err)
			}
			ts_.AnswerFileContent = string(sOut)
		}
		ts = append(ts, ts_)
	}
	return ts, nil
}
