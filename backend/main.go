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
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Access-Control-Request-Method",
			"Access-Control-Request-Headers",
			"Access-Control-Allow-Origin",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
			"X-Requested-With",
			"Origin,Accept",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// AllowAllOrigins:  true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	jwtMiddleware, err := auth.NewJwtMiddleware()
	if err != nil {
		log.Fatal(err)
		return
	}

	router.POST("/login", jwtMiddleware.LoginHandler)
	router.GET("/refresh_token", jwtMiddleware.RefreshHandler)
	api := router.Group("/api").Use(jwtMiddleware.MiddlewareFunc())
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
