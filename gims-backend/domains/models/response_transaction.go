package models

type TransactionJoinGold struct {
	Transaction   Transaction   `json:"transaction"`
	GoldDetail    GoldDetail    `json:"gold_detail"`
	GoldInventory GoldInventory `json:"gold_inventory"`
}

type Report struct {
	Transactions       []TransactionJoinGold `json:"transactions"`
	TotalPrice         float64               `json:"total_price"`
	IncomePrice        float64               `json:"income_price"`
	OutcomePrice       float64               `json:"outcome_price"`
	BuyPrice           float64               `json:"buy_price"`  //sum of all transaction type buy
	SellPrice          float64               `json:"sell_price"` //sum of all transaction type sell
	ChangeIncomePrice  float64               `json:"change_income_price"`
	ChangeOutcomePrice float64               `json:"change_outcome_price"`
	TotalChangePrice   float64               `json:"total_change_price"`
}

type Dashboard struct {
	SellTransaction    []TransactionJoinGold `json:"sell_transaction"`
	BuyTransaction     []TransactionJoinGold `json:"buy_transaction"`
	ChangeTransaction  []TransactionJoinGold `json:"change_transaction"`
	TotalPrice         float64               `json:"total_price"`
	IncomePrice        float64               `json:"income_price"`
	OutcomePrice       float64               `json:"outcome_price"`
	BuyPrice           float64               `json:"buy_price"`  //sum of all transaction type buy
	SellPrice          float64               `json:"sell_price"` //sum of all transaction type sell
	ChangeIncomePrice  float64               `json:"change_income_price"`
	ChangeOutcomePrice float64               `json:"change_outcome_price"`
	TotalChangePrice   float64               `json:"total_change_price"`
	GoldTypeCount      map[string]int        `json:"gold_type_count"`
	WeightCount        map[string]int       `json:"weight_count"`
	UserCount          map[string]int        `json:"user_count"`
}
