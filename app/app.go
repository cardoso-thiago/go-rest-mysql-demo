package app

import (
	"github.com/cardoso-thiago/go-rest-mysql-demo/app/db"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     db.ArtistDB
}

func New() *App {
	app := &App{
		Router: mux.NewRouter(),
	}
	app.initRoutes()
	return app
}

func (app *App) initRoutes() {
	app.Router.HandleFunc("/api/artist", app.CreateArtistHandler()).Methods("POST")
	app.Router.HandleFunc("/api/artist", app.GetArtistsHandler()).Methods("GET")
}
