package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func initMongoClient() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// MongoDBへの接続を確認
	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB")
}

// レスポンスとして返すデータ
type Data struct {
	Message string `json:"message"`
}

func main() {
	initMongoClient()
	// Ginルーターを作成
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // リクエストを許可するオリジンを指定
	router.Use(cors.New(config))

	router.Use(func(c *gin.Context) {
		c.Set("mongoClient", mongoClient)
		c.Next()
	})

	// エンドポイントのハンドラー関数を設定

	router.GET("/api/assignments", GetAssignments)

	// サーバーをポート3000で起動
	router.Run(":3000")
	fmt.Println("Server is running.")
}

// 後々以下は別ファイルに移行する

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
	client, exists := c.Get("mongoClient")
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
	var problems []Problems
	if err = cur.All(context.Background(), &problems); err != nil {
		log.Fatal(err)
	}

	//　続いてstatusを決定する。submittionのテーブルみに行って、userでまず引っ掛ける。その後problemIdごとに全てのテストケースでACになっているsubmittionが存在するかチェック
	submissionCollection := dbClient.Database("dev").Collection("submission")

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
