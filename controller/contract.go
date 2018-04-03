package controller

import (
	"fmt"
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
	CallbackContractAsync(*contract, true)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error saving contract: %s", err))
		return model.OkMessage(false, err.Error())
	}
	log.Infoln(fmt.Sprintf("❤️ ❤️ ❤️   Contract created ID: %s", contract.ID))
	return model.Ok(true)



}

//GetContract returns all contracts stored in the datastore
func GetContract() []model.Contract {
	var contracts []model.Contract
	mongo.GetCollection(model.Contract{}).Find(nil).All(&contracts)
	return contracts
}

//GetContractResult returns all contracts results stored in the datastore
func GetContractResult() (results []model.ContractResult) {
	contracts := GetContract()
	for _, element := range contracts {
		item := model.ContractResult{Contract: element}
		item.Bids = GetBidsByContractID(element.ID)
		item.Wins = GetWinsByContractID(element.ID)
		results = append(results, item)
	}
	return results
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

//CallbackContractAsync checks an incoming DB row to see if it is interesting
func CallbackContractAsync(contract model.Contract, interesting bool) {
	go callbackContract(contract, interesting)
}

//callbackContract checks an incoming DB row to see if it is interesting
// This method is long running and should be callled asynchronously!
func callbackContract(contract model.Contract, interesting bool) {

	//Check if died of old age
	if contract.RunUntil != 0 && util.MakeTimestamp() > contract.RunUntil {
		if interesting {
			log.Infoln(fmt.Sprintf("Thinking: 🙈 🙉 🙊 Contract is ancient. Ignoring, ID: %s", contract.ID))
		}
		return
	}

	//Sleep for 10 seconds in-case I have bid in past life
	time.Sleep(10 * time.Second)

	bids := GetBidsByContractID(contract.ID)
	for _, b := range bids {
		if b.FromAccount == model.SystemWallet.ID {
			log.Infoln(fmt.Sprintf("Thinking: 🍻 🍻 🍻  I've already voted. Not voting for contract again, ID: %s", contract.ID))
			return //Already voted..
		}
	}

	//For now sleep for 10.. this should allow wins to come through
	time.Sleep(10 * time.Second)

	wins := GetWinsByContractID(contract.ID)
	if len(wins) > 0 {
		log.Infoln(fmt.Sprintf("Thinking: 😒 😒 😒  Contract has been won. Ignoring, ID: %s", contract.ID))
		return
	}

	log.Infoln(fmt.Sprintf("Thinking: ☺️ ☺️ ☺️  I'd like to consider bidding on this contract, ID: %s", contract.ID))
	considerContract(contract)

}

//considerContract checks an incoming DB row to see if I can run it and vote for it
func considerContract(contract model.Contract) {
	log.Infoln(fmt.Sprintf("Asking: ❓ ❓ ❓  Can I run this contract: %s ", contract.ID))

	//Check if I can run this spec
	canHandle := true //TODO

	if canHandle {
		log.Infoln(fmt.Sprintf("> 🤩 🤩 🤩  Actually bidding on this contract: %s ", contract.ID))
		CreateBidFromContract(contract)
		//Check for wins in a minute
		go func() {
			CheckForWinsLater(contract)
		}()
	}

}
