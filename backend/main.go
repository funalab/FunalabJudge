package main

import (
	"fmt"
	"go-test/db"
	"go-test/env"
	"go-test/routes/assignment"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

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
	env.LoadEnv()
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // リクエストを許可するオリジンを指定
	router.Use(cors.New(config))

	router.GET("/", tutorialHandler)
	router.GET("/assignmentInfo/:id", assignment.AssignmentInfoHandler)

	router.Run(":3000")
	fmt.Println("Server is running.")
}
