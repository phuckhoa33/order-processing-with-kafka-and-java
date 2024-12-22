package postgresql

import (
	"github.com/phuckhoa33/web-crawler/internal/domain/models/base"
)

type Transaction struct {
	*base.AuditModelBase

	UserID    string `json:"user_id" gorm:"type:varchar(255)"`
	OrderID   string `json:"order_id" gorm:"type:varchar(255)"`
	Amount    float64   `json:"amount"`
	Action    string    `json:"action"` // e.g., "pay", "trade"
}