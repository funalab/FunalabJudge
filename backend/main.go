package main

import (
	"fmt"
	"go-test/auth"
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
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
		MaxAge:       24 * time.Hour,
		AllowHeaders: []string{"content-type"}, // 他はなくても現状動く
		// 以下の項目は、全てを許可しない設定にしても認証機能に影響はなかった, セキュリティの観点で設定が必要な可能性はある
		// AllowMethods: []string{},
	}))

	authMiddleware, err := auth.NewJwtMiddleware()
	if err != nil {
		log.Fatal(err)
		return
	}

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router.POST("/login", authMiddleware.LoginHandler)
	router.POST("/logout", authMiddleware.LogoutHandler)
	router.GET("/refresh_token", authMiddleware.RefreshHandler)
	authed := router.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		authed.GET("/test", func(c *gin.Context) {
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
