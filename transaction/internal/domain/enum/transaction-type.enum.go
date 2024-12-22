package enum

type TransactionType string

const (
	PayTransaction TransactionType = "PAY"
	RefundTransaction TransactionType = "REFUND"
)