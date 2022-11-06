package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"musicStreamingPlatform/config"
	db2 "musicStreamingPlatform/db"
	"musicStreamingPlatform/handler"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(config *config.Config) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DB.Host, config.DB.Port, config.DB.User, config.DB.Password, config.DB.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	a.DB = db
	a.Router = mux.NewRouter()
	a.setRouters()
	db2.SetUpTables(db)
}

func (a *App) setRouters() {
	a.Post("/release/track", a.handleRequest(handler.ReleaseNewTrack))
	a.Post("/release/album", a.handleRequest(handler.ReleaseNewAlbum))
	a.Get("/getQueriedSongs/{songTitle}", a.handleRequest(handler.GetQueriedSongs))
	a.Delete("/favourites/deleteFavouriteTrack/{songTitle}", a.handleRequest(handler.DeleteFromFavouriteList))
	a.Post("/favourites/addNewFavouriteTrack", a.handleRequest(handler.AddNewFavouriteInstances))
	a.Get("/favourites/AllFavouriteTracks/{username}", a.handleRequest(handler.GetAllFavourite))
	//a.Post("/addUser", a.handleRequest(handler.CreateUser))
	/*
		a.Post("/projects", a.handleRequest(handler.CreateProject))
		a.Get("/projects/{title}", a.handleRequest(handler.GetProject))
		a.Put("/projects/{title}", a.handleRequest(handler.UpdateProject))
		a.Delete("/projects/{title}", a.handleRequest(handler.DeleteProject))
		a.Put("/projects/{title}/archive", a.handleRequest(handler.ArchiveProject))
		a.Delete("/projects/{title}/archive", a.handleRequest(handler.RestoreProject))

		a.Get("/projects/{title}/tasks", a.handleRequest(handler.GetAllTasks))
		a.Post("/projects/{title}/tasks", a.handleRequest(handler.CreateTask))
		a.Get("/projects/{title}/tasks/{id:[0-9]+}", a.handleRequest(handler.GetTask))
		a.Put("/projects/{title}/tasks/{id:[0-9]+}", a.handleRequest(handler.UpdateTask))
		a.Delete("/projects/{title}/tasks/{id:[0-9]+}", a.handleRequest(handler.DeleteTask))
		a.Put("/projects/{title}/tasks/{id:[0-9]+}/complete", a.handleRequest(handler.CompleteTask))

	*/
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE").Queries("songTitle", "{songTitle}")
}

type RequestHandlerFunction func(Db *sql.DB, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}
