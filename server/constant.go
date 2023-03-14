package server

const (
	BASE_URL         = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	REQUEST_TIMEOUT  = 200
	DB_TIMEOUT       = 10
	DB_DIR           = "./server/db"
	DB_FILE          = "./server/db/cotacao.db"
	SQL_CREATE_TABLE = `
		CREATE TABLE IF NOT EXISTS cotacao (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			codein TEXT,
			name TEXT,
			high REAL,
			low REAL,
			varbid REAL,
			pctchange REAL,
			bid REAL,
			ask REAL,
			timestamp INTEGER,
			createdate TEXT
		);
	`
	SQL_INSERT = `
		INSERT INTO cotacao 
			(codein, name, high, low, varbid, pctchange, bid, ask, timestamp, createdate) 
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`
)
