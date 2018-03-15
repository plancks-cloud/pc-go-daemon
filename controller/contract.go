package controller

import (
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
)

//CreateContract creates a new contract
func CreateContract() model.MessageOK {
	return model.OkMessage(true)
}

//GetContract returns all contracts stored in the datastore
func GetContract() []model.Contract {
	var contracts []model.Contract
	contracts = append(contracts, model.Contract{})
	return contracts
}
