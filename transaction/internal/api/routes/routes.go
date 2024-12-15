package routes

import (
	"log"

	"github.com/phuckhoa33/web-crawler/cmd/server"
	health_check_route "github.com/phuckhoa33/web-crawler/internal/api/routes/health-check"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes(server *server.Server) {

	// Version 1 of router
	v1 := server.Gin.Group(server.Config.App.AppApiPrefix + "/v1")

	// Configure for health check route
	health_check_route.ConfigureHealthCheckRoutes(server, v1)

	// Configure path of swagger
	server.Gin.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Talk out swagger host
	log.Printf("Swagger listening and serving at http://%s:%s/docs/index.html", server.Config.App.AppHost, server.Config.App.AppPort)
}
