package controllers

import (
	"encoding/json"
	"github.com/hajjboy95/go-user-service/models"
	"github.com/hajjboy95/go-user-service/services"
	"github.com/hajjboy95/go-user-service/utils"
	"net/http"
)

type LoginController struct {
	loginService services.Loginable
}

func NewLoginController(loginService services.Loginable) *LoginController {
	return &LoginController{
		loginService: loginService,
	}
}
func (lg *LoginController)Login(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		message := utils.Message(false, err.Error())
		utils.Respond(http.StatusBadRequest, w, message)
		return
	}

	resp, err := lg.loginService.Login(user)

	if err != nil {
		message := utils.Message(false, err.Error())
		utils.Respond(http.StatusUnauthorized, w, message)
		return
	}
	utils.Respond(http.StatusOK, w, resp)
}

func (lg *LoginController) ControllerName() string {
	return "loginController"
}