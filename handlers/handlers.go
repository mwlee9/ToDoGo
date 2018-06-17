package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/mwlee9/todogoml/models"

	"github.com/julienschmidt/httprouter"

	"github.com/microcosm-cc/bluemonday"
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

// Allows access through basic verification on homepage
var AllowAccess int

// #################################################Serve Favicon#####################

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "golang_gopher.ico")
}

// #################################################Render Pages#####################

// Dash ...
func Dash(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	TblName = "dash"

	t, err := template.ParseFiles("views/dash.html", "partials/head.html", "partials/foot.html", "partials/footer.html")

	checkErr(err)

	t.Execute(w, "dash")

}

// Work ...
func Work(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if AllowAccess == 1 {
		TblName = "work"

		t, err := template.ParseFiles("views/work.html", "partials/head.html", "partials/foot.html", "partials/footer.html")

		checkErr(err)

		t.Execute(w, "work")

	}

}

// Weekend ...
func Weekend(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if AllowAccess == 1 {
		TblName = "weekend"

		t, err := template.ParseFiles("views/weekend.html", "partials/head.html", "partials/foot.html", "partials/footer.html")

		checkErr(err)

		t.Execute(w, "weekend")
	}

}

// Groceries ...
func Groceries(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if AllowAccess == 1 {
		TblName = "groceries"

		t, err := template.ParseFiles("views/groceries.html", "partials/head.html", "partials/foot.html", "partials/footer.html")

		checkErr(err)

		t.Execute(w, "groceries")
	}

}

// Resolutions ...
func Resolutions(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	TblName = "resolutions"

	t, err := template.ParseFiles("views/resolutions.html", "partials/head.html", "partials/foot.html", "partials/footer.html")

	checkErr(err)

	t.Execute(w, "resolutions")

}

// Hobby ...
func Hobby(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if AllowAccess == 1 {
		TblName = "hobby"

		t, err := template.ParseFiles("views/hobby.html", "partials/head.html", "partials/foot.html", "partials/footer.html")

		checkErr(err)

		t.Execute(w, "hobby")
	}

}

// Design ...
func Design(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if AllowAccess == 1 {
		TblName = "design"

		t, err := template.ParseFiles("views/design.html", "partials/head.html", "partials/foot.html", "partials/footer.html")

		checkErr(err)

		t.Execute(w, "design")
	}

}

// ###########################################################################################

// Login ...
func LoginForm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	AllowAccess = 0
	r.ParseForm()
	usernameIn := r.FormValue("username")
	passwordIn := r.FormValue("password")
	p := bluemonday.UGCPolicy()
	username := p.Sanitize(usernameIn)
	password := p.Sanitize(passwordIn)

	if username == os.Getenv("TODOGOML_USERNAME") && password == os.Getenv("TODOGOML_PASSWORD") {
		AllowAccess = 1
	}

	TblName = "dash"

	t, err := template.ParseFiles("views/dash.html", "partials/head.html", "partials/foot.html", "partials/footer.html")

	checkErr(err)

	t.Execute(w, "dash")

}

// Login ...
func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	TblName = "dash"

	t, err := template.ParseFiles("views/login.html", "partials/head.html", "partials/foot.html", "partials/footer.html")

	checkErr(err)

	t.Execute(w, "login")

}

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
	deletedRecordId := ps.ByName("id")

	// Query for the rows that are going to be deleted, to display before deletion.
	rows := models.GetOneTask(deletedRecordId, TblName)

	tempTask := task{}

	for rows.Next() {
		err2 := rows.Scan(&tempTask.ID, &tempTask.Name, &tempTask.Body, &tempTask.Priority)

		checkErr(err2)
	}

	// fmt.Fprintf(w, "The following record was deleted: ")
	json.NewEncoder(w).Encode(tempTask)

	// Actually delete the record
	models.DeleteOneTask(deletedRecordId, TblName)

}

// NewTask ...
func NewTask(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	r.ParseForm()
	categoryIn := r.FormValue("category")
	taskIn := r.FormValue("task")
	priorityIn := r.FormValue("priority")

	p := bluemonday.UGCPolicy()
	category := p.Sanitize(categoryIn)
	task := p.Sanitize(taskIn)
	priority := p.Sanitize(priorityIn)

	// returns the route name (also named to the table name for convenience)
	tblName := models.NewTask(category, task, priority, TblName)

	viewFp := "views/" + tblName + ".html"
	t, err := template.ParseFiles(viewFp, "partials/head.html", "partials/foot.html", "partials/footer.html")

	checkErr(err)

	// Remember, for partials each html file must be named properly.
	t.Execute(w, tblName)

}

// EditOneTask ...
func EditOneTask(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Get the id of the record to be edited.
	editedRecordId := ps.ByName("id")

	err := r.ParseForm()
	checkErr(err)

	taskIn := r.FormValue("taskBody")
	categoryIn := r.FormValue("taskCategory")
	priorityIn := r.FormValue("taskPriority")

	p := bluemonday.UGCPolicy()
	category := p.Sanitize(categoryIn)
	task := p.Sanitize(taskIn)
	priority := p.Sanitize(priorityIn)

	// Actually edit the record
	models.EditOneTask(category, task, priority, TblName, editedRecordId)

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}

}
