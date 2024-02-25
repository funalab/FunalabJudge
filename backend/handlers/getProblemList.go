package handlers

import (
	"context"
	"log"
	"net/http"
	"os"

	"go-test/assignment"
	"go-test/db/submission"
	"go-test/myTypes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetProblemListHandler(c *gin.Context) {
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
	var problems []myTypes.ProblemWithStatus
	if err = cur.All(context.Background(), &problems); err != nil {
		log.Fatal(err)
	}
	//　続いてstatusを決定する。submissionのテーブルみに行って、userでまず引っ掛ける。その後problemIdごとに全てのテストケースでACになっているsubmittionが存在するかチェック
	// TODO userでの絞り込みは現状してない, 全てのユーザーの全ての提出をみている
	submissionCollection := dbClient.Database(dbName).Collection(submitCol)
	resps := make([]myTypes.ProblemRespWithDateInfo, 0)
	for _, problem := range problems {
		filter := bson.M{"problemId": problem.ProblemId}
		submissionCursor, err := submissionCollection.Find(context.TODO(), filter)
		if err != nil {
			log.Fatal(err)
		}
		var submissions []submission.Submission
		if err = submissionCursor.All(context.TODO(), &submissions); err != nil {
			log.Fatal(err)
		}

		// 対応するsubmissionから全てがACの提出が存在するか確認
		// TODO 全てがACではなく、problemのboarderScore以上に修正が必要
		for _, submission := range submissions {
			allAC := true
			for _, object := range submission.Results {
				if object.Status != "AC" {
					allAC = false
					break
				}
			}
			if allAC {
				problem.Status = true
				break
			} else {
				problem.Status = false
			}
		}
		resp := assignment.TranslatePathIntoProblemResp(problemsCollection, problem.ProblemId)
		resp2 := new(myTypes.ProblemRespWithDateInfo)
		resp2.ProblemResp = *resp
		resp2.Status = problem.Status
		resp2.OpenDate = problem.OpenDate
		resp2.CloseDate = problem.CloseDate
		resps = append(resps, *resp2)
	}

	c.JSON(http.StatusOK, resps)
}
