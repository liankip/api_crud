package db

import (
	"database/sql"
	"embed"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
)

var embedMigrations embed.FS

func RunMigrations() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to open sql connection: %v", err)
	}

	goose.SetBaseFS(embedMigrations)

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
}
