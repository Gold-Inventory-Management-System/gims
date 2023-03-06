package domains

import "github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"

type GoldUseCase interface{
	NewGold(goldDetail *models.InputNewGoldDetail) error
	ConvertIDStringToUint32(id string) (uint32, error)
	ConvertIDStringToUint64(id string) (uint64, error)
	AddGold(newGoldInventory *models.InputNewGoldInventory) error
	GetAllGoldDetail() ([]models.GoldDetail, error)
	FindGoldDetailByCode(code string) ([]models.GoldDetail, error)
	FindGoldDetailByDetail(g *models.GoldDetail) ([]models.GoldDetail, error)
	EditGoldDetail(goldDetail *models.GoldDetail) error
	GetAllGoldDetailJoinInventory() ([]models.GoldDetailJoinInventory, error)
	SetStatusGoldDetailToDelete(goldDetailID []uint32) error
	SetStatusGoldDetailToNormal(goldDetailID uint32) error
	SetStatusGoldInventory(goldInventoryIDs []uint32, status string) error
	GetGoldDetailJoinInventoryByDetail(g *models.GoldDetail) ([]models.GoldDetailJoinInventory, error)
	GetGoldDetailByGoldDetailID(goldDetailID uint32) (*models.GoldDetail, error)
	DeleteGoldInventoryByIDArray(ids []uint32) error
	SetTagSerialNumberGoldInventory(input *models.InputSetTagSerialNumber) error
	QueryGoldByTagSerialNumber(tagSerialnumber uint64) (*models.GoldJoin, error)
	QueryGoldJoinByGoldInventoryIDArray(ids []uint32) []models.GoldJoin
	CheckFrontGold(arrayOfSerialNumber []uint64) (*models.CheckGoldResponse, error)
	GetAllFrontGold() ([]models.GoldJoin, error)
}

type GoldRepository interface{
	NewGoldDetail(g *models.GoldDetail) (uint32, error)
	NewGoldInventory(newGoldInventory *models.InputNewGoldInventory) error
	QueryAllGoldDetail() ([]models.GoldDetail, error)
	QueryGoldDetailByGoldDetailID(goldDetailID uint32) (*models.GoldDetail, error)
	QueryGoldDetailByCode(code string) ([]models.GoldDetail, error)
	CheckGoldDetail(g *models.GoldDetail) error
	QueryGoldDetailByDetail(g *models.GoldDetail) ([]models.GoldDetail, error)
	UpdateGoldDetail(goldDetail *models.GoldDetail) error
	QueryAllGoldDetailJoinInventory() ([]models.GoldDetailJoinInventory, error)
	SetStatusGoldDetail(goldDetailID uint32, setStatus string) error
	UpdateGoldInventoryStatus(goldInventoryID uint32, status string) error
	CheckGoldInventoryByGoldInventoryID(id uint32) (models.GoldInventory, error)
	UpdateGoldInventoryIsSold(goldInventoryID uint32, isSold int) error
	AppendGoldDetailToTransactionJoinGold(transactionJoinGold []models.TransactionJoinGold) error
	AppendGoldInventoryToTransactionJoinGold(transactionJoinGold []models.TransactionJoinGold) error
	QueryAllGoldInventoryByGoldDetailID(gold_detail_id uint32) ([]models.GoldInventory, error)
	QueryGoldDetailJoinInventoryByDetail(g *models.GoldDetail) ([]models.GoldDetailJoinInventory, error)
	DeleteGoldInventoryByID(id uint32) error
	QueryGoldByTagSerialNumber(serialNumber uint64) *models.GoldInventory
	SetTagSerialNumberGoldInventory(id uint32, serialNumber uint64) error 
	QueryAllGoldInventoryStatusFront() ([]models.GoldInventory, map[uint32]string, error) 
	QueryGoldInventoryByTagSerialNumberArray(tagSerialNumberArray []uint64) []models.GoldInventory
	QueryGoldInventoryByGoldInventoryID(id uint32) (*models.GoldInventory, error)
	QueryFrontGold() ([]models.GoldJoin, error)
}