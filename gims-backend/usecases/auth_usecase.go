package usecases

import (
	"fmt"
	"time"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func (uu *userUseCase) SignIn(u *models.User) (*models.User, string, error){
	user, err := uu.userRepo.FindUser(u.Username)
	if err != nil {
		return user, "", err
	}
	if user.Username != u.Username{
		return user, "", fmt.Errorf("can't find user")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err != nil{
		return user, "", fmt.Errorf("wrong password")
	}
	accessToken, err := createToken(user)
	return user, accessToken, err
}

func createToken(u *models.User) (string, error){
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = u.Username
	claims["role"] = u.Role
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()
	accessToken, err := token.SignedString([]byte("secret")) //มาเปลี่ยนอีกที
    return accessToken, err
}