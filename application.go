package main

import (
	"fmt"
	"log"
	"net/http"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/api"
	"github.com/gorilla/mux"
)

const port = 8080

func main() {

	initAll()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/createContract", api.CorsHandler(api.CreateContract)).Methods("POST", "OPTIONS")
	router.HandleFunc("/createWallet", api.CorsHandler(api.CreateWallet)).Methods("POST", "OPTIONS")
	router.HandleFunc("/dockerListServices", api.CorsHandler(api.DockerListServices)).Methods("GET", "OPTIONS") //ADD JSON RETURN
	router.HandleFunc("/forceSync", api.CorsHandler(api.ForceSync)).Methods("GET", "OPTIONS")
	router.HandleFunc("/getContract", api.CorsHandler(api.GetContract)).Methods("GET", "OPTIONS")
	router.HandleFunc("/ping", api.CorsHandler(api.Ping)).Methods("GET", "OPTIONS")
	router.HandleFunc("/setCurrentWallet", api.CorsHandler(api.SetCurrentWallet)).Methods("POST", "OPTIONS")
	router.HandleFunc("/getOneContract", api.CorsHandler(api.GetOneContract)).Methods("GET", "OPTIONS")

	fmt.Printf("Listening [:%v]", port)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), router))
}

func initAll() {
	mongo.Init()
}
