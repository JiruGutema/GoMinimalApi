// Package config contains cinfiguration settings for the application. like db connectin
package config


import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func ConnectDatabase() {
	fmt.Println("connecting to the database...")
	dsn := "postgres://postgres:1441@localhost:5432/gominimalapi"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	DB = db
	fmt.Println("Database Connected Successfully!")
}
