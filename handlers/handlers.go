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

type animal struct {
	ID      int
	Name    string
	Species string
}

var animals []animal

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
	animals := animals[:0]

	for rows.Next() {
		//REMEMBER: using a := will redefine that slice EVERY TIME, if you want to append, to an existing slice you must use = only.
		tempAnimal := animal{}

		err2 := rows.Scan(&tempAnimal.ID, &tempAnimal.Name, &tempAnimal.Species)

		checkErr(err2)

		animals = append(animals, tempAnimal)

	}
	json.NewEncoder(w).Encode(animals)
	defer rows.Close()
}

// GetOneTask ...
func GetOneTask(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	params := ps.ByName("id")

	rows := models.GetOneTask(params)

	tempAnimal := animal{}

	for rows.Next() {

		err2 := rows.Scan(&tempAnimal.ID, &tempAnimal.Name, &tempAnimal.Species)

		checkErr(err2)
	}

	defer rows.Close()

	json.NewEncoder(w).Encode(tempAnimal)

}

// DeleteOneTask ...
func DeleteOneTask(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Get the id of the record to be deleted.
	deletedRecord := ps.ByName("id")

	// Query for the rows that are going to be deleted, to display before deletion.
	rows := models.GetOneTask(deletedRecord)

	tempAnimal := animal{}

	for rows.Next() {
		err2 := rows.Scan(&tempAnimal.ID, &tempAnimal.Name, &tempAnimal.Species)

		checkErr(err2)
	}

	// fmt.Fprintf(w, "The following record was deleted: ")
	json.NewEncoder(w).Encode(tempAnimal)

	// Actually delete the record
	models.DeleteOneTask(deletedRecord)

}

// NewTask ...
func NewTask(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	r.ParseForm()
	name := r.FormValue("name")
	species := r.FormValue("species")

	models.NewTask(name, species)

	t, err3 := template.ParseFiles("views/home.html")

	if err3 != nil {
		fmt.Println("ERROR3")
	}

	t.Execute(w, "Home")

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}

}
