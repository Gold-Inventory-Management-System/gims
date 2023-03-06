package models

import "time"

type GoldInventory struct {
	GoldInventoryID uint32    `json:"gold_inventory_id"`
	GoldDetailID    uint32    `json:"gold_detail_id"`
	Status          string    `json:"status"` //safe or front
	DateIn          time.Time `json:"date_in"`
	DateSold        time.Time `json:"date_sold"`
	Note            string    `json:"note"`
	IsSold          int       `json:"is_sold"`
	TagSerialNumber uint64    `json:"tag_serail_number"`
}
