package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

type EchoLogStore struct{}

func (e *EchoLogStore) connect() *sql.DB {
	db, err := sql.Open("sqlite3", "echo.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func (e *EchoLogStore) migrate() {
	m, err := migrate.New(
		"file://db/migrations",
		"sqlite3://echo.db")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}

func (e *EchoLogStore) insertLog(id, message, createdAt string) {
	db := e.connect()
	_, err := db.Exec("INSERT INTO echo_log (id, message, created_at) VALUES (?, ?, ?)", id, message, createdAt)
	if err != nil {
		log.Fatal(err)
	}
}

func (e *EchoLogStore) printLogs() {
	db := e.connect()
	rows, err := db.Query("SELECT id, message, created_at FROM echo_log")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, message, createdAt string
		rows.Scan(&id, &message, &createdAt)
		fmt.Printf("ID: %s, Message: %s, Created At: %s\n", id, message, createdAt)
	}
}

func main() {
	store := &EchoLogStore{}

	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		store.migrate()
	} else {
		store.insertLog("id1", "Hello, World!", "2023-08-18 10:30:00")
		store.printLogs()
	}
}
