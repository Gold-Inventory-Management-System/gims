package repositories

import (
	"fmt"
	"time"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/infrastructure/database"
)

func (gr *goldRepository) NewGoldInventory(newGoldInventory *models.InputNewGoldInventory) error {
	insertGoldInventorySql := `INSERT INTO gold_inventories (gold_inventory_id, gold_detail_id, status, date_in, date_sold, note, is_sold, tag_serial_number) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`
	for i := 0; i < newGoldInventory.Quantity; i++ {
		loc, _ := time.LoadLocation("Asia/Jakarta")
		if _, err := gr.db.Exec(gr.ctx, insertGoldInventorySql, database.GenerateUUID(), newGoldInventory.GoldDetailID, "safe", time.Now().In(loc), time.Time{}, newGoldInventory.Note, 0, 0); err != nil {
			return err
		}
	}
	return nil
}

func (gr *goldRepository) UpdateGoldInventoryStatus(goldInventoryID uint32, status string) error {
	updateGoldInventoryStatus := `UPDATE gold_inventories SET status = $1 WHERE gold_inventory_id = $2`
	_, err := gr.db.Exec(gr.ctx, updateGoldInventoryStatus, status, goldInventoryID)
	return err
}

func (gr *goldRepository) UpdateGoldInventoryIsSold(goldInventoryID uint32, isSold int) error {
	updateGoldInventoryIsSold := `UPDATE gold_inventories SET is_sold = $1 WHERE gold_inventory_id = $2`
	_, err := gr.db.Exec(gr.ctx, updateGoldInventoryIsSold, isSold, goldInventoryID)
	return err
}

func (gr *goldRepository) CheckGoldInventoryByGoldInventoryID(id uint32) (models.GoldInventory, error) {
	var goldInventory models.GoldInventory
	queryGoldInventoryByGoldInventoryIDSql := `SELECT * FROM gold_inventories WHERE gold_inventory_id = $1 AND is_sold = $2;`
	rows, err := gr.db.Query(gr.ctx, queryGoldInventoryByGoldInventoryIDSql, id, 0)
	if err != nil {
		return goldInventory, err
	}
	for rows.Next() {
		if err = rows.Scan(&goldInventory.GoldInventoryID, &goldInventory.GoldDetailID, &goldInventory.Status, &goldInventory.DateIn, &goldInventory.DateSold, &goldInventory.Note, &goldInventory.IsSold, &goldInventory.TagSerialNumber); err != nil {
			return goldInventory, err
		}
	}
	if goldInventory.GoldInventoryID == 0 {
		return goldInventory, fmt.Errorf("product not found")
	}
	return goldInventory, err
}

func (gr *goldRepository) AppendGoldInventoryToTransactionJoinGold(transactionJoinGold []models.TransactionJoinGold) error {
	queryGoldInventoryByGoldInventoryIDSql := `SELECT * FROM gold_inventories WHERE gold_inventory_id = $1;`
	for i := range transactionJoinGold {
		if transactionJoinGold[i].Transaction.TransactionType == "buy" {
			continue
		}
		row := gr.db.QueryRow(gr.ctx, queryGoldInventoryByGoldInventoryIDSql, transactionJoinGold[i].Transaction.GoldInventoryID)
		if err := row.Scan(&transactionJoinGold[i].GoldInventory.GoldInventoryID, &transactionJoinGold[i].GoldInventory.GoldDetailID, &transactionJoinGold[i].GoldInventory.Status, &transactionJoinGold[i].GoldInventory.DateIn, &transactionJoinGold[i].GoldInventory.DateSold, &transactionJoinGold[i].GoldInventory.Note, &transactionJoinGold[i].GoldInventory.IsSold, &transactionJoinGold[i].GoldInventory.TagSerialNumber); err != nil {
			return err
		}
	}
	return nil
}

func (gr *goldRepository) QueryAllGoldInventoryByGoldDetailID(gold_detail_id uint32) ([]models.GoldInventory, error) {
	queryGoldInventoryByGoldDetailIDSql := `SELECT * FROM gold_inventories WHERE gold_detail_id = ? AND is_sold = ?;`
	var inventories []models.GoldInventory
	rowsGoldInventories, err := gr.gormDb.Raw(queryGoldInventoryByGoldDetailIDSql, gold_detail_id, 0).Rows()
	if err != nil {
		return inventories, err
	}
	for rowsGoldInventories.Next() {
		var inventory models.GoldInventory
		if err = gr.gormDb.ScanRows(rowsGoldInventories, &inventory); err != nil {
			return inventories, err
		}
		inventories = append(inventories, inventory)
	}
	return inventories, nil
}

func (gr *goldRepository) DeleteGoldInventoryByID(id uint32) error {
	deleteGoldInventoryByID := `DELETE FROM gold_inventories WHERE gold_inventory_id = $1;`
	_, err := gr.db.Exec(gr.ctx, deleteGoldInventoryByID, id)
	return err
}

func (gr *goldRepository) QueryGoldInventoryByGoldInventoryID(id uint32) (*models.GoldInventory, error) {
	var goldInventory *models.GoldInventory
	queryGoldInventoryByGoldInventoryIDSql := `SELECT * from gold_inventories WHERE gold_inventory_id = ?;`
	gr.gormDb.Raw(queryGoldInventoryByGoldInventoryIDSql, id).Scan(&goldInventory)
	if goldInventory.GoldInventoryID == 0 {
		return goldInventory, fmt.Errorf("inventory not found")
	}
	return goldInventory, nil
}

func (gr *goldRepository) QueryAllGoldInventoryStatusFront() ([]models.GoldInventory, map[uint32]string, error) {
	var inventories []models.GoldInventory
	queryGoldInventoryStatusFront := `SELECT * from gold_inventories WHERE status = 'front' AND is_sold = 0;`
	mapGoldInventoryID := map[uint32]string{}
	rowsGoldInventories, err := gr.gormDb.Raw(queryGoldInventoryStatusFront).Rows()
	if err != nil {
		return inventories, mapGoldInventoryID, err
	}
	for rowsGoldInventories.Next() {
		var inventory models.GoldInventory
		if err = gr.gormDb.ScanRows(rowsGoldInventories, &inventory); err != nil {
			return inventories, mapGoldInventoryID, err
		}
		inventories = append(inventories, inventory)
		if inventory.TagSerialNumber == 0 {
			mapGoldInventoryID[inventory.GoldInventoryID] = "tag empty"
			continue
		}
		mapGoldInventoryID[inventory.GoldInventoryID] = "found"
	}
	return inventories, mapGoldInventoryID, nil
}