package db

import (
	"fmt"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mem"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/util"
	"github.com/hashicorp/go-memdb"
	log "github.com/sirupsen/logrus"
)

const contractTable = "Contract"

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

//GetContract returns all contracts stored in the database
func GetContract() (contracts []model.Contract) {
	res, err := mem.GetAll(contractTable)
	return iteratorToManyContracts(res, err)
}

//GetContractResult returns all contracts results stored in the database
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
	item, err := mem.GetUniqueById(contractTable, id)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error Getting one contract: %s", err))
		return model.Contract{}, err
	}
	return item.(model.Contract), nil
}

//ExpiredContract checks if a contract has expired
func ExpiredContract(contract *model.Contract) bool {
	return ExpiredContractBy(contract, 0)
}

//ExpiredContractBy checks if a contract has expired
func ExpiredContractBy(contract *model.Contract, seconds int) bool {
	if contract.SecondsToLive == 0 {
		return false
	}

	now := util.MakeTimestamp()
	expires := contract.Timestamp + (1000 * (contract.SecondsToLive + 60 + int64(seconds)))

	return now > expires
}

//DeleteContract deletes a contract from the database
func DeleteContract(contract *model.Contract) {
	_, err := mem.Delete(contractTable, "id", contract.ID)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error deleting contract: %s", err))
	}

}

//ContractExists checks if there is a contract by that ID in the db
func ContractExists(id string) bool {
	res, err := mem.GetUniqueById(contractTable, id)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error getting services: %s", err))
		//TODO: discuss this... super hard
		return false
	}
	return res != nil

}

//UpdateContract upserts the given contract
func UpdateContract(contract *model.Contract) (err error) {
	err = contract.Upsert()
	return
}

func iteratorToManyContracts(iterator memdb.ResultIterator, err error) (items []model.Contract) {
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	if iterator == nil {
		return items
	}
	more := true
	for more {
		next := iterator.Next()
		if next == nil {
			more = false
			continue
		}
		item := next.(model.Contract)
		items = append(items, item)
	}
	log.Infoln(fmt.Sprintf("Contract iterator counts: %d", len(items)))
	return items

}
