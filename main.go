package main

import (
	"log"
	"my-stats/internal/adapters/handlers"
	"my-stats/internal/adapters/persistence"
	"my-stats/internal/domain/providers"
	"my-stats/internal/ports/services"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	_ = os.Setenv("HOST", "0.0.0.0")
	_ = os.Setenv("PORT", "8088")

	// Repositories
	cfwidgetRepo := persistence.NewMemoryProjectsRepository(providers.UpdateCFWidget)

	// Services
	cfwidgetService := services.NewProjectsService(cfwidgetRepo)

	// Handlers
	cfwidgetHandler := handlers.NewProjectsHandler(cfwidgetService)

	// Router
	router = gin.Default()
	router.Use(cors.Default())

	router.GET("cfwidget", cfwidgetHandler.GetProjects)

	log.Printf("Server running on http://localhost:%s\n", os.Getenv("PORT"))

	address := os.Getenv("HOST") + ":" + os.Getenv("PORT")

	var err error
	if strings.EqualFold(os.Getenv("SSL_ENABLED"), "true") {
		err = router.Run(address, os.Getenv("SSL_CERT"), os.Getenv("SSL_KEY"))
	} else {
		err = router.Run(address)
	}

	if err != nil {
		log.Fatalf("Error starting the server: %s", err)
	}
}
