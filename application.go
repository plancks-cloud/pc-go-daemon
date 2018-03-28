package main

import (
	"fmt"
	"net/http"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/repo"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/api"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const port = 8080

func main() {

	initAll()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/createContract", api.CorsHandler(api.CreateContract)).Methods("POST", "OPTIONS")
	router.HandleFunc("/updateContract", api.CorsHandler(api.UpdateContract)).Methods("POST", "OPTIONS")
	router.HandleFunc("/createWallet", api.CorsHandler(api.CreateWallet)).Methods("POST", "OPTIONS")
	router.HandleFunc("/dockerListServices", api.CorsHandler(api.DockerListServices)).Methods("GET", "OPTIONS")               //ADD JSON RETURN
	router.HandleFunc("/dockerListRunningServices", api.CorsHandler(api.DockerListRunningServices)).Methods("GET", "OPTIONS") //ADD JSON RETURN
	router.HandleFunc("/forceSync", api.CorsHandler(api.ForceSync)).Methods("GET", "OPTIONS")
	router.HandleFunc("/getContract", api.CorsHandler(api.GetContract)).Methods("GET", "OPTIONS")
	router.HandleFunc("/ping", api.CorsHandler(api.Ping)).Methods("GET", "OPTIONS")
	router.HandleFunc("/setCurrentWallet", api.CorsHandler(api.SetCurrentWallet)).Methods("POST", "OPTIONS")
	router.HandleFunc("/getOneContract", api.CorsHandler(api.GetOneContract)).Methods("GET", "OPTIONS")

	log.Info(fmt.Sprintf("READY: Listening [:%v]", port))
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), router))
}

func initAll() {
	if model.GetEnvLogFormat() == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}

	log.Info("Starting")
	mongo.Init()
	repo.Init()

	controller.PullAll()
}
