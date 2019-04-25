package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strings"
)

type User struct {
	gorm.Model
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token";sql:"-"`
}

func (u *User) Validate() error {
	if !strings.Contains(u.Email, "@") {
		err := errors.New("Invalid Email")
		return err
	}

	if len(u.Password) < 6 {
		err := errors.New("Invalid Password, Password must be greater the 6 chars")
		return err
	}

	return nil
}