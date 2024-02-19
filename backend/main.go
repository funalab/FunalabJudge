package main

import (
	"fmt"
	"go-test/api"
	"go-test/assignment"
	"go-test/auth"
	"go-test/compile"
	"go-test/db"
	"go-test/env"
	"go-test/submission"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	env.LoadEnv()
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge:       24 * time.Hour,
		AllowHeaders: []string{"content-type"}, // 他はなくても現状動く
		// 以下の項目は、全てを許可しない設定にしても認証機能に影響はなかった, セキュリティの観点で設定が必要な可能性はある
		// AllowMethods: []string{},
	}))

	err, client := db.Mongo_connectable()
	if err != nil {
		log.Printf("Connection err: %v\n", err.Error())
	}

	router.Use(func(c *gin.Context) {
		c.Set("mongoClient", client)
		c.Next()
	})

	authMiddleware, err := auth.NewJwtMiddleware()
	if err != nil {
		log.Fatal(err)
		return
	}
	router.POST("/login", authMiddleware.LoginHandler)
	router.POST("/logout", authMiddleware.LogoutHandler)
	router.GET("/refresh_token", authMiddleware.RefreshHandler)
	authed := router.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		// ユーザーごとにアクセス権が異なるエンドポイントには、userNameかsubmissionIdを含める
		authed.POST("/changePass/:userName", auth.ChangeUserPass)
		authed.GET("/getAssignmentStatus/:userName", api.GetAssignments)
		authed.GET("/assignmentInfo/:problemId", assignment.AssignmentInfoHandler)
		authed.GET("/submissions/:userName", submission.SubmissionQueueHandler)
		authed.GET("/submission/:submissionId", submission.SubmissionHandler)
		authed.POST("/compile", compile.CompileHandler)
		authed.GET("/maxSubmissionId", submission.MaxSubmissionIdHandler)
		authed.POST("/addSubmission", submission.AddSubmissionHandler)
	}

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// サーバーをポート3000で起動
	router.Run(":3000")
	fmt.Println("Server is running.")
}

// // ハッシュ化した初期パスワードを生成するコード
// // 他のmain関数とimportを全てコメントアウトして実行する
// package main

// import (
// 	"fmt"
// 	"go-test/auth"
// )

// func main() {
// 	hash, _ := auth.HashPassword("password")
// 	fmt.Println(hash)
// }
