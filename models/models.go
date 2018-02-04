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

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS dash (id SERIAL PRIMARY KEY, category TEXT, task TEXT, priority INTEGER);")
	checkErr(err)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS work (id SERIAL PRIMARY KEY, category TEXT, task TEXT, priority INTEGER);")
	checkErr(err)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS weekend (id SERIAL PRIMARY KEY, category TEXT, task TEXT, priority INTEGER);")
	checkErr(err)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS groceries (id SERIAL PRIMARY KEY, category TEXT, task TEXT, priority INTEGER);")
	checkErr(err)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS resolutions (id SERIAL PRIMARY KEY, category TEXT, task TEXT, priority INTEGER);")
	checkErr(err)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS hobby (id SERIAL PRIMARY KEY, category TEXT, task TEXT, priority INTEGER);")
	checkErr(err)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS design (id SERIAL PRIMARY KEY, category TEXT, task TEXT, priority INTEGER);")

	checkErr(err)

	defer db.Close()

}

// GetAllTasks ...
func GetAllTasks(tblName string) *sql.Rows {
	db := InitDatabase()

	var rows *sql.Rows
	var err error

	switch tblName {

	case "dash":
		rows, err = db.Query("SELECT * FROM dash ORDER BY priority, category")

	case "work":
		rows, err = db.Query("SELECT * FROM work ORDER BY priority, category")

	case "weekend":
		rows, err = db.Query("SELECT * FROM weekend ORDER BY priority, category")
	case "groceries":
		rows, err = db.Query("SELECT * FROM groceries ORDER BY priority, category")

	case "resolutions":
		rows, err = db.Query("SELECT * FROM resolutions ORDER BY priority, category")

	case "hobby":
		rows, err = db.Query("SELECT * FROM hobby ORDER BY priority, category")

	case "design":
		rows, err = db.Query("SELECT * FROM design ORDER BY priority, category")

	}

	checkErr(err)

	defer db.Close()

	return rows
}

// GetOneTask ...
//This GetOneTask func is needed in order to properly select a rec to delete!
func GetOneTask(param string, tblName string) *sql.Rows {
	db := InitDatabase()

	var rows *sql.Rows
	var err error

	switch tblName {

	case "dash":
		rows, err = db.Query("SELECT * FROM dash WHERE id = $1", param)
	case "work":
		rows, err = db.Query("SELECT * FROM work WHERE id = $1", param)
	case "weekend":
		rows, err = db.Query("SELECT * FROM weekend WHERE id = $1", param)
	case "groceries":
		rows, err = db.Query("SELECT * FROM groceries WHERE id = $1", param)
	case "resolutions":
		rows, err = db.Query("SELECT * FROM resolutions WHERE id = $1", param)
	case "hobby":
		rows, err = db.Query("SELECT * FROM hobby WHERE id = $1", param)
	case "design":
		rows, err = db.Query("SELECT * FROM design WHERE id = $1", param)

	}

	checkErr(err)

	defer db.Close()

	return rows
}

// DeleteOneTask ...
func DeleteOneTask(param string, tblName string) sql.Result {
	db := InitDatabase()

	var stmt *sql.Stmt
	var err error

	switch tblName {

	case "dash":
		stmt, err = db.Prepare("DELETE FROM dash WHERE id = $1;")
	case "work":
		stmt, err = db.Prepare("DELETE FROM work WHERE id = $1;")
	case "weekend":
		stmt, err = db.Prepare("DELETE FROM weekend WHERE id = $1;")
	case "groceries":
		stmt, err = db.Prepare("DELETE FROM groceries WHERE id = $1;")
	case "resolutions":
		stmt, err = db.Prepare("DELETE FROM resolutions WHERE id = $1;")
	case "hobby":
		stmt, err = db.Prepare("DELETE FROM hobby WHERE id = $1;")
	case "design":
		stmt, err = db.Prepare("DELETE FROM design WHERE id = $1;")

	}

	checkErr(err)

	res, err := stmt.Exec(param)

	checkErr(err)

	defer db.Close()

	return res

}

// NewTask ...
func NewTask(category string, task string, priority string, tblName string) string {

	db := InitDatabase()

	var stmt *sql.Stmt
	var err error

	switch tblName {

	case "dash":
		stmt, err = db.Prepare("INSERT INTO dash (category, task, priority) values ($1,$2,$3)")
	case "work":
		stmt, err = db.Prepare("INSERT INTO work (category, task, priority) values ($1,$2,$3)")
	case "weekend":
		stmt, err = db.Prepare("INSERT INTO weekend (category, task, priority) values ($1,$2,$3)")
	case "groceries":
		stmt, err = db.Prepare("INSERT INTO groceries (category, task, priority) values ($1,$2,$3)")
	case "resolutions":
		stmt, err = db.Prepare("INSERT INTO resolutions (category, task, priority) values ($1,$2,$3)")
	case "hobby":
		stmt, err = db.Prepare("INSERT INTO hobby (category, task, priority) values ($1,$2,$3)")
	case "design":
		stmt, err = db.Prepare("INSERT INTO design (category, task, priority) values ($1,$2,$3)")

	}

	checkErr(err)

	_, err = stmt.Exec(category, task, priority)

	checkErr(err)

	defer db.Close()

	return tblName

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
