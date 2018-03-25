package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
)

//CreateContract creates a new contract
func CreateContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var contract model.Contract
	err := decoder.Decode(&contract)
	if err != nil {
		fmt.Println(fmt.Sprintf("There was a problem decoding the post message: %s", err))
		json.NewEncoder(w).Encode(model.OkMessage(false))
	}
	json.NewEncoder(w).Encode(controller.CreateContract(&contract))
}

//GetContract returns a contract by the ID given
func GetContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(controller.GetContract())
}

//GetOneContract returns a contract by the ID given
func GetOneContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	item := r.URL.Query().Get("id")
	contract, err := controller.GetOneContract(item)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error getting contract (%s): %s", item, err))
	}
	json.NewEncoder(w).Encode(contract)
}
