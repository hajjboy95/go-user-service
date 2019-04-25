package services

import (
	"github.com/hajjboy95/go-user-service/common"
	"github.com/hajjboy95/go-user-service/models"
	"github.com/hajjboy95/go-user-service/utils"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type Registrable interface {
	RegisterUser(user *models.User) (map[string]interface{}, error)
}

type RegistrationService struct {
	db *gorm.DB
}

func (rs *RegistrationService) ServiceName() string {
	return "RegistrationService"
}

func (rs *RegistrationService) RegisterUser(user *models.User) (map[string]interface{}, error) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}

	err = common.CheckIfUserExists(user)
	if err != nil {
		return nil, err
	}

	return rs.createUser(user)
}

func (rs *RegistrationService) createUser(user *models.User) (map[string]interface{}, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	if err != nil {
		return nil, err
	}

	rs.db.Create(user)
	if user.ID <= 0 {
		return nil, errors.New("Failed to create an account")
	}

	tokenString, err := common.CreateTokenStringForUser(user)
	if err != nil {
		return nil, err
	}
	user.Token = tokenString
	resp := utils.Message(true, "User Registered")
	resp["user"] = user

	return resp, nil
}

func NewRegistrationService(db *gorm.DB) *RegistrationService{
	return &RegistrationService{db: db}
}
