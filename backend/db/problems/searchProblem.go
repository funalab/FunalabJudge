package problems

import (
	"context"
	"go-test/db"
	"log"
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

func ReadTestcaseContent(p Problem) ProblemWithTestcase {
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
		Testcases:     parseTestcaseWithPathIntoTestcase(p.TestcaseWithPaths),
	}
}

func parseTestcaseWithPathIntoTestcase(tws []TestcaseWithPath) []Testcase {
	staticDir := os.Getenv("STATIC_DIR")
	ts := make([]Testcase, 0)
	for _, tw := range tws {
		sIn, err := os.ReadFile(filepath.Join(staticDir, tw.InputFilePath))
		if err != nil {
			log.Fatalf("Failed to parse into sample json: %v\n", err.Error())
		}
		sOut, err := os.ReadFile(filepath.Join(staticDir, tw.OutputFilePath))
		if err != nil {
			log.Fatalf("Failed to parse into sample json: %v\n", err.Error())
		}
		ts = append(ts, Testcase{
			TestcaseId:        tw.TestcaseId,
			InputFileContent:  string(sIn),
			OutputFileContent: string(sOut),
		})
	}
	return ts
}
