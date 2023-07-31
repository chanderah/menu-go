package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib" // library bindings for pgx
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	"github.com/joho/godotenv"
)

var (
	Db *sqlx.DB
)

var schema =
`	CREATE TABLE IF NOT EXISTS tutorial.posts (
		id BIGSERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		content TEXT NOT NULL,
		created_at timestamp default NOW()
	);
`

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
	db, err:= sqlx.Connect("pgx", dsn);
	if err != nil {
		log.Fatal("Failed to connect to the Database!", err.Error())
	}
	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal("Failed to execute the Schema.", err.Error())
	}

	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
	Db = db;
}