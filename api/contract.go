package api

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller"
)

//CreateContract creates a new contract
func CreateContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(controller.CreateContract())
}

//CreateWallet creates a new wallet
func CreateWallet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello creator, %q", html.EscapeString(r.URL.Path))
}

//DockerListServices lists all docker services running
func DockerListServices(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello creator, %q", html.EscapeString(r.URL.Path))
}

//ForceSync forces a sync
func ForceSync(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello creator, %q", html.EscapeString(r.URL.Path))
}

//GetContract returns a contract by the ID given
func GetContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(controller.GetContract())
}

//Ping perform a health check
func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(controller.HealthCheck())
}

//SetCurrentWallet sets the currently used wallet
func SetCurrentWallet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello creator, %q", html.EscapeString(r.URL.Path))
}

//CorsHandler does cors check
func CorsHandler(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			fmt.Fprintf(w, "Options")
		} else {
			f.ServeHTTP(w, r)
		}
	}
}
