package util

import (
	"fmt"
	"log"
	"os"

	"github.com/chanderah/menu-go/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func sync(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.ProductOption{})
	DB = db
}

func GetConnectionMySql() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}

	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tb_",
			SingularTable: true,
			// TablePrefix:   "tutorial.",
			// SingularTable: false,
		},
	})

	if err != nil {
		log.Fatal("Failed to connect to the Database!")
	}
	sync(db)
}

func GetConnectionPostgres() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}
	// connStr := os.Getenv("DB_URL")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB_NAME")
	DBUsername := os.Getenv("DB_USERNAME")
	DBPassword := os.Getenv("DB_PASSWORD")
	SSL := os.Getenv("SSL")
	TimeZone := os.Getenv("TimeZone")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s", DBHost, DBPort, DBUsername, DBPassword, DBName, SSL, TimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix:   "tb_",
			// SingularTable: true,
			TablePrefix:   "tutorial.",
			SingularTable: false,
		},
	})

	if err != nil {
		log.Fatal("Failed to connect to the Database!")
	}
	sync(db)
}
