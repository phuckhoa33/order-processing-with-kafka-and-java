package services

import (
	"fmt"

	"github.com/phuckhoa33/web-crawler/cmd/server"
	models "github.com/phuckhoa33/web-crawler/internal/domain/models/postgresql"
	"github.com/phuckhoa33/web-crawler/internal/domain/repositories"
	"github.com/phuckhoa33/web-crawler/internal/domain/requests"
)


type TransactionManagementService struct {
	Repo *repositories.TransactionRepository
	Server *server.Server
}

func NewTransactionManagementService(server *server.Server) *TransactionManagementService {
	return &TransactionManagementService{
		Repo:   repositories.NewTransactionRepository(server.DB),
		Server: server,
	}
}

// Create a new data source
func (s *TransactionManagementService) CreatePayTransaction(dto requests.CreateTransactionRequestDto) error {
	// Create data source
	transaction := models.Transaction{
		UserID: dto.UserID,
		OrderID: dto.OrderID,
		Amount: 899,
		Action: string(dto.Action),
	}

	if err := s.Repo.Create(&transaction); err != nil {
		fmt.Println(err)
		return fmt.Errorf("cannot create transaction with order id %s", dto.OrderID)
	}

	return nil
}

