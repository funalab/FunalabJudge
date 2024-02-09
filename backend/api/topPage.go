// api/top_page.go

package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetAssignments は /api/assignments で呼び出される
func GetAssignments(c *gin.Context) {
	client, exists := c.Get("dbClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database client not available"})
		return
	}
	dbClient := client.(*mongo.Client)

	// まずはproblemsのテーブルから全ての問題をとってくる
	problemsCollection := dbClient.Database("dev").Collection("problems")
	cur, err := problemsCollection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	if err = cur.All(context.Background(), &problems); err != nil {
		log.Fatal(err)
	}
	//　続いてstatusを決定する。これがちょい大変。submittionのテーブルみに行って、userでまず引っ掛ける。その後problemIdごとに全てのテストケースでACになっているsubmittionが存在するかチェック
	submissionCollection := dbClient.Database("dev").Collection("submittion")

	for _, problem := range problems {
		fmt.Printf("Problem ID: %v\n", problem["_id"])
		filter := bson.D{{"problem_id", problem["_id"]}}
		submissionCursor, err := submissionCollection.Find(context.TODO(), filter)
		if err != nil {
			log.Fatal(err)
		}
		var submissions []bson.M
		if err = submissionCursor.All(context.TODO(), &submissions); err != nil {
			log.Fatal(err)
		}

		// 対応するsubmissionの情報を出力
		for _, submission := range submissions {
			fmt.Printf("Submission: %+v\n", submission)
		}
	}

}
