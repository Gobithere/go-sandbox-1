package db

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func InitDB(connstr string) {
	var err error
	DB, err = sql.Open("postgres", connstr)
	if err != nil {
		panic("DB Connection Failed!")
	} else {
		fmt.Println("DB Connection Success!")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(4)
}

func CreateTables() {
	createEventTableSQL := `CREATE TABLE IF NOT EXISTS events (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	dateTime TIMESTAMP NOT NULL,
	user_id INTEGER
	)`

	_, err := DB.Exec(createEventTableSQL)
	if err != nil {
		panic(err)
	}

}
