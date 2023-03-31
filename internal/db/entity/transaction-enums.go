package entity

type TransactionStatus = string

const (
	TransactionStatusDraft      TransactionStatus = "DRAFT"
	TransactionStatusFailed     TransactionStatus = "FAILED"
	TransactionStatusRejected   TransactionStatus = "REJECTED"
	TransactionStatusSuccess    TransactionStatus = "SUCCESS"
	TransactionStatusProcessing TransactionStatus = "PROCESSING"
	TransactionStatusCreated    TransactionStatus = "CREATED"
)
