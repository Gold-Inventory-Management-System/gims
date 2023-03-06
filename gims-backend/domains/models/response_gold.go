package models

type GoldJoin struct {
	GoldDetail    *GoldDetail    `json:"gold_detail"`
	GoldInventory *GoldInventory `json:"gold_inventory"`
}

type CheckGoldResponse struct {
	Result            string     `json:"result"`
	MissFrontGold     []GoldJoin `json:"miss_front_gold"`
	SafeGold          []GoldJoin `json:"safe_gold"`
	TagEmptyFrontGold []GoldJoin `json:"tag_empty_front_gold"`
}