package main

import (
	"context"
	"fmt"
	"go-test/api"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-test/db"
	"go-test/env"
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

type Data struct {
	Message string `json:"message"`
}

func tutorialHandler(c *gin.Context) {
	if db.Mongo_connectable() {
		data := Data{
			Message: "Hello fron Gin and mongo!!",
		}
		c.JSON(200, data)
	}
}

func main() {
	err, mongoClient := db.Mongo_connectable()
	// Ginルーターを作成
	env.LoadEnv()
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // リクエストを許可するオリジンを指定
	router.Use(cors.New(config))

	router.Use(func(c *gin.Context) {
		c.Set("mongoClient", mongoClient)
		c.Next()
	})

	router.GET("/api/assignments", api.GetAssignments)
	router.GET("/", tutorialHandler)
	router.GET("/assignmentInfo/:id", api.AssignmentInfoHandler)

	router.Run(":3000")
	fmt.Println("Server is running.")
}
