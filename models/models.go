package models

import (
	"database/sql"
	"fmt"
)

// Database
func InitDatabase() *sql.DB {
	conn := "postgres://swyreijf:hlR3e6UqP7YEsy6nq_BIChyRE8SPINoP@nutty-custard-apple.db.elephantsql.com:5432/swyreijf"
	db, err := sql.Open("postgres", conn)

	checkErr(err)

	return db

}

// CreateTable ...
func CreateTable() {

	db := InitDatabase()

	// _, err := db.Exec("CREATE TABLE IF NOT EXISTS animals (id SERIAL PRIMARY KEY, name TEXT, species TEXT);")
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS animals (id SERIAL PRIMARY KEY, name TEXT, species TEXT);")

	if err != nil {
		fmt.Println("ERROR")
	}

	defer db.Close()

}

// GetAllTasks ...
func GetAllTasks() *sql.Rows {
	db := InitDatabase()

	rows, err := db.Query("SELECT * FROM animals")

	if err != nil {
		fmt.Println("ERROR1")
	}

	defer db.Close()

	return rows
}

// GetOneTask ...
func GetOneTask(params string) *sql.Rows {
	db := InitDatabase()
	rows, err := db.Query("SELECT * FROM animals WHERE id =" + `'` + params + `'`)

	checkErr(err)

	defer db.Close()

	return rows
}

// DeleteOneTask ...
func DeleteOneTask(params string) sql.Result {
	db := InitDatabase()
	stmt, err := db.Prepare("DELETE FROM animals WHERE id = $1;")

	checkErr(err)

	res, err1 := stmt.Exec(params)

	checkErr(err1)

	defer db.Close()

	return res

}

// NewTask ...
func NewTask(name string, species string) {

	db := InitDatabase()

	stmt, err := db.Prepare("INSERT INTO ANIMALS (name, species) values ($1,$2)")
	checkErr(err)

	_, err = stmt.Exec(name, species)

	checkErr(err)

	defer db.Close()

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
