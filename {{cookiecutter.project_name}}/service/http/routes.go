package http

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/polygens/{{cookiecutter.project_name}}/service/http/middleware"
)

type HTTP struct {
	r *mux.Router
}

func Setup(r *mux.Router) *HTTP {
	http := &HTTP{r}

	http.setupRoutes()

	return http
}

func (http *HTTP) setupRoutes() {
	http.r.Use(middleware.Metrics)
	http.r.Handle("/metrics", promhttp.Handler()).Methods("GET")
	http.r.HandleFunc("/ping", handler.Health).Methods("GET")
	http.r.HandleFunc("/ready", handler.Health).Methods("GET")
	http.r.HandleFunc("/live", handler.Health).Methods("GET")
}
