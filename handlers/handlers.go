package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/mwlee9/todogoml/models"

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

// TblName ...
// Allows data passage between functions for determine which table a webpage is on
var TblName string

// ##################################Render Pages##############################################

// Dash ...
func Dash(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	TblName = "dash"

	t, err := template.ParseFiles("views/dash.html", "partials/head.html", "partials/foot.html", "partials/footer.html")

	checkErr(err)

	t.Execute(w, "dash")

}

// Work ...
func Work(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	TblName = "work"

	t, err := template.ParseFiles("views/work.html", "partials/head.html", "partials/foot.html", "partials/footer.html")

	checkErr(err)

	t.Execute(w, "work")

}

// Weekend ...
func Weekend(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	TblName = "weekend"

	t, err := template.ParseFiles("views/weekend.html", "partials/head.html", "partials/foot.html", "partials/footer.html")

	checkErr(err)

	t.Execute(w, "weekend")

}

// Grocery ...
func Grocery(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	TblName = "grocery"

	t, err := template.ParseFiles("views/grocery.html", "partials/head.html", "partials/foot.html", "partials/footer.html")

	checkErr(err)

	t.Execute(w, "grocery")

}

// Resolutions ...
func Resolutions(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	TblName = "resolutions"

	t, err := template.ParseFiles("views/resolutions.html", "partials/head.html", "partials/foot.html", "partials/footer.html")

	checkErr(err)

	t.Execute(w, "resolutions")

}

// ###########################################################################################

// GetAllTasks ...
func GetAllTasks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows := models.GetAllTasks(TblName)
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
//This GetOneTask func is needed in order to properly select a rec to delete!
func GetOneTask(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	params := ps.ByName("id")

	rows := models.GetOneTask(params, TblName)

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
	rows := models.GetOneTask(deletedRecord, TblName)

	tempTask := task{}

	for rows.Next() {
		err2 := rows.Scan(&tempTask.ID, &tempTask.Name, &tempTask.Body, &tempTask.Priority)

		checkErr(err2)
	}

	// fmt.Fprintf(w, "The following record was deleted: ")
	json.NewEncoder(w).Encode(tempTask)

	// Actually delete the record
	models.DeleteOneTask(deletedRecord, TblName)

}

// NewTask ...
func NewTask(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	r.ParseForm()
	category := r.FormValue("category")
	task := r.FormValue("task")
	priority := r.FormValue("priority")

	// returns the route name (also named to the table name for convenience)
	tblName := models.NewTask(category, task, priority, TblName)

	viewFp := "views/" + tblName + ".html"
	t, err := template.ParseFiles(viewFp, "partials/head.html", "partials/foot.html", "partials/footer.html")

	checkErr(err)

	// Remember, for partials each html file must be named properly.
	t.Execute(w, tblName)

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}

}
