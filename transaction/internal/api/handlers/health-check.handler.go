package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/phuckhoa33/web-crawler/cmd/server"
	wrapper_responses "github.com/phuckhoa33/web-crawler/internal/domain/responses"
)

type HealthCheckHandler struct {
	server *server.Server
}

func NewHealthCheckHandler(server *server.Server) *HealthCheckHandler {
	return &HealthCheckHandler{
		server: server,
	}
}

// HealthCheck godoc
// @Summary Health check
// @Description Check health of server
// @ID health-check
// @Tags health-check
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Failure 500 {string} string "server error"
// @Router /health-check/database [get]
func (h *HealthCheckHandler) DatabaseHealthCheck(context *gin.Context) {
	db := h.server.DB

	if err := db.DB().Ping(); err != nil {
		wrapper_responses.ErrorResponse(context, 500, "Database is not connected")
		return
	}

	wrapper_responses.Response(context, 200, "ok")
}