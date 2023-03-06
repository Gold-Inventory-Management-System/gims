package repositories

import (
	"context"
	"fmt"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"gorm.io/gorm"
)

type userRepository struct{
	ctx context.Context
	db *pgxpool.Pool
	gormDb *gorm.DB
}

func NewUserRepository(db *pgxpool.Pool, gormDb *gorm.DB) domains.UserRepository{
	return &userRepository{ctx: context.Background(), db: db, gormDb: gormDb}
}

func (ur *userRepository) InsertUser(u *models.User) error{
	insertUserSql := `INSERT INTO users (username, password, role) VALUES ($1, $2, $3)`
	if _, err := ur.db.Exec(ur.ctx, insertUserSql, u.Username, string(u.Password), u.Role); err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CheckUsername(username string) error{
	checkUsernameSql := `SELECT * FROM users WHERE username = $1;`
	rows, err := ur.db.Query(ur.ctx, checkUsernameSql, username);
	if err != nil {
		return err
	}
	u := new(models.User)
	for rows.Next(){
		if err = rows.Scan(&u.ID, &u.Username, &u.Password, &u.Role); err != nil{
			return err
		}
	}
	if u.Username != "" {
		return fmt.Errorf("username already used")
	}
	return nil
}

func (ur *userRepository) FindUser(username string) (*models.User, error) {
	u := new(models.User)
	checkUsernameSql := `SELECT * FROM users WHERE username = $1;`
	rows, err := ur.db.Query(ur.ctx, checkUsernameSql, username);
	if err != nil {
		return u, err
	}
	for rows.Next(){
		if err = rows.Scan(&u.ID, &u.Username, &u.Password, &u.Role); err != nil{
			return u, err
		}
	}
	return u, nil
}

func (ur *userRepository) QueryAllUser() ([]models.UserResponse, error) {
	var users []models.UserResponse
	queryAllUserSql := `SELECT username, role FROM users;`
	rows, err := ur.db.Query(ur.ctx, queryAllUserSql);
	if err != nil {
		return users, err
	}
	for rows.Next(){
		var u models.UserResponse
		if err = rows.Scan(&u.Username, &u.Role); err != nil{
			return users, err
		}
		if u.Role == "admin"{
			continue
		}
		users = append(users, u)
	}
	return users, nil
}

func (ur *userRepository) UpdateUsername(oldUsername, newUsername string) error{
	updateUsernameSql := `UPDATE users SET username = $1 WHERE username = $2;`
	_, err := ur.db.Exec(ur.ctx, updateUsernameSql, newUsername, oldUsername)
	return err
}

func (ur *userRepository) UpdatePassword(username, newHashPassword string) error{
	updatePasswordSql := `UPDATE users SET password = $1 WHERE username = $2;`
	_, err := ur.db.Exec(ur.ctx, updatePasswordSql, newHashPassword, username)
	return err
}

func (ur *userRepository) DeleteUser(username string) error {
	deleteUserSql := `DELETE FROM users WHERE username = $1;`
	_, err := ur.db.Exec(ur.ctx, deleteUserSql, username)
	return err
}