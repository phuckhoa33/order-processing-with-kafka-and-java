package postgresql

import (
	"github.com/google/uuid"
	"github.com/phuckhoa33/web-crawler/internal/domain/models/base"
)

type Transaction struct {
	*base.AuditModelBase

	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid"`
	OrderID   uuid.UUID `json:"order_id" gorm:"type:uuid"`
	Amount    float64   `json:"amount"`
	Action    string    `json:"action"` // e.g., "pay", "trade"
}