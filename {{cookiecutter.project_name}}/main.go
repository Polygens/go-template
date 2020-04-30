package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/polygens/{{cookiecutter.project_name}}/service"
	"github.com/polygens/{{cookiecutter.project_name}}/service/config"
)

var version string
var svc *service.Service

func main() {
	log.Infof("Starting %s version: %s", filepath.Base(os.Args[0]), version)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}

	logLvl, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatalf("Failed to set log level: %s", err)
	}

	log.SetLevel(logLvl)

	r := mux.NewRouter()

	svc = service.Setup(r, cfg)
	
	closeHandler()

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(int(cfg.HTTPPort)), r))
}

func closeHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Info("Service interrupted, exiting gracefully")
		svc.Close()
		os.Exit(0)
	}()
}
