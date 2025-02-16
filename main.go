package main

import (
	"github.com/gin-gonic/gin"
	"github.com/papu-nika/new_cloud_cost_back/api"
	_ "github.com/papu-nika/new_cloud_cost_back/docs"

	"github.com/papu-nika/new_cloud_cost_back/server"
)

func main() {
	r := gin.Default()

	server := server.Server{}

	api.RegisterHandlers(r, &server)
	// api.RegisterHandlersWithOptions(r, &server, api.GinServerOptions{})

	// Swagger UI の提供
	addSwaggerUI(r)

	r.Run(":8080")
}
