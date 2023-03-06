package usecases

import (
	"fmt"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
)

type transactionUsecase struct {
	transactionRepo domains.TransactionRepository
	goldRepo        domains.GoldRepository
}

func NewTransactionUsecase(tr domains.TransactionRepository, gr domains.GoldRepository) domains.TransactionUseCase {
	return &transactionUsecase{transactionRepo: tr, goldRepo: gr}
}

func (tu *transactionUsecase) NewTransactionTypeBuy(transaction *models.Transaction) error {
	transaction.SetTimeNow()
	transaction.BuyPrice = transaction.Price
	transaction.Price = -transaction.Price
	return tu.transactionRepo.InsertNewTransaction(transaction)
}

func (tu *transactionUsecase) NewTransactionTypeSell(transaction *models.Transaction) error {
	goldInventory, err := tu.goldRepo.CheckGoldInventoryByGoldInventoryID(transaction.GoldInventoryID)
	if err != nil {
		return err
	}
	goldDetail, err := tu.goldRepo.QueryGoldDetailByGoldDetailID(goldInventory.GoldDetailID)
	if err != nil {
		return err
	}
	transaction.GoldDetailID = goldInventory.GoldDetailID
	transaction.Weight = goldDetail.Weight
	transaction.SetTimeNow()
	if err := tu.goldRepo.UpdateGoldInventoryIsSold(transaction.GoldInventoryID, 1); err != nil {
		return err
	}
	transaction.SellPrice = transaction.Price
	return tu.transactionRepo.InsertNewTransaction(transaction)
}

func (tu *transactionUsecase) NewTransactionTypeChange(transaction *models.Transaction) error {
	goldInventory, err := tu.goldRepo.CheckGoldInventoryByGoldInventoryID(transaction.GoldInventoryID)
	if err != nil {
		return err
	}
	if _, err := tu.goldRepo.QueryGoldDetailByGoldDetailID(goldInventory.GoldDetailID); err != nil {
		return err
	}
	transaction.GoldDetailID = goldInventory.GoldDetailID
	transaction.Price = transaction.SellPrice - transaction.BuyPrice
	transaction.SetTimeNow()
	if err := tu.goldRepo.UpdateGoldInventoryIsSold(transaction.GoldInventoryID, 1); err != nil {
		return err
	}
	return tu.transactionRepo.InsertNewTransaction(transaction)
}

func (tu *transactionUsecase) RollBackTransaction(transactionID uint32) error {
	transaction, err := tu.transactionRepo.QueryTransactionByTransactionID(transactionID)
	if err != nil {
		return err
	}
	if transaction.TransactionType != "buy" {
		if err := tu.goldRepo.UpdateGoldInventoryIsSold(transaction.GoldInventoryID, 0); err != nil {
			return err
		}
	}
	return tu.transactionRepo.DeleteTransaction(transaction.TransactionID)
}

func (tu *transactionUsecase) appendGoldToTransaction(tjgs []models.TransactionJoinGold) ([]models.TransactionJoinGold, error) {
	if err := tu.goldRepo.AppendGoldDetailToTransactionJoinGold(tjgs); err != nil {
		return tjgs, err
	}
	if err := tu.goldRepo.AppendGoldInventoryToTransactionJoinGold(tjgs); err != nil {
		return tjgs, err
	}
	return tjgs, nil
}

func (tu *transactionUsecase) GetAllTransactionJoinGold() ([]models.TransactionJoinGold, error) {
	tjgs, err := tu.transactionRepo.QueryAllTransaction()
	if err != nil {
		return tjgs, err
	}
	return tu.appendGoldToTransaction(tjgs)
}

func (tu *transactionUsecase) GetTransactionByTransactionType(transactionType string) ([]models.TransactionJoinGold, error) {
	tjgs, err := tu.transactionRepo.QueryTransactionByTransactionType(transactionType)
	if err != nil {
		return tjgs, err
	}
	return tu.appendGoldToTransaction(tjgs)
}

func (tu *transactionUsecase) GetTransactionByTimeInterval(timeRange string) ([]models.TransactionJoinGold, error) {
	tjgs, err := tu.transactionRepo.QueryTransactionByTimeInterval(timeRange)
	if err != nil {
		return tjgs, err
	}
	return tu.appendGoldToTransaction(tjgs)
}

func (tu *transactionUsecase) GetTransactionByDate(date string) ([]models.TransactionJoinGold, error) {
	tjgs, err := tu.transactionRepo.QueryTransactionFromTo(date, date)
	if err != nil {
		return tjgs, err
	}
	return tu.appendGoldToTransaction(tjgs)
}

func (tu *transactionUsecase) GetTransactionFromTo(from, to string) ([]models.TransactionJoinGold, error) {
	tjgs, err := tu.transactionRepo.QueryTransactionFromTo(from, to)
	if err != nil {
		return tjgs, err
	}
	return tu.appendGoldToTransaction(tjgs)
}

func (tu *transactionUsecase) GetReport(interval string) (*models.Report, error) {
	report, err := tu.transactionRepo.MakeReport(interval)
	if err != nil {
		return report, err
	}
	tjgs, err1 := tu.appendGoldToTransaction(report.Transactions)
	report.Transactions = tjgs
	return report, err1
}

func (tu *transactionUsecase) appendGoldToTransactionDashboard(dashboard *models.Dashboard) error {
	if err := tu.goldRepo.AppendGoldDetailToTransactionJoinGold(dashboard.SellTransaction); err != nil {
		return err
	}
	if err := tu.goldRepo.AppendGoldInventoryToTransactionJoinGold(dashboard.SellTransaction); err != nil {
		return err
	}
	if err := tu.goldRepo.AppendGoldDetailToTransactionJoinGold(dashboard.ChangeTransaction); err != nil {
		return err
	}
	if err := tu.goldRepo.AppendGoldInventoryToTransactionJoinGold(dashboard.ChangeTransaction); err != nil {
		return err
	}
	for _, tmp := range dashboard.SellTransaction {
		dashboard.GoldTypeCount[tmp.GoldDetail.Type] += 1
		dashboard.WeightCount[fmt.Sprintf("%.3f", tmp.GoldDetail.Weight)] += 1
		dashboard.UserCount[tmp.Transaction.Username] += 1
	}
	for _, tmp := range dashboard.ChangeTransaction {
		dashboard.GoldTypeCount[tmp.GoldDetail.Type] += 1
		dashboard.WeightCount[fmt.Sprintf("%.3f", tmp.GoldDetail.Weight)] += 1
		dashboard.UserCount[tmp.Transaction.Username] += 1
	}
	return nil
}

func (tu *transactionUsecase) GetDashboard(from, to string) (*models.Dashboard, error) {
	dashboard, err := tu.transactionRepo.MakeDashboard(from, to)
	if err != nil {
		return dashboard, err
	}
	tu.appendGoldToTransactionDashboard(dashboard)
	return dashboard, nil
}