package {{projectName}}

import (
	"github.com/gorilla/mux"

	"github.com/polygens/{{projectName}}/config"
)

type App struct {
	router *mux.Router
	cfg    *config.Config
}

var app *App

// Init creates and starts the app
func Init(router *mux.Router, cfg *config.Config) {
	app = &App{router, cfg, dht}

	app.setupRoutes()
}
