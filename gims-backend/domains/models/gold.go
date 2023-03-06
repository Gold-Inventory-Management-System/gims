package models

type GoldDetailJoinInventory struct {
	GoldDetailID uint32          `json:"gold_detail_id"`
	Code         string          `json:"code"`
	Type         string          `json:"type"`
	Detail       string          `json:"detail"`
	Weight       float64         `json:"weight"`
	GoldPercent  float64         `json:"gold_percent"`
	GoldSmithFee float64         `json:"gold_smith_fee"`
	Picture      string          `json:"picture"`
	Status       string          `json:"status"`
	Inventories  []GoldInventory `json:"inventories"`
}