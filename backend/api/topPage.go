// api/top_page.go

package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Problems struct {
	ProblemId          int       `bson:"problemId"`
	Name               string    `bson:"name"`
	ExecutionTime      int       `bson:"executionTime"`
	MemoryLimit        int       `bson:"memoryLimit"`
	Statement          string    `bson:"statement"`
	ProblemConstraints string    `bson:"problemConstraints"`
	InputFormat        string    `bson:"inputFormat"`
	OutputFormat       string    `bson:"outputFormat"`
	OpenDate           time.Time `bson:"openDate"`
	CloseDate          time.Time `bson:"closeDate"`
	BorderScore        int       `bson:"borderScore"`
	Status             bool      `bson:"status"`
}

type Result struct {
	TestId int    `bson:"testId"`
	Status string `bson:"status"`
}

type Submission struct {
	UserId        int       `bson:"userId"`
	ProblemId     int       `bson:"problemId"`
	SubmittedDate time.Time `bson:"submittedDate"`
	Results       []Result  `bson:"results"`
}

func GetAssignments(c *gin.Context) {
	dbName := os.Getenv("DB_NAME")
	prbCol := os.Getenv("PROBLEMS_COLLECTION")
	client, exists := c.Get("mongoClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database client not available"})
		return
	}
	dbClient := client.(*mongo.Client)

	// まずはproblemsのテーブルから全ての問題をとってくる
	problemsCollection := dbClient.Database(dbName).Collection(prbCol)
	cur, err := problemsCollection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var problems []Problems
	if err = cur.All(context.Background(), &problems); err != nil {
		log.Fatal(err)
	}
	//　続いてstatusを決定する。submittionのテーブルみに行って、userでまず引っ掛ける。その後problemIdごとに全てのテストケースでACになっているsubmittionが存在するかチェック
	submissionCollection := dbClient.Database(dbName).Collection("submission")

	for i, problem := range problems {
		filter := bson.D{{Key: "problemId", Value: problem.ProblemId}}
		submissionCursor, err := submissionCollection.Find(context.TODO(), filter)
		if err != nil {
			log.Fatal(err)
		}
		var submissions []Submission
		if err = submissionCursor.All(context.TODO(), &submissions); err != nil {
			log.Fatal(err)
		}

		// 対応するsubmissionから全てがACの提出が存在するか確認
		fmt.Printf("%v\n", submissions)
		for _, submission := range submissions {
			allAC := true
			for _, object := range submission.Results {
				if object.Status != "AC" {
					allAC = false
					break
				}
			}
			if allAC {
				problems[i].Status = true
				break
			} else {
				problems[i].Status = false
			}
		}
	}
	c.JSON(http.StatusOK, problems)
}
