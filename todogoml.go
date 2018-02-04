package main

import (
	"net/http"
	"os"

	"github.com/mwlee9/todogoml/handlers"
	"github.com/mwlee9/todogoml/models"

	"github.com/julienschmidt/httprouter"

	_ "github.com/lib/pq"
)

// Server
func main() {

	models.CreateTable()
	// Set a default value in the event the server disconnects and home hasn't been visited yet.
	handlers.TblName = "dash"

	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handlers
	router := httprouter.New()
	router.GET("/", handlers.Dash)
	router.GET("/all", handlers.GetAllTasks)
	router.GET("/item/:id/", handlers.GetOneTask) //REMEMBER: a req without an ID still routes to this URL, passing in "favicon.ico" as the param. can't be /:id, must be diff branch url/item/:id
	//This GetOneTask func is needed in order to properly select a rec to delete!
	router.DELETE("/item/:id", handlers.DeleteOneTask)
	router.POST("/", handlers.NewTask)
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	// REMEMBER: Chrome caches the stylesheet when it can't be found...use firefox instead, or delete history in chrome.

	// Web Pages
	router.GET("/work", handlers.Work)
	router.GET("/weekend", handlers.Weekend)
	// router.GET("/groceries", handlers.Weekend)
	// router.GET("/resolutions", handlers.Weekend)
	// router.GET("/hobby", handlers.Weekend)
	// router.GET("/design", handlers.Weekend)

	http.ListenAndServe(getPort(), router)
}

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8000"
}
