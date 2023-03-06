package usecases

import (
	"fmt"
	"strconv"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
)

type goldUseCase struct {
	goldRepo domains.GoldRepository
}

func NewGoldUseCase(gr domains.GoldRepository) domains.GoldUseCase {
	return &goldUseCase{gr}
}

func (gu *goldUseCase) ConvertIDStringToUint32(id string) (uint32, error) {
	goldDetailID, err := strconv.ParseUint(id, 10, 32)
	return uint32(goldDetailID), err
}

func (gu *goldUseCase) ConvertIDStringToUint64(id string) (uint64, error) {
	goldDetailID, err := strconv.ParseUint(id, 10, 64)
	return uint64(goldDetailID), err
}

func (gu *goldUseCase) NewGold(g *models.InputNewGoldDetail) error {
	newGoldDetail := &models.GoldDetail{Code: g.Code, Type: g.Type, Detail: g.Detail, Weight: g.Weight, GoldPercent: g.GoldPercent, GoldSmithFee: g.GoldSmithFee, Picture: g.Picture, Status: g.Status}
	if err := gu.goldRepo.CheckGoldDetail(newGoldDetail); err != nil {
		return err
	}
	id, err := gu.goldRepo.NewGoldDetail(newGoldDetail)
	if err != nil {
		return err
	}
	newGoldInventory := &models.InputNewGoldInventory{GoldDetailID: id, Note: g.Note, Quantity: g.Quantity}
	err = gu.goldRepo.NewGoldInventory(newGoldInventory)
	if err != nil {
		return err
	}
	return nil
}

func (gu *goldUseCase) AddGold(newGoldInventory *models.InputNewGoldInventory) error {
	err := gu.goldRepo.NewGoldInventory(newGoldInventory)
	return err
}

func (gu *goldUseCase) GetAllGoldDetail() ([]models.GoldDetail, error) {
	res, err := gu.goldRepo.QueryAllGoldDetail()
	return res, err
}

func (gu *goldUseCase) FindGoldDetailByCode(code string) ([]models.GoldDetail, error) {
	details, err := gu.goldRepo.QueryGoldDetailByCode(code)
	return details, err
}

func (gu *goldUseCase) FindGoldDetailByDetail(g *models.GoldDetail) ([]models.GoldDetail, error) {
	details, err := gu.goldRepo.QueryGoldDetailByDetail(g)
	return details, err
}

func (gu *goldUseCase) EditGoldDetail(goldDetail *models.GoldDetail) error {
	return gu.goldRepo.UpdateGoldDetail(goldDetail)
}

func (gu *goldUseCase) GetAllGoldDetailJoinInventory() ([]models.GoldDetailJoinInventory, error) {
	return gu.goldRepo.QueryAllGoldDetailJoinInventory()
}

func (gu *goldUseCase) SetStatusGoldDetailToDelete(goldDetailID []uint32) error {
	for _, id := range goldDetailID {
		if err := gu.goldRepo.SetStatusGoldDetail(id, "delete"); err != nil {
			return err
		}
	}
	return nil
}

func (gu *goldUseCase) SetStatusGoldDetailToNormal(goldDetailID uint32) error {
	return gu.goldRepo.SetStatusGoldDetail(goldDetailID, "normal")
}

func (gu *goldUseCase) SetStatusGoldInventory(goldInventoryIDs []uint32, status string) error {
	for _, id := range goldInventoryIDs{
		if err := gu.goldRepo.UpdateGoldInventoryStatus(id, status); err != nil {
			return err
		}
	}
	return nil
}

func (gu *goldUseCase) GetGoldDetailJoinInventoryByDetail(g *models.GoldDetail) ([]models.GoldDetailJoinInventory, error) {
	return gu.goldRepo.QueryGoldDetailJoinInventoryByDetail(g)
}

func (gu *goldUseCase) GetGoldDetailByGoldDetailID(goldDetailID uint32) (*models.GoldDetail, error) {
	return gu.goldRepo.QueryGoldDetailByGoldDetailID(goldDetailID)
}

func (gu *goldUseCase) DeleteGoldInventoryByIDArray(ids []uint32) error {
	for _, id := range ids {
		if err := gu.goldRepo.DeleteGoldInventoryByID(id); err != nil {
			return err
		}
	}
	return nil
}

func (gu *goldUseCase) SetTagSerialNumberGoldInventory(input *models.InputSetTagSerialNumber) error {
	goldInventory := gu.goldRepo.QueryGoldByTagSerialNumber(input.TagSerialNumber)
	if goldInventory.GoldInventoryID != 0 {
		return fmt.Errorf("this tag serial number is already used")
	}
	return gu.goldRepo.SetTagSerialNumberGoldInventory(input.GoldInventoryID, input.TagSerialNumber)
}

func (gu *goldUseCase) QueryGoldByTagSerialNumber(tagSerialnumber uint64) (*models.GoldJoin, error) {
	var gold *models.GoldJoin = new(models.GoldJoin)
	var err error
	gold.GoldInventory = gu.goldRepo.QueryGoldByTagSerialNumber(tagSerialnumber)
	goldDetailID := gold.GoldInventory.GoldDetailID
	if goldDetailID == 0{
		return gold, fmt.Errorf("gold not found")
	}
	gold.GoldDetail, err = gu.goldRepo.QueryGoldDetailByGoldDetailID(goldDetailID)
	return gold, err
}

func (gu *goldUseCase) QueryGoldJoinByGoldInventoryIDArray(ids []uint32) []models.GoldJoin{
	res := []models.GoldJoin{}
	for _, id := range ids {
		var goldJoin models.GoldJoin
		goldInventory, _ := gu.goldRepo.QueryGoldInventoryByGoldInventoryID(id)
		goldDetail, _ := gu.goldRepo.QueryGoldDetailByGoldDetailID(goldInventory.GoldDetailID)
		goldJoin.GoldDetail = goldDetail
		goldJoin.GoldInventory = goldInventory
		res = append(res, goldJoin)
	}
	return res
}

func (gu *goldUseCase) CheckFrontGold(arrayOfSerialNumber []uint64) (*models.CheckGoldResponse, error) {
	res := new(models.CheckGoldResponse)
	var miss []uint32
	var notin []uint32
	var tagEmpty []uint32
	resQueryGoldByTag := gu.goldRepo.QueryGoldInventoryByTagSerialNumberArray(arrayOfSerialNumber)
	_, mapFrontGoldID, err := gu.goldRepo.QueryAllGoldInventoryStatusFront()
	if err != nil {
		return res, err
	}
	//for debug
	// fmt.Println(mapFrontGoldID)
	for _, inventory := range resQueryGoldByTag {
		if mapFrontGoldID[inventory.GoldInventoryID] == "found" {
			mapFrontGoldID[inventory.GoldInventoryID] = "check"
		}else if mapFrontGoldID[inventory.GoldInventoryID] == "" {
			mapFrontGoldID[inventory.GoldInventoryID] = "notin"
		}
	}
	for k, v := range mapFrontGoldID {
		if v == "found" {
			miss = append(miss, k)
		}else if v == "notin" {
			notin = append(notin, k)
		}else if v == "tag empty" {
			tagEmpty = append(tagEmpty, k)
		}
	}

	res.MissFrontGold = gu.QueryGoldJoinByGoldInventoryIDArray(miss) //ทองที่ไม่ครบ
	res.SafeGold = gu.QueryGoldJoinByGoldInventoryIDArray(notin) //ทองที่ส่งมามีแท็กแต่เป็นทองที่ไม่อยู่่ที่หน้าร้าน
	res.TagEmptyFrontGold = gu.QueryGoldJoinByGoldInventoryIDArray(tagEmpty) //ทองที่อยู่หน้าร้านแต่ไม่มีแท็ก

	if len(res.MissFrontGold) != 0 || len(res.SafeGold) != 0 || len(res.TagEmptyFrontGold) != 0 {
		if len(res.MissFrontGold) != 0 {
			res.Result += "There are miss front gold"
		}
		if len(res.SafeGold) != 0 {
			if len(res.Result) != 0 {
				res.Result += " And there are safe gold"
			}else {
				res.Result += "There are safe gold"
			}
		}
		if len(res.TagEmptyFrontGold) != 0 {
			if len(res.Result) != 0 {
				res.Result += " And there are tag empty front gold"
			}else {
				res.Result += "There are tag empty front gold"
			}
		}
	}else {
		res.Result = "Complete"
	}
	return res, err
}

func (gu *goldUseCase) GetAllFrontGold() ([]models.GoldJoin, error) {
	goldJoins, err := gu.goldRepo.QueryFrontGold()
	return goldJoins, err
}