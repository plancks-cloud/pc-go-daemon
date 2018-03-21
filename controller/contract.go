package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
)

//CreateContract creates a new contract
func CreateContract(r *http.Request) model.MessageOK {
	decoder := json.NewDecoder(r.Body)
	var contract model.Contract
	err := decoder.Decode(&contract)
	if err != nil {
		fmt.Println(fmt.Sprintf("There was a massive issue decoding the post message: %s", err))
		return model.OkMessage(false)
	}
	err = contract.Push()
	if err != nil {
		fmt.Println(fmt.Sprintf("Error saving contract: %s", err))
		return model.OkMessage(false)
	}
	return model.OkMessage(true)
}

//GetContract returns all contracts stored in the datastore
func GetContract() []model.Contract {
	var contracts []model.Contract
	contracts = append(contracts, model.Contract{})
	return contracts
}
