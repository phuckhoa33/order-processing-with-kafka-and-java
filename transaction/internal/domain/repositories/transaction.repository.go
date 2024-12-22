package repositories

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	models "github.com/phuckhoa33/web-crawler/internal/domain/models/postgresql"
)


type TransactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) Create(transaction *models.Transaction) error {
	return r.DB.Create(transaction).Error
}

func (r *TransactionRepository) FindByID(id uuid.UUID) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := r.DB.First(&transaction, "id = ?", id.String()).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepository) FindOne(fields map[string]interface{}) (*models.Transaction, error) {
	var transaction models.Transaction
	query := r.DB

	for field, value := range fields {
		query = query.Where(field+" = ?", value)
	}

	if err := query.First(&transaction).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepository) FindAll() ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := r.DB.Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *TransactionRepository) Update(transaction *models.Transaction) error {
	return r.DB.Save(transaction).Error
}

func (r *TransactionRepository) Delete(id uuid.UUID) error {
	return r.DB.Delete(&models.Transaction{}, "id = ?", id.String()).Error
}

func (r *TransactionRepository) DeleteAll() error {
	return r.DB.Delete(&models.Transaction{}).Error
}