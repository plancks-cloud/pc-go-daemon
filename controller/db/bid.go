package db

import (
	"math/rand"

	"fmt"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/util"
	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
)

//GetBid returns all contracts stored in the datastore
func GetBid() (bids []model.Bid) {
	mongo.GetCollection(model.Bid{}).Find(nil).All(&bids)
	return
}

//GetBidsByContractID returns all bids for a contract
func GetBidsByContractID(contractID string) (bids []model.Bid) {
	mongo.GetCollection(model.Bid{}).Find(bson.M{"contractId": contractID}).All(&bids)
	return
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
	bid := model.Bid{}
	_, err := mongo.GetCollection(&bid).RemoveAll(bson.M{"contractId": id})
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
