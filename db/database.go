package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	db_name := ""
	db_host := ""
	db_port := ""
	db_user := ""
	db_pass := ""

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", db_host, db_port, db_user, db_pass, db_name)
	DB, err := sql.Open("postgres", dsn)

	if err != nil {
		fmt.Println("Database connection failed :", err)
	} else {
		fmt.Println("Connected to Postgres Database")
	}
	return DB
}
