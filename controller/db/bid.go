package db

import (
	"math/rand"

	"fmt"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mem"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/util"
	"github.com/hashicorp/go-memdb"
	log "github.com/sirupsen/logrus"
)

const bidTable = "Bid"

//GetBid returns all contracts stored in the datastore
func GetBid() (bids []model.Bid) {
	res, err := mem.GetAll(bidTable)
	return iteratorToManyBids(res, err)
}

//GetBidsByContractID returns all bids for a contract
func GetBidsByContractID(contractID string) (bids []model.Bid) {
	res, err := mem.GetAllByFieldAndValue(bidTable, contractId, contractID)
	return iteratorToManyBids(res, err)
}

//CreateBidFromContract inserts a new bid for a contract
func CreateBidFromContract(contract model.Contract) {
	log.Infoln(fmt.Sprintf("✔️  Actually bidding on: %s", contract.ID))
	bid := model.Bid{}
	bid.ContractID = contract.ID
	bid.FromAccount = model.SystemWallet.ID
	bid.Votes = rand.Intn(100000)
	bid.Timestamp = util.MakeTimestamp()
	bid.Signature = model.SystemWallet.GetSignature()
	bid.Push()

}

//DeleteBidsByContractID deletes a contract with an ID
func DeleteBidsByContractID(id string) {
	_, err := mem.Delete(bidTable, contractId, id)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error deleting bids by contractId: %s", err))
	}

}

//HaveIBidOnContract checks if a wallet has bid
func HaveIBidOnContract(id string) bool {
	bids := GetBidsByContractID(id)
	myID := model.SystemWallet.ID
	for _, bid := range bids {
		if bid.FromAccount == myID {
			return true
		}
	}
	return false
}

func iteratorToManyBids(iterator memdb.ResultIterator, err error) (items []model.Bid) {
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
		item := next.(model.Bid)
		items = append(items, item)
	}
	log.Infoln(fmt.Sprintf("Bid iterator counts: %d", len(items)))
	return items

}
