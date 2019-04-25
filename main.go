package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hajjboy95/go-user-service/controllers"
	"github.com/hajjboy95/go-user-service/db"
	"github.com/hajjboy95/go-user-service/middleware"
	"github.com/hajjboy95/go-user-service/services"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	dbs := db.GetDb()

	registrationService := services.NewRegistrationService(dbs)
	registrationController := controllers.NewRegistrationController(registrationService)

	loginService := services.NewLoginService(dbs)
	loginController := controllers.NewLoginController(loginService)

	router.HandleFunc("/v1/api/user/new", registrationController.CreateAccount).Methods("POST")
	router.HandleFunc("/v1/api/user/login", loginController.Login).Methods("POST")

	authenticationMiddleWare := middleware.NewAuthenticationMiddleware(
		[]string {"/v1/api/user/new", "/v1/api/user/login"})

	router.Use(authenticationMiddleWare.Authenticate)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		fmt.Print(err)
	}
}