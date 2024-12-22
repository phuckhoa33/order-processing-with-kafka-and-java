package repository_interfaces

import models "github.com/phuckhoa33/web-crawler/internal/domain/models/postgresql"

type TransactionRepository interface {
	Create(data *models.Transaction) error
	FindByID(id uint) (*models.Transaction, error)
	FindOne(fields map[string]interface{}) (*models.Transaction, error)
	FindAll() ([]models.Transaction, error)
	Update(data *models.Transaction) error
	Delete(id uint) error
	DeleteAll() error
}