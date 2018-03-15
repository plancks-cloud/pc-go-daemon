package api

import (
	"encoding/json"
	"net/http"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller"
)

//CreateContract creates a new contract
func CreateContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(controller.CreateContract())
}

//GetContract returns a contract by the ID given
func GetContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(controller.GetContract())
}
