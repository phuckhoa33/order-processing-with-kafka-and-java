package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phuckhoa33/web-crawler/cmd/server"
	"github.com/phuckhoa33/web-crawler/internal/domain/requests"
	wrapper_responses "github.com/phuckhoa33/web-crawler/internal/domain/responses"
	"github.com/phuckhoa33/web-crawler/internal/domain/services"
)

type TransactionManagementHandler struct {
	Server *server.Server
	Service *services.TransactionManagementService
}

func NewTransactionManagementHandler(server *server.Server) *TransactionManagementHandler {
	return &TransactionManagementHandler{
		Server: server,
		Service: services.NewTransactionManagementService(server),
	}
}

// CreatePayTransaction godoc
// @Summary Create a pay transaction
// @Description Create a pay transaction
// @ID create-pay-transaction
// @Tags transaction-management
// @Accept json
// @Produce json
// @Param transaction body requests.CreatePayTransaction true "Data source object"
// @Success 200
// @Failure 400 {string} string "BAD_REQUEST"
// @Router /transactions/pay [post]
func (h *TransactionManagementHandler) CreatePayTransaction(context *gin.Context) {
	// get request
	request := new(requests.CreateTransactionRequestDto)
	if err := context.Bind(&request); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	// Check field is empty
	if err := request.Validate(); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	// Create data source
	if err := h.Service.CreatePayTransaction(*request); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	wrapper_responses.Response(context, http.StatusOK, nil)
}
