package common

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/hajjboy95/go-user-service/db"
	"github.com/hajjboy95/go-user-service/models"
	"github.com/jinzhu/gorm"
	"os"
)

func CreateTokenStringForUser(user *models.User) (string, error) {
	tk := &models.Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte(os.Getenv("token_password")))
	return "Bearer " + tokenString, err
}

func CheckIfUserExists(user *models.User) error {
	temp := &models.User{}

	err := db.GetDb().Table("users").Where("email = ?", user.Email).First(temp).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		err := errors.New("Record not found")
		return err
	}

	if temp.Email != "" {
		err := errors.New("Email Already in Use")
		return err
	}
	return nil
}