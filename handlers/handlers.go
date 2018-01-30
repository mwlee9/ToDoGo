package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"../models"

	"github.com/julienschmidt/httprouter"
)

// Types - Remember, names must be capital to be exported for the json package to use.

type task struct {
	ID       int
	Name     string
	Body     string
	Priority int
}

var tasks []task

// Home ...
func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	t, err3 := template.ParseFiles("views/home.html")

	if err3 != nil {
		fmt.Println("ERROR3")
	}

	t.Execute(w, "Home")

}

// GetAllTasks ...
func GetAllTasks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows := models.GetAllTasks()
	tasks := tasks[:0]

	for rows.Next() {
		//REMEMBER: using a := will redefine that slice EVERY TIME, if you want to append, to an existing slice you must use = only.
		tempTask := task{}

		err2 := rows.Scan(&tempTask.ID, &tempTask.Name, &tempTask.Body, &tempTask.Priority)

		checkErr(err2)

		tasks = append(tasks, tempTask)

	}
	json.NewEncoder(w).Encode(tasks)
	defer rows.Close()
}

// GetOneTask ...
func GetOneTask(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	params := ps.ByName("id")

	rows := models.GetOneTask(params)

	tempTask := task{}

	for rows.Next() {

		err2 := rows.Scan(&tempTask.ID, &tempTask.Name, &tempTask.Body, &tempTask.Priority)

		checkErr(err2)
	}

	defer rows.Close()

	json.NewEncoder(w).Encode(tempTask)

}

// DeleteOneTask ...
func DeleteOneTask(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Get the id of the record to be deleted.
	deletedRecord := ps.ByName("id")

	// Query for the rows that are going to be deleted, to display before deletion.
	rows := models.GetOneTask(deletedRecord)

	tempTask := task{}

	for rows.Next() {
		err2 := rows.Scan(&tempTask.ID, &tempTask.Name, &tempTask.Body, &tempTask.Priority)

		checkErr(err2)
	}

	// fmt.Fprintf(w, "The following record was deleted: ")
	json.NewEncoder(w).Encode(tempTask)

	// Actually delete the record
	models.DeleteOneTask(deletedRecord)

}

// NewTask ...
func NewTask(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	r.ParseForm()
	name := r.FormValue("name")
	body := r.FormValue("body")
	priority := r.FormValue("priority")

	models.NewTask(name, body, priority)

	t, err3 := template.ParseFiles("views/home.html")

	checkErr(err3)

	t.Execute(w, "Home")

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}

}
