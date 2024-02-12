package main

import (
	"fmt"
	"go-test/auth"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Ginルーターを作成
	router := gin.Default()

	// ここからCorsの設定
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
		// 以下の項目は、全てを許可しない設定にしても認証機能に影響はなかった, セキュリティの観点で設定が必要な可能性はある
		// AllowMethods: []string{},  あってもなくても認証機能に影響はなかった
		// AllowHeaders: []string{},  あってもなくても認証機能に影響はなかった
	}))

	jwtMiddleware, err := auth.NewJwtMiddleware()
	if err != nil {
		log.Fatal(err)
		return
	}

	router.POST("/login", jwtMiddleware.LoginHandler)
	router.GET("/refresh_token", jwtMiddleware.RefreshHandler)
	api := router.Group("").Use(jwtMiddleware.MiddlewareFunc())
	{
		api.GET("/test", func(c *gin.Context) {
			// userID := auth.UserIdInJwt(c)
			c.JSON(http.StatusOK, gin.H{
				"userID": "test",
			})
		})
	}

	// サーバーをポート3000で起動
	router.Run(":3000")
	fmt.Println("Server is running.")
}
