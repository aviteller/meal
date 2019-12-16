package models

import (
	"database/sql"
	"fmt"
)

// Creates all tables in the sql database
func InitTables() {

	GetDB().Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY autoincrement, name text, email text, password text, token text, created_at text, updated_at text, deleted int DEFAULT 0)")

	GetDB().Exec("CREATE TABLE IF NOT EXISTS children (id INTEGER PRIMARY KEY autoincrement, user_id int, name text, age int, gender text, created_at text, updated_at text, deleted int DEFAULT 0)")

}

//GetDB opens connection to sqlite db
func GetDB() *sql.DB {
	//:databaselocked.sqlite?cache=shared&mode=rwc
	database, err := sql.Open("sqlite3", "./meals.db")
	if err != nil {
		fmt.Println(err)
	}

	return database
}
