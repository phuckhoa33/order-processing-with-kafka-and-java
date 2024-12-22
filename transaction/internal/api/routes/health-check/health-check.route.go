package health_check_route

import (
	"github.com/gin-gonic/gin"
	"github.com/phuckhoa33/web-crawler/cmd/server"
	"github.com/phuckhoa33/web-crawler/internal/api/handlers"
)

func ConfigureHealthCheckRoutes(server *server.Server, route *gin.RouterGroup) {
	// import hanlders
	healthCheckHandler := handlers.NewHealthCheckHandler(server)
	// Conigure route
	healthCheckRoute := route.Group("/health-check")
	{
		healthCheckRoute.GET("/database", healthCheckHandler.DatabaseHealthCheck)
	}
}
