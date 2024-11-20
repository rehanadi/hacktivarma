package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {

	databaseName := "phar1"
	dsn := fmt.Sprintf("root:@tcp(127.0.0.1:3306)/%s?parseTime=true", databaseName)
	DB, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println("Database connection failed :", err)
	} else {
		fmt.Println("Connected to MySQL")
	}
	return DB
}
