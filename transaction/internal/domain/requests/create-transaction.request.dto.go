package requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/phuckhoa33/web-crawler/internal/domain/enum"
)

type CreateTransactionRequestDto struct {
	UserID     string                    `json:"userId" validate:"required" example:"ME232323232"`
	OrderID      string                    `json:"orderId" validate:"required" example:"OR3232323"`
	Action       enum.TransactionType `json:"action" validate:"required" example:"PAY"`
}

func (r *CreateTransactionRequestDto) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.UserID, validation.Required, validation.Length(1, 255)),
		validation.Field(&r.OrderID, validation.Required, validation.Length(1, 255)),
		validation.Field(&r.Action, validation.Required, validation.In(
			enum.PayTransaction,
			enum.RefundTransaction,
		)),
	)
}