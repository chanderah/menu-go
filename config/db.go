package config

import (
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib" // library bindings for pgx
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDb(){
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file!")
	}

	// connStr := os.Getenv("DB_URL")
	DBHost:= os.Getenv("DB_HOST")
	DBPort:= os.Getenv("DB_PORT")
	DBName:= os.Getenv("DB_NAME")
	DBUsername:= os.Getenv("DB_USERNAME")
	DBPassword:= os.Getenv("DB_PASSWORD")
	SSL:= os.Getenv("SSL")
	TimeZone:= os.Getenv("TimeZone")

	dsn:= fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s", DBHost, DBPort, DBUsername, DBPassword, DBName, SSL, TimeZone)
	_, err = sqlx.Connect("pgx", dsn);
	if err != nil {
		log.Fatal("Failed to connect to the Database!", err.Error())
	} else {
		fmt.Println("Connection is established!.")
	}

	// data := model.Post{}
    // rows, err := db.Queryx("SELECT * FROM tutorial.posts")
    // for rows.Next() {
    //     err := rows.StructScan(&data)
    //     if err != nil {
    //         log.Fatalln(err)
    //     }
    //     fmt.Printf("%#v\n", data)
    // }

	// dsn:= fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s", DBHost, DBPort, DBUsername, DBPassword, DBName, SSL, TimeZone)
	// db, err:= gorm.Open(postgres.Open(dsn), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Info),
	// 	NamingStrategy: schema.NamingStrategy{
	// 		// TablePrefix: "tb_",
	// 		TablePrefix: "tutorial.",
	// 		SingularTable: false,
	// 	},
	// })

	// if err != nil {
	// 	log.Fatal("Failed to connect to the Database!")
	// }
	// db.AutoMigrate(&model.Post{});
	// Db = db;
}