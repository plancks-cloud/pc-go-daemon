package controller

import (
	"fmt"
	"math/rand"
	"time"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/util"
	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
)

//CreateContract creates a new contract
func CreateContract(contract *model.Contract) model.MessageOK {
	err := contract.Push()
	if err != nil {
		log.Errorln(fmt.Sprintf("Error saving contract: %s", err))
		return model.OkMessage(false, err.Error())
	}
	return model.Ok(true)
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
	err := mongo.GetCollection(&contract).Find(bson.M{"_id": id}).One(&contract)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error getting contract: %s", err))
	}
	return contract, err
}

//UpdateContract upserts the given contract
func UpdateContract(contract *model.Contract) error {
	err := contract.Upsert()
	return err
}

//callbackContract checks an incoming DB row to see if it is interesting
// This method is long running and should be callled asynchronously!
func callbackContract(contract model.Contract) {

	//Check if died of old age
	if contract.RunUntil != 0 && util.MakeTimestamp() > contract.RunUntil {
		log.Infoln(fmt.Sprintf("Contract has died of old age: %s", contract.ID))
		return
	}

	//Sleep for 5 seconds incase I have bid in past life
	time.Sleep(5 * time.Second)

	bids := GetBidsByContractID(contract.ID)
	found := false
	for _, b := range bids {
		if b.FromAccount == model.SystemWallet.ID {
			found = true
			break
		}
	}

	if found {
		log.Infoln(fmt.Sprintf("I have voted on this contract: %s", contract.ID))
		return
	}

	//Sleep for 5 seconds incase I have bid in past life
	time.Sleep(5 * time.Second)

	//For now sleep for 10.. this should allow wins to come through
	time.Sleep(10 * time.Second)

	wins := GetWinsByContractID(contract.ID)
	if len(wins) > 0 {
		log.Infoln(fmt.Sprintf("Someone already one - move on: %s", contract.ID))
		return
	}

}

//considerContract checks an incoming DB row to see if I can run it and vote for it
func considerContract(contract model.Contract) {

	//Need to get current wallet

	//Check if I can run this spec
	//TODO
	canHandle := true

	if canHandle {
		bid := model.Bid{}
		bid.ContractID = contract.ID
		bid.FromAccount = model.SystemWallet.ID
		bid.Votes = rand.Intn(100000)
		bid.Timestamp = util.MakeTimestamp()
		bid.Signature = model.SystemWallet.GetSignature()
		bid.Push()
	}

}
