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
	hangarRepo := persistence.NewMemoryProjectsRepository(providers.UpdateHangar)
	modrinthRepo := persistence.NewMemoryProjectsRepository(providers.UpdateModrinth)
	spigetRepo := persistence.NewMemoryProjectsRepository(providers.UpdateSpiget)

	// Services
	cfwidgetService := services.NewProjectsService(cfwidgetRepo)
	hangarService := services.NewProjectsService(hangarRepo)
	modrinthService := services.NewProjectsService(modrinthRepo)
	spigetService := services.NewProjectsService(spigetRepo)

	// Handlers
	cfwidgetHandler := handlers.NewProjectsHandler(cfwidgetService)
	hangarHandler := handlers.NewProjectsHandler(hangarService)
	modrinthHandler := handlers.NewProjectsHandler(modrinthService)
	spigetHandler := handlers.NewProjectsHandler(spigetService)

	// Router
	router = gin.Default()
	router.Use(cors.Default())

	router.GET("cfwidget", cfwidgetHandler.GetProjects)
	router.GET("hangar", hangarHandler.GetProjects)
	router.GET("modrinth", modrinthHandler.GetProjects)
	router.GET("spiget", spigetHandler.GetProjects)

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
