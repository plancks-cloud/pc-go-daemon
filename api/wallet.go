package api

import (
	"encoding/json"
	"net/http"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller"
)

//CreateWallet creates a new wallet
func CreateWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(controller.CreateWallet())
}

//SetCurrentWallet sets the currently used wallet
func SetCurrentWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(controller.SetCurrentWallet())
}
