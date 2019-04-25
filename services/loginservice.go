package services

import (
	"errors"
	"github.com/hajjboy95/go-user-service/common"
	"github.com/hajjboy95/go-user-service/models"
	"github.com/hajjboy95/go-user-service/utils"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Loginable interface {
	Login(user *models.User) (map[string]interface{}, error)
}

type LoginService struct {
	db *gorm.DB
}

func (ls *LoginService) ServiceName() string {
	return "loginService"
}

func (ls *LoginService) Login(user *models.User) (map[string]interface{}, error) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}
	return ls.loginUser(user)
}

func (ls *LoginService) loginUser(user *models.User) (map[string]interface{}, error) {
	u := &models.User{}
	err := ls.db.Table("users").Where("email = ?", user.Email).First(u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("Email not registered, Please Register")
		}
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if err != nil {
		return nil, err
	}

	tokenString, err := common.CreateTokenStringForUser(user)

	if err != nil {
		return nil, err
	}
	u.Token = tokenString
	resp := utils.Message(true, "Login Success")
	resp["user"] = u

	return resp, nil
}

func NewLoginService(db *gorm.DB) *LoginService {
	return &LoginService{ db: db}
}
