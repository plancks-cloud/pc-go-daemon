package main

import (
	"fmt"
	"log"
	"net/http"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/api"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
	"github.com/gorilla/mux"
)

const port = 8085

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/createContract", api.CorsHandler(api.CreateContract)).Methods("POST", "OPTIONS")
	router.HandleFunc("/createWallet", api.CorsHandler(api.CreateWallet)).Methods("POST", "OPTIONS")
	router.HandleFunc("/dockerListServices", api.CorsHandler(api.DockerListServices)).Methods("GET", "OPTIONS") //ADD JSON RETURN
	router.HandleFunc("/forceSync", api.CorsHandler(api.ForceSync)).Methods("GET", "OPTIONS")
	router.HandleFunc("/getContract", api.CorsHandler(api.GetContract)).Methods("GET", "OPTIONS")
	router.HandleFunc("/ping", api.CorsHandler(api.Ping)).Methods("GET", "OPTIONS")
	router.HandleFunc("/setCurrentWallet", api.CorsHandler(api.SetCurrentWallet)).Methods("POST", "OPTIONS")

	mongoDB()

	fmt.Printf("Listening [:%v]", port)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), router))
}

func mongoDB() {
	fmt.Println("Doing it!")
	mongo.Init()
	mongo.Push(model.OkMessage(true))
}
