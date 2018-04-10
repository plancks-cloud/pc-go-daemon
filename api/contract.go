package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller/db"
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
	json.NewEncoder(w).Encode(db.CreateContract(&contract))

	go func() {
		//Ensure we kick off async processes
		model.DoorBellRemote <- true
	}()

}

//GetContract returns a contract by the ID given
func GetContract(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db.GetContract())
}

//GetContractResult returns a contract by the ID given
func GetContractResult(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db.GetContractResult())
}

//GetOneContract returns a contract by the ID given
func GetOneContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	item := r.URL.Query().Get("id")
	contract, err := db.GetOneContract(item)
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
	err = db.UpdateContract(&contract)
	if err != nil {
		log.Errorln(fmt.Sprintf("There was a problem updating the document: %s", err))
		encoder.Encode(model.OkMessage(false, err.Error()))
		return
	}
	json.NewEncoder(w).Encode(encoder.Encode(model.Ok(true)))
}
