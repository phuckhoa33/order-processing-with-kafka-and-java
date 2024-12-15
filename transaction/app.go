package application

import (
	"log"

	"github.com/phuckhoa33/web-crawler/cmd/server"
	"github.com/phuckhoa33/web-crawler/internal/api/routes"
	"github.com/phuckhoa33/web-crawler/internal/config"
)

func Start(config *config.Config) {
	// Initiate server
	app := server.NewServer(config)

	// Configure routes
	routes.ConfigureRoutes(app)

	// Run port
	err := app.Run(config.App.AppPort)

	// Check err when run port
	if err != nil {
		log.Fatal("Port already used, please change to another port")
	}
}
