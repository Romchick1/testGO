package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Импорт драйвера
)

var DB *sql.DB

func InitDB() {
	dsn := "postgres://postgres:qwe1234567890@localhost:5432/product_db?sslmode=disable" // Замените yourpassword
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to DB")
}
