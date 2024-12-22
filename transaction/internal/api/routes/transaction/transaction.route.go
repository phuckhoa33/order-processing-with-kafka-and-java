package transaction_route

import (
	"github.com/gin-gonic/gin"
	"github.com/phuckhoa33/web-crawler/cmd/server"
	"github.com/phuckhoa33/web-crawler/internal/api/handlers"
)

func ConfigureTransactionRoute(server *server.Server, route *gin.RouterGroup) {
	// import handler
	handler := handlers.NewTransactionManagementHandler(server)

	// Configure route
	transactionRoute := route.Group("/transactions")
	{
		transactionRoute.POST("/pay", handler.CreatePayTransaction)
	}
}