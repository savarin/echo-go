package main

import (
	"database/sql"
	"fmt"
	"log"

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

func (e *EchoLogStore) createTable() {
	db := e.connect()
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS echo_log (id TEXT NOT NULL PRIMARY KEY, message TEXT NOT NULL, created_at TEXT NOT NULL);`)
	if err != nil {
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
	store.createTable()
	store.insertLog("id1", "Hello, World!", "2023-08-18 10:30:00")
	store.printLogs()
}
