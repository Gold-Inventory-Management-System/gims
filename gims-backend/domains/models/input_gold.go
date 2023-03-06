package models

type InputNewGoldInventory struct {
	GoldDetailID uint32 `json:"gold_detail_id"`
	Quantity     int    `json:"quantity"`
	Note         string `json:"note"`
}

type InputNewGoldDetail struct {
	GoldDetailID uint32  `json:"gold_detail_id"`
	Code         string  `json:"code"`
	Type         string  `json:"type"`
	Detail       string  `json:"detail"`
	Weight       float64 `json:"weight"`
	GoldPercent  float64 `json:"gold_percent"`
	GoldSmithFee float64 `json:"gold_smith_fee"`
	Picture      string  `json:"picture"`
	Status       string  `json:"status"`
	Note         string  `json:"note"`
	Quantity     int     `json:"quantity"`
}

type InputSetStatusGoldInventory struct {
	GoldInventoryID []uint32 `json:"gold_inventory_id"`
	Status          string   `json:"status"`
}

type InputSetStatusGoldDetail struct {
	GoldDetailID []uint32 `json:"gold_detail_id"`
	Status       string   `json:"status"`
}

type InputDeleteGoldInventory struct {
	GoldInventoryID []uint32 `json:"gold_inventory_id"`
}

type InputSetTagSerialNumber struct {
	GoldInventoryID uint32 `json:"gold_inventory_id"`
	TagSerialNumber uint64 `json:"tag_serial_number"`
}

type CheckGoldBody struct {
	TagSerialArray []uint64 `json:"tag_serial_array"`
}