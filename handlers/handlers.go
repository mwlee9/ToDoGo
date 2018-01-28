package handlers

import (
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

	rows := models.GetAllTasks()

	tempAnimal := animal{ID: 0, Name: "", Species: ""}

	for rows.Next() {
		err2 := rows.Scan(&tempAnimal.ID, &tempAnimal.Name, &tempAnimal.Species)
		fmt.Println(tempAnimal.ID, tempAnimal.Name, tempAnimal.Species)
		if err2 != nil {
			fmt.Println("ERROR2")
			fmt.Println(err2)
		}

	}

	t, err3 := template.ParseFiles("views/home.html")

	if err3 != nil {
		fmt.Println("ERROR3")
	}

	t.Execute(w, "Home")

}

// GetOneTask ...
func GetOneTask(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	params := ps.ByName("id")

	rows := models.GetOneTask(params)

	tempAnimal1 := animal{ID: 0, Name: "", Species: ""}

	for rows.Next() {
		err2 := rows.Scan(&tempAnimal1.ID, &tempAnimal1.Name, &tempAnimal1.Species)
		fmt.Println(tempAnimal1.ID, tempAnimal1.Name, tempAnimal1.Species)
		if err2 != nil {
			fmt.Println("ERROR2")
			fmt.Println(err2)
		}

	}

}

// DeleteTask ...
func DeleteOneTask(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	models.DeleteOneTask(ps.ByName("id"))

}

// NewTask ...
func NewTask(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	r.ParseForm()
	name := r.FormValue("name")
	species := r.FormValue("species")

	models.NewTask(name, species)

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
