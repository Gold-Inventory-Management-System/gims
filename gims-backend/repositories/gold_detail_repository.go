package repositories

import (
	"context"
	"fmt"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/infrastructure/database"
	"github.com/jackc/pgx/v4/pgxpool"
	"gorm.io/gorm"
)

type goldRepository struct {
	ctx context.Context
	db  *pgxpool.Pool
	gormDb *gorm.DB
}

func NewGoldRepository(db *pgxpool.Pool, gormDb *gorm.DB) domains.GoldRepository {
	return &goldRepository{ctx: context.Background(), db: db, gormDb: gormDb}
}

func (gr *goldRepository) NewGoldDetail(g *models.GoldDetail) (uint32, error) {
	id := database.GenerateUUID()
	insertGoldDetailSql := `INSERT INTO gold_details (gold_detail_id, code, type, detail, weight, gold_percent, gold_smith_fee, picture, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`
	if _, err := gr.db.Exec(gr.ctx, insertGoldDetailSql, id, g.Code, g.Type, g.Detail, g.Weight, g.GoldPercent, g.GoldSmithFee, g.Picture, "normal"); err != nil {
		return 0, err
	}
	return id, nil
}

func (gr *goldRepository) CheckGoldDetail(g *models.GoldDetail) error {
	queryGoldDetailByDetail := `
		SELECT * 
		FROM gold_details
		WHERE code = $1 AND type = $2 AND detail = $3 AND weight = $4 AND gold_percent = $5 AND gold_smith_fee = $6 AND status = $7;
	`
	rows, err := gr.db.Query(gr.ctx, queryGoldDetailByDetail, g.Code, g.Type, g.Detail, g.Weight, g.GoldPercent, g.GoldSmithFee, "normal")
	if err != nil {
		return err
	}
	detail := new(models.GoldDetail)
	for rows.Next() {
		if err = rows.Scan(&detail.GoldDetailID, &detail.Code, &detail.Type, &detail.Detail, &detail.Weight, &detail.GoldPercent, &detail.GoldSmithFee, &detail.Picture, &detail.Status); err != nil {
			return err
		}
	}
	if detail.Type != "" {
		return fmt.Errorf("this detail is already exist see code %s", detail.Code)
	}
	return nil
}

func (gr *goldRepository) QueryAllGoldDetail() ([]models.GoldDetail, error) {
	var res []models.GoldDetail
	queryAllGoldDetail := `SELECT * FROM gold_details WHERE status = $1;`
	rows, err := gr.db.Query(gr.ctx, queryAllGoldDetail, "normal")
	if err != nil {
		return res, err
	}
	for rows.Next() {
		var detail models.GoldDetail
		if err = rows.Scan(&detail.GoldDetailID, &detail.Code, &detail.Type, &detail.Detail, &detail.Weight, &detail.GoldPercent, &detail.GoldSmithFee, &detail.Picture, &detail.Status); err != nil {
			return res, err
		}
		res = append(res, detail)
	}
	return res, nil
}

func (gr *goldRepository) QueryGoldDetailByGoldDetailID(goldDetailID uint32) (*models.GoldDetail, error) {
	var detail models.GoldDetail
	queryGoldDetailByCodeSql := `
		SELECT *
		FROM gold_details
		WHERE gold_detail_id = $1 AND status = $2;
	`
	rows, err := gr.db.Query(gr.ctx, queryGoldDetailByCodeSql, goldDetailID, "normal")
	if err != nil {
		return &detail, err
	}
	for rows.Next() {
		if err = rows.Scan(&detail.GoldDetailID, &detail.Code, &detail.Type, &detail.Detail, &detail.Weight, &detail.GoldPercent, &detail.GoldSmithFee, &detail.Picture, &detail.Status); err != nil {
			return &detail, err
		}
	}
	if detail.GoldDetailID == 0 {
		return &detail, fmt.Errorf("detail not found")
	}
	return &detail, nil
}

func (gr *goldRepository) QueryGoldDetailByCode(code string) ([]models.GoldDetail, error) {
	var res []models.GoldDetail
	queryGoldDetailByCodeSql := `
		SELECT *
		FROM gold_details
		WHERE code = $1 AND status = $2;
	`
	rows, err := gr.db.Query(gr.ctx, queryGoldDetailByCodeSql, code, "normal")
	if err != nil {
		return res, err
	}
	for rows.Next() {
		var detail models.GoldDetail
		if err = rows.Scan(&detail.GoldDetailID, &detail.Code, &detail.Type, &detail.Detail, &detail.Weight, &detail.GoldPercent, &detail.GoldSmithFee, &detail.Picture, &detail.Status); err != nil {
			return res, err
		}
		res = append(res, detail)
	}
	return res, nil
}

func checkGoldDetail(g, goldDetail models.GoldDetail) bool {
	if g.Code != "" {
		if g.Code != goldDetail.Code {
			return false
		}
	}
	if g.Type != "" {
		if g.Type != goldDetail.Type {
			return false
		}
	}
	if g.Weight != 0 {
		if g.Weight != goldDetail.Weight {
			return false
		}
	}
	if g.GoldPercent != 0 {
		if g.GoldPercent != goldDetail.GoldPercent {
			return false
		}
	}
	if g.GoldSmithFee != 0 {
		if g.GoldSmithFee != goldDetail.GoldSmithFee {
			return false
		}
	}
	if g.Detail != "" {
		if g.Detail != goldDetail.Detail {
			return false
		}
	}
	return true
}

func (gr *goldRepository) QueryGoldDetailByDetail(g *models.GoldDetail) ([]models.GoldDetail, error) {
	var res []models.GoldDetail
	queryAllGoldDetailSql := `SELECT * FROM gold_details WHERE status = $1;`
	rows, err := gr.db.Query(gr.ctx, queryAllGoldDetailSql, "normal")
	if err != nil {
		return res, err
	}
	for rows.Next() {
		var goldDetail models.GoldDetail
		if err = rows.Scan(&goldDetail.GoldDetailID, &goldDetail.Code, &goldDetail.Type, &goldDetail.Detail, &goldDetail.Weight, &goldDetail.GoldPercent, &goldDetail.GoldSmithFee, &goldDetail.Picture, &goldDetail.Status); err != nil {
			return res, err
		}
		if !checkGoldDetail(*g, goldDetail){
			continue
		}
		res = append(res, goldDetail)
	}
	return res, nil
}

func (gr *goldRepository) UpdateGoldDetail(goldDetail *models.GoldDetail) error{
	updateGoldDetailSql := `
		UPDATE gold_details
		SET code = $1, type = $2, detail = $3, weight = $4, gold_percent = $5, gold_smith_fee = $6, picture = $7
		WHERE gold_detail_id = $8;
	`
	_, err := gr.db.Exec(gr.ctx, updateGoldDetailSql, goldDetail.Code, goldDetail.Type, goldDetail.Detail, goldDetail.Weight, goldDetail.GoldPercent, goldDetail.GoldSmithFee, goldDetail.Picture, goldDetail.GoldDetailID)
	return err
}

func (gr *goldRepository) SetStatusGoldDetail(goldDetailID uint32, setStatus string) error{
	updateStatusGoldDetailToDelete := `
		UPDATE gold_details
		SET status = $1
		WHERE gold_detail_id = $2;
	`
	_, err := gr.db.Exec(gr.ctx, updateStatusGoldDetailToDelete, setStatus, goldDetailID)
	return err
}

func (gr *goldRepository) AppendGoldDetailToTransactionJoinGold(transactionJoinGold []models.TransactionJoinGold) error {
	queryGoldDetailByGoldDetailIDSql := `SELECT * FROM gold_details WHERE gold_detail_id = $1;`
	for i := range transactionJoinGold {
		if transactionJoinGold[i].Transaction.TransactionType == "buy" {
			continue
		}
		row := gr.db.QueryRow(gr.ctx, queryGoldDetailByGoldDetailIDSql, transactionJoinGold[i].Transaction.GoldDetailID)
		if err := row.Scan(&transactionJoinGold[i].GoldDetail.GoldDetailID, &transactionJoinGold[i].GoldDetail.Code, &transactionJoinGold[i].GoldDetail.Type, &transactionJoinGold[i].GoldDetail.Detail, &transactionJoinGold[i].GoldDetail.Weight, &transactionJoinGold[i].GoldDetail.GoldPercent, &transactionJoinGold[i].GoldDetail.GoldSmithFee, &transactionJoinGold[i].GoldDetail.Picture, &transactionJoinGold[i].GoldDetail.Status); err != nil {
			return err
		}
	}
	return nil
}