package domains

import "github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"

type TransactionUseCase interface {
	NewTransactionTypeBuy(transaction *models.Transaction) error
	NewTransactionTypeSell(transaction *models.Transaction) error
	NewTransactionTypeChange(transaction *models.Transaction) error
	RollBackTransaction(transactionID uint32) error
	GetAllTransactionJoinGold() ([]models.TransactionJoinGold, error)
	GetTransactionByTransactionType(transactionType string) ([]models.TransactionJoinGold, error)
	GetTransactionByTimeInterval(timeRange string) ([]models.TransactionJoinGold, error)
	GetTransactionByDate(date string) ([]models.TransactionJoinGold, error)
	GetTransactionFromTo(from, to string) ([]models.TransactionJoinGold, error)
	GetReport(interval string) (*models.Report, error)
	GetDashboard(from, to string) (*models.Dashboard, error)
}

type TransactionRepository interface {
	InsertNewTransaction(transaction *models.Transaction) error
	DeleteTransaction(transactionID uint32) error
	QueryTransactionByTransactionID(transactionID uint32) (*models.Transaction, error)
	QueryAllTransaction() ([]models.TransactionJoinGold, error)
	QueryTransactionByTransactionType(transactionType string) ([]models.TransactionJoinGold, error)
	QueryTransactionByTimeInterval(timeRange string) ([]models.TransactionJoinGold, error)
	QueryTransactionFromTo(from, to string) ([]models.TransactionJoinGold, error)
	MakeReport(interval string) (*models.Report, error)
	MakeDashboard(from, to string) (*models.Dashboard, error)
}