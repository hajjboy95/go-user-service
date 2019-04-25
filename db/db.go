package db

import (
	"fmt"
	"github.com/hajjboy95/go-user-service/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type DB struct {

}

var db *gorm.DB

func init() {
	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
		log.Panic(e) //not for production
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		log.Panic(err) //not for production
	}

	db = conn
	db.Debug().AutoMigrate(&models.User{})
}

func GetDb() *gorm.DB {
	return db
}