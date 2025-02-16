package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func addSwaggerUI(r *gin.Engine) {
	r.GET("/swagger/doc.json", func(c *gin.Context) {
		file, err := os.Open("openapi.yaml")
		if err != nil {
			log.Printf("Failed to open openapi.yaml: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load OpenAPI definition"})
			return
		}
		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			log.Printf("Failed to read openapi.yaml: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read OpenAPI definition"})
			return
		}
		c.Data(http.StatusOK, "application/json", content)
	})
	r.Use(static.Serve("/swagger", static.LocalFile("./docs", true)))
}
