package server

import (
	"database/sql"
	"log"
	"os"
)

func initDB() {
	os.MkdirAll(DB_DIR, 0755)
	os.Create(DB_FILE)
	db, err := sql.Open("sqlite3", DB_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec(SQL_CREATE_TABLE)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("database started")
}
