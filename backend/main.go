package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func mongo_connectable() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	mongoClient, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI("mongodb://localhost:27017/"),
	)
	if err != nil {
		log.Fatalf("connection error :%v", err)
		return false
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("ping mongodb error :%v", err)
		return false
	}
	cancel()
	if err := mongoClient.Disconnect(ctx); err != nil {
		log.Fatalf("mongodb disconnect error : %v", err)
		return false
	}
	return true
}

// レスポンスとして返すデータ
type Data struct {
	Message string `json:"message"`
}

func main() {
	// Ginルーターを作成
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // リクエストを許可するオリジンを指定
	router.Use(cors.New(config))

	// エンドポイントのハンドラー関数を設定
	router.GET("/", func(c *gin.Context) {
		if mongo_connectable() {
			// レスポンスデータの作成
			data := Data{
				Message: "Hello from Gin and mongo!!",
			}
			// JSON形式でレスポンスを返す
			c.JSON(200, data)
		}
	})
	router.POST("/login", func(c *gin.Context) {
		authUser(c)
	})

	// サーバーをポート3000で起動
	router.Run(":3000")
	fmt.Println("Server is running.")
}
