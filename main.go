package main

import (
	"my-stats/internal/adapters/persistence"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	_ = os.Setenv("HOST", "0.0.0.0")
	_ = os.Setenv("PORT", "8088")

	// Repositories

	// Services

	// Handlers

	// Router
	persistence.Update()
	router = gin.Default()
	router.Use(cors.Default())

	//router.GET("/", handlers.HandleMain)
}
