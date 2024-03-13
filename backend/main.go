package main

import (
	"flag"
	"fmt"
	"go-test/auth"
	"go-test/db"
	"go-test/handlers"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	var (
		releaseFlag = flag.Bool("release", false, "flag for release mode")
	)
	flag.Parse()
	if *releaseFlag {
		gin.SetMode(gin.ReleaseMode)
	}

	if err := godotenv.Load("../frontend/.env"); err != nil {
		log.Fatal("Failed to load .env file.")
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://" + os.Getenv("PUBLIC_SERVER_IP") + ":" + os.Getenv("FRONTEND_PORT"),
		},
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge:       24 * time.Hour,
		AllowHeaders: []string{"content-type"}, // 他はなくても現状動く
		// 以下の項目は、全てを許可しない設定にしても認証機能に影響はなかった, セキュリティの観点で設定が必要な可能性はある
		// AllowMethods: []string{},
	}))

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	f, _ := os.Create("gin.log")
	log.SetOutput(f)
	gin.DefaultWriter = io.MultiWriter(f)
	router.Use(gin.LoggerWithFormatter(loggerFormatter))

	client, err := db.Mongo_connectable()
	if err != nil {
		log.Fatalf("Connection err: %v\n", err.Error())
	}
	router.Use(func(c *gin.Context) {
		c.Set("mongoClient", client)
		c.Next()
	})

	authMiddleware, err := auth.NewJwtMiddleware()
	if err != nil {
		log.Fatal(err)
	}
	router.POST("/login", authMiddleware.LoginHandler)
	router.POST("/logout", authMiddleware.LogoutHandler)
	router.GET("/refresh_token", authMiddleware.RefreshHandler)
	authed := router.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		// ユーザーごとにアクセス権が異なるエンドポイントには、userNameかsubmissionIdを含める
		authed.POST("/changePassword/:userName", handlers.ChangePasswordHandler)
		authed.GET("/getProblemList/:userName", handlers.GetProblemListHandler)
		authed.GET("/getProblem/:problemId", handlers.GetProblemHandler)
		authed.GET("/getSubmissionList/:userName", handlers.GetSubmissionListHandler)
		authed.GET("/getSubmission/:submissionId", handlers.GetSubmissionHandler)
		authed.POST("/addSubmission/:userName", handlers.AddSubmissionHandler)
		authed.GET("/getSubmittedFiles/:submissionId", handlers.GetSubmittedFilesHandler)

	}

	router.GET("/getB3Status", handlers.GetB3StatusHandler)
	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// サーバー起動
	router.Run(os.Getenv("PRIVATE_SERVER_IP") + ":" + os.Getenv("BACKEND_PORT"))
	fmt.Println("Server is running.")
}

func loggerFormatter(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s | %d | %s | %s | %s \"%s\" %s\n",
		param.TimeStamp.Format(time.RFC3339),
		param.StatusCode,
		param.Latency,
		param.ClientIP,
		param.Method,
		param.Path,
		param.Request.UserAgent(),
	)
}
