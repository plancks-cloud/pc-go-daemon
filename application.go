package main

import (
	"fmt"
	"net/http"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/api"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const port = 8080

func main() {

	initAll()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/createContract", api.CorsHandler(api.CreateContract)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/updateContract", api.CorsHandler(api.UpdateContract)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/createWallet", api.CorsHandler(api.CreateWallet)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/dockerListServices", api.CorsHandler(api.DockerListServices)).Methods("GET", "OPTIONS")               //ADD JSON RETURN
	router.HandleFunc("/api/dockerListRunningServices", api.CorsHandler(api.DockerListRunningServices)).Methods("GET", "OPTIONS") //ADD JSON RETURN
	router.HandleFunc("/api/forceSync", api.CorsHandler(api.ForceSync)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/getContract", api.CorsHandler(api.GetContract)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/ping", api.CorsHandler(api.Ping)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/setCurrentWallet", api.CorsHandler(api.SetCurrentWallet)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/getOneContract", api.CorsHandler(api.GetOneContract)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/getContractResult", api.CorsHandler(api.GetContractResult)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/getServiceState", api.CorsHandler(api.GetServiceState)).Methods("GET", "OPTIONS")

	log.Info(fmt.Sprintf("READY: Listening [:%v]", port))
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), router))
}

func initAll() {
	if model.GetEnvLogFormat() == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}

	log.Info("Starting")
	mongo.Init()
	model.InitRepo()

	// go func() {
	// 	//Wake up the function
	// 	util.Options(model.DBSyncURL)
	// 	for {
	// 		//Sync and sleep
	// 		log.Infoln(fmt.Sprintf("> Time to sync"))
	// 		controller.PullAll()
	// 		controller.PushAll()
	// 		time.Sleep(1 * time.Minute)
	// 	}
	// }()

	controller.SyncDatabase()
	controller.ReconServices()
}
