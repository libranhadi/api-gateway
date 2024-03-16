package config

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	connectionString := "postgres://postgres:postgres@postgresql:5432/service_users?sslmode=disable"
	count := 0
	for {
		if count > 3 {
			break
		}

		count += 1
		var err error
		db, err = sql.Open("postgres", connectionString)
		if err != nil {
			fmt.Println("Failed to connect to database. Retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
			continue
		}

		db.SetConnMaxLifetime(time.Minute * 10)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)

		if err := db.Ping(); err != nil {
			fmt.Println("Failed to ping database. Retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}
	fmt.Println("Connected to database...")
}

func NewPostgresContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func GetPostgresDB() *sql.DB {
	return db
}
