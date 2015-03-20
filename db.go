package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

type Db struct {
	connection *sql.DB
}

func (db *Db) initdb() error {
	_, err := db.connection.Exec(`
		CREATE TABLE IF NOT EXISTS Todos (id INTEGER PRIMARY KEY NOT NULL, content TEXT NOT NULL, modificationtime TIMESTAMP NOT NULL)
	`)
	if err != nil {
		return err
	}
	if debug {
		log.Printf("Table 'Todos' OK")
	}
	return nil
}

func (db *Db) InsertTodo(todo string) error {
	tx, err := db.connection.Begin()
	if err != nil {
		log.Fatal(err)
		return err
	}
	stmt, err := tx.Prepare(`INSERT INTO Todos(content, modificationtime) VALUES(?, ?)`)
	defer stmt.Close()
	_, err = stmt.Exec(todo, time.Now())
	if err != nil {
		log.Fatal(err)
		return err
	}
	tx.Commit()
	return nil
}

func (db *Db) Close() {
	db.connection.Close()
}

func OpenDatabase(dbpath string) (*Db, error) {
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		return nil, err
	}
	dbobj := Db{
		connection: db,
	}
	err = dbobj.initdb()
	if err != nil {
		return nil, err
	}
	return &dbobj, nil
}
