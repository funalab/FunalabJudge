package main

import (
	"context"
	"fmt"
	"go-test/api"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	router.GET("/api/assignments", api.GetAssignments)

	// サーバーをポート3000で起動
	router.Run(":3000")
	fmt.Println("Server is running.")
}
