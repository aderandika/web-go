package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Inisial untuk koneksi ke database
var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:@/go_products?parseTime=true")
	if err != nil {
		panic(err)
	}

	log.Println("Database Terhubung")
	DB = db
}
