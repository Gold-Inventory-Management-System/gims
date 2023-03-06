package repositories

import "github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"

func (gr *goldRepository) QueryGoldByTagSerialNumber(serialNumber uint64) *models.GoldInventory {
	var goldInventory models.GoldInventory
	queryGoldByTagSerialNumberSql := `SELECT * FROM gold_inventories WHERE tag_serial_number = ? AND is_sold = ?;`
	gr.gormDb.Raw(queryGoldByTagSerialNumberSql, serialNumber, 0).Scan(&goldInventory)
	return &goldInventory
}

func (gr *goldRepository) SetTagSerialNumberGoldInventory(id uint32, serialNumber uint64) error {
	updateTagSerialNumberGoldInventorySql := `UPDATE gold_inventories SET tag_serial_number = $1 WHERE gold_inventory_id = $2 and is_sold = $3;`
	_, err := gr.db.Exec(gr.ctx, updateTagSerialNumberGoldInventorySql, serialNumber, id, 0)
	return err
}

func (gr *goldRepository) QueryGoldInventoryByTagSerialNumberArray(tagSerialNumberArray []uint64) []models.GoldInventory {
	queryGoldInventoryByTagSerialNumberSql := `SELECT * from gold_inventories WHERE tag_serial_number = ? AND is_sold = 0;`
	var inventories []models.GoldInventory
	for _, tagSerialNumber := range tagSerialNumberArray {
		var goldInventory models.GoldInventory
		gr.gormDb.Raw(queryGoldInventoryByTagSerialNumberSql, tagSerialNumber).Scan(&goldInventory)
		if goldInventory.GoldDetailID != 0 {
			inventories = append(inventories, goldInventory)
		}
	}
	return inventories
}