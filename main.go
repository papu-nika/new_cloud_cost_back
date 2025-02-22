package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/papu-nika/new_cloud_cost_back/api"
	_ "github.com/papu-nika/new_cloud_cost_back/docs"

	"github.com/papu-nika/new_cloud_cost_back/server"
)

func main() {
	r := gin.Default()

	s := server.Server{}

	api.RegisterHandlers(r, &s)
	// api.RegisterHandlersWithOptions(r, &server, api.GinServerOptions{})

	r.GET("/events", server.SSEHandler)           // SSE 接続
	r.POST("/post", server.PostHandler)           // クライアントからの POST を受信
	r.GET("/messages", server.GetMessagesHandler) // クライアントからの GET を受信

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:5173"}, // React(Vite)のオリジンを許可
		AllowMethods:  []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:  []string{"Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
		// AllowCredentials: true, // Cookie などの認証情報を送信可能にする
		// AllowAllOrigins: true, // 全てのオリジンを許可
	}))

	// Swagger UI の提供
	addSwaggerUI(r)

	r.Run(":8080")
}
