package db

import (
	"fmt"
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
	log.Infoln(fmt.Sprintf("❤️  Contract created ID: %s", contract.ID))
	return model.Ok(true)

}

//GetContract returns all contracts stored in the datastore
func GetContract() (contracts []model.Contract) {
	mongo.GetCollection(model.Contract{}).Find(nil).All(&contracts)
	return
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
	return
}

//GetOneContract returns a single contract
func GetOneContract(id string) (contract model.Contract, err error) {
	err = mongo.GetCollection(&contract).Find(bson.M{"_id": id}).One(&contract)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error getting contract: %s", err))
	}
	return
}

//ExpiredContract checks if a contract has expired
func ExpiredContract(contract *model.Contract) bool {
	now := util.MakeTimestamp()
	if contract.RunUntil == 0 {
		return false
	}
	return now > contract.RunUntil
}

//ExpiredContractBy checks if a contract has expired
func ExpiredContractBy(contract *model.Contract, seconds int) bool {
	now := util.MakeTimestamp()
	if contract.RunUntil == 0 {
		return false
	}

	return now > contract.RunUntil+int64(seconds*1000)
}

func DeleteContract(contract *model.Contract) {
	err := mongo.GetCollection(&contract).Remove(bson.M{"_id": contract.ID})
	if err != nil {
		log.Errorln(fmt.Sprintf("Error deleting contract: %s", err))
	}

}

//ContractExists checks if there is a contract by that ID in the db
func ContractExists(id string) bool {
	var contract model.Contract
	count, err := mongo.GetCollection(&contract).Find(bson.M{"_id": id}).Count()
	if err != nil {
		log.Errorln(fmt.Sprintf("Error getting contract: %s", err))
	}
	return count == 1
}

//UpdateContract upserts the given contract
func UpdateContract(contract *model.Contract) (err error) {
	err = contract.Upsert()
	return
}

//considerContract checks an incoming DB row to see if I can run it and vote for it
func considerContract(contract model.Contract) {
	log.Infoln(fmt.Sprintf("❓  Asking: Can I run this contract: %s ", contract.ID))

	//Check if I can run this spec
	canHandle := true //TODO

	if canHandle {
		log.Infoln(fmt.Sprintf("🤩  > Actually bidding on this contract: %s ", contract.ID))
		CreateBidFromContract(contract)
		//Check for wins in a minute
		go func() {
			CheckForWinsLater(contract)
		}()
	}

}
