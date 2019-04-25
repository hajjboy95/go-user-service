package controllers

import (
	"encoding/json"
	"github.com/hajjboy95/go-user-service/models"
	"github.com/hajjboy95/go-user-service/services"
	"github.com/hajjboy95/go-user-service/utils"
	"net/http"
)

type RegistrationController struct {
	registrationService services.Registrable
}

func NewRegistrationController(registrationService services.Registrable) *RegistrationController {
	return &RegistrationController{
		registrationService: registrationService,
	}
}

func (rc *RegistrationController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		message := utils.Message(false, err.Error())
		utils.Respond(http.StatusConflict, w, message)
		return
	}

	resp, err := rc.registrationService.RegisterUser(user)

	if err != nil {
		message := utils.Message(false, err.Error())
		utils.Respond(http.StatusConflict, w, message)
		return
	}

	utils.Respond(http.StatusCreated,  w, resp)
}

func (rc *RegistrationController) ControllerName() string {
	return "registrationController"
}