package main

import (
	"fmt"
	"go-test/api"
	"go-test/assignment"
	"go-test/db"
	"go-test/env"
	"go-test/submission"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Data struct {
	Message string `json:"message"`
}

func tutorialHandler(c *gin.Context) {
	err, _ := db.Mongo_connectable()
	if err == nil {
		data := Data{
			Message: "Hello fron Gin and mongo!!",
		}
		c.JSON(200, data)
	}
}

func main() {
	env.LoadEnv()
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // リクエストを許可するオリジンを指定
	router.Use(cors.New(config))
	err, client := db.Mongo_connectable()
	if err != nil {
		log.Printf("Connection err: %v\n", err.Error())
	}

	router.Use(func(c *gin.Context) {
		c.Set("mongoClient", client)
		c.Next()
	})

	router.GET("/", api.GetAssignments)
	router.GET("/assignmentInfo/:id", assignment.AssignmentInfoHandler)
	router.GET("/submissions/:userId", submission.SubmissionQueueHandler)

	router.Run(":3000")
	fmt.Println("Server is running.")
}
