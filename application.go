package main

import (
	"fmt"
	"net/http"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/api"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller/community"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller/health"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller/remote"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mem"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const port = 8080

func main() {

	initAll()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/createContract", api.CORSHandler(api.CreateContract)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/updateContract", api.CORSHandler(api.UpdateContract)).Methods("POST", "OPTIONS")
	// router.HandleFunc("/api/createWallet", api.CORSHandler(api.CreateWallet)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/dockerListServices", api.CORSHandler(api.DockerListServices)).Methods("GET", "OPTIONS")               //ADD JSON RETURN
	router.HandleFunc("/api/dockerListRunningServices", api.CORSHandler(api.DockerListRunningServices)).Methods("GET", "OPTIONS") //ADD JSON RETURN
	router.HandleFunc("/api/forceSync", api.CORSHandler(api.ForceSync)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/getContract", api.CORSHandler(api.GetContract)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/ping", api.CORSHandler(api.Ping)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/setCurrentWallet", api.CORSHandler(api.SetCurrentWallet)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/getOneContract", api.CORSHandler(api.GetOneContract)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/getContractResult", api.CORSHandler(api.GetContractResult)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/getServiceStateResult", api.CORSHandler(api.GetServiceStateResult)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/🦖", api.CORSHandler(api.CheckStatus)).Methods("GET", "OPTIONS")

	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), router))
	log.Info(fmt.Sprintf("READY: Listening [:%v]", port))
}

func initAll() {
	if model.GetEnvLogFormat() == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{})
	}

	// log.SetLevel(log.ErrorLevel)

	log.Info("Starting")
	mem.Init()
	model.InitRepo()

	remote.Init()
	remote.ScheduleRemoteSync()
	community.ScheduleCommunityActivities()
	health.ScheduleHealthCheck()

}
