package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

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
		log.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", err))
		json.NewEncoder(w).Encode(model.OkMessage(false, err.Error()))
	}
	json.NewEncoder(w).Encode(controller.CreateContract(&contract))
}

//GetContract returns a contract by the ID given
func GetContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(controller.GetContract())
}

//GetContractResult returns a contract by the ID given
func GetContractResult(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(controller.GetContractResult())
}

//GetOneContract returns a contract by the ID given
func GetOneContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	item := r.URL.Query().Get("id")
	contract, err := controller.GetOneContract(item)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error getting contract (%s): %s", item, err))
	}
	json.NewEncoder(w).Encode(contract)
}

//UpdateContract upserts a new contract
func UpdateContract(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var contract model.Contract
	err := decoder.Decode(&contract)
	if err != nil {
		log.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", err))
		encoder.Encode(model.OkMessage(false, err.Error()))
		return
	}
	err = controller.UpdateContract(&contract)
	if err != nil {
		log.Errorln(fmt.Sprintf("There was a problem updating the document: %s", err))
		encoder.Encode(model.OkMessage(false, err.Error()))
		return
	}
	json.NewEncoder(w).Encode(encoder.Encode(model.Ok(true)))
}
