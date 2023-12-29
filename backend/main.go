package main

import (
	"github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
)

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
		// レスポンスデータの作成
		data := Data{
			Message: "Hello from Gin!!",
		}

		// JSON形式でレスポンスを返す
		c.JSON(200, data)
	})

	// サーバーをポート3000で起動
	router.Run(":3000")
}

