package controller

import (
	"fmt"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
)

//CreateContract creates a new contract
func CreateContract(contract *model.Contract) model.MessageOK {
	err := contract.Push()
	if err != nil {
		log.Errorln(fmt.Sprintf("Error saving contract: %s", err))
		return model.OkMessage(false)
	}
	return model.OkMessage(true)
}

//GetContract returns all contracts stored in the datastore
func GetContract() []model.Contract {
	var contracts []model.Contract
	mongo.GetCollection(model.Contract{}).Find(nil).All(&contracts)
	for _, contract := range contracts {

		log.Infoln(fmt.Sprintf("Contract Acccount: %s - ID: %s", contract.Account, contract.ID))
	}
	return contracts
}

//GetOneContract returns a single contract
func GetOneContract(id string) (model.Contract, error) {
	var contract model.Contract
	log.Infoln(fmt.Sprintf("Horrors! %s", bson.ObjectIdHex(id)))
	err := mongo.GetByID(&contract, id)
	return contract, err
}
