// api/top_page.go

package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"go-test/types"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAssignments(c *gin.Context) {
	dbName := os.Getenv("DB_NAME")
	prbCol := os.Getenv("PROBLEMS_COLLECTION")
	submitCol := os.Getenv("SUBMISSION_COLLECTION")
	client, exists := c.Get("mongoClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database client not available"})
		return
	}
	dbClient := client.(*mongo.Client)

	problemsCollection := dbClient.Database(dbName).Collection(prbCol)
	cur, err := problemsCollection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var problems []types.ProblemWithStatus
	if err = cur.All(context.Background(), &problems); err != nil {
		log.Fatal(err)
	}
	//　続いてstatusを決定する。submissionのテーブルみに行って、userでまず引っ掛ける。その後problemIdごとに全てのテストケースでACになっているsubmittionが存在するかチェック
	submissionCollection := dbClient.Database(dbName).Collection(submitCol)

	for i, problem := range problems {
		fmt.Printf("problemPath: %v\n", problem.ProblemPath)
		filter := bson.D{{Key: "problemId", Value: problem.ProblemId}}
		submissionCursor, err := submissionCollection.Find(context.TODO(), filter)
		if err != nil {
			log.Fatal(err)
		}
		var submissions []types.Submission
		if err = submissionCursor.All(context.TODO(), &submissions); err != nil {
			log.Fatal(err)
		}

		// 対応するsubmissionから全てがACの提出が存在するか確認
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
