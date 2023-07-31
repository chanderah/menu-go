package config

import (
	"fmt"
	"log"
	"os"

	"github.com/chanderah/menu-go/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func ConnectDb(){
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	// connStr := os.Getenv("DB_URL")
	Host:= os.Getenv("HOST")
	Port:= os.Getenv("PORT")
	DBName:= os.Getenv("DB_NAME")
	DBUsername:= os.Getenv("DB_USERNAME")
	DBPassword:= os.Getenv("DB_PASSWORD")

	dsn:= fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=verify-full TimeZone=Asia/Jakarta", Host, Port, DBUsername, DBPassword, DBName)
	db, err:= gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to the Database!")
	}

	db.AutoMigrate(&model.Post{});
	Db = db
}