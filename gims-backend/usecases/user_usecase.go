package usecases

import (
	"fmt"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct{
	userRepo domains.UserRepository
}

func NewUserUseCase(ur domains.UserRepository) domains.UserUseCase{
	return &userUseCase{
		userRepo: ur,
	}
}

func (uu *userUseCase) GenerateHashPasswordAndReplaceInUserModel(u *models.User) error{
	p := []byte(u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(p, bcrypt.DefaultCost)
	if err != nil{
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (uu *userUseCase) GenerateHashPassword(password string) (string, error){
	p := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(p, bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func (uu *userUseCase) RegisterAdmin(u *models.User, secret string) error{
	if secret != "BaBaBuBu"{
		return fmt.Errorf("invalid secret")
	}
	if err := uu.userRepo.CheckUsername(u.Username); err != nil{
		return err
	}
	u.Role = "admin"
	if err := uu.GenerateHashPasswordAndReplaceInUserModel(u); err != nil {
		return err
	}
	err := uu.userRepo.InsertUser(u)
	return err
}

func (uu *userUseCase) Register(u *models.User) error{
	if err := uu.userRepo.CheckUsername(u.Username); err != nil{
		return err
	}
	if err := uu.GenerateHashPasswordAndReplaceInUserModel(u); err != nil {
		return err
	}
	err := uu.userRepo.InsertUser(u)
	return err
}

func (uu* userUseCase) QueryAllUser() ([]models.UserResponse, error){
	return uu.userRepo.QueryAllUser()
}

func (uu *userUseCase) RenameUsername(oldUsername, newUsername string) error{
	if err := uu.userRepo.CheckUsername(newUsername); err != nil {
		return err
	}
	err := uu.userRepo.UpdateUsername(oldUsername, newUsername)
	return err
}

func (uu *userUseCase) UpdatePassword(username, password string) error{
	hashPassword, err := uu.GenerateHashPassword(password)
	if err != nil {
		return err
	}
	err = uu.userRepo.UpdatePassword(username, hashPassword)
	return err
}

func (uu *userUseCase) DeleteUser(username string) error {
	return uu.userRepo.DeleteUser(username)
}