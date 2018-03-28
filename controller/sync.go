package controller

import (
	"encoding/json"
	"fmt"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/util"
	log "github.com/sirupsen/logrus"
)

//getDbSyncURL contains the location of the PUSH endpoint
func getDbSyncURL() string {
	return "https://us-central1-plancks-cloud.cloudfunctions.net/pc-function-db-sync-v1"
}

//PushAll gets all rows in DB and pushes to DB
func PushAll() {
	PushAllWallets()
	PushAllContracts()
	PushAllBids()
	PushAllWins()
}

//PushAllWallets pushes all wallets to cloud
func PushAllWallets() {
	wallets := GetWallet()
	body := model.WalletSyncable{Collection: "Wallet", Index: "_id", Indexes: nil, Rows: wallets}
	util.Post(getDbSyncURL(), body.ToJSON())

}

//PushAllContracts pushes all contracts to cloud
func PushAllContracts() {
	contracts := GetContract()
	body := model.ContractSyncable{Collection: "Contract", Index: "_id", Indexes: nil, Rows: contracts}
	util.Post(getDbSyncURL(), body.ToJSON())

}

//PushAllBids pushes all bids to cloud
func PushAllBids() {
	bids := GetBid()
	body := model.BidSyncable{Collection: "Bid", Index: "_id", Indexes: nil, Rows: bids}
	util.Post(getDbSyncURL(), body.ToJSON())

}

//PushAllWins pushes all wins to cloud
func PushAllWins() {
	wins := GetWin()
	body := model.WinSyncable{Collection: "Win", Index: "_id", Indexes: nil, Rows: wins}
	util.Post(getDbSyncURL(), body.ToJSON())

}

//PullAll gets all rows in cloud DB and puts them in local DB
func PullAll() {
	contracts := PullAllContracts()
	for _, contract := range contracts {
		contract.Upsert()
	}
	wallets := PullAllWallets()
	for _, item := range wallets {
		item.Upsert()
	}
	bids := PullAllBids()
	for _, item := range bids {
		item.Upsert()
	}
	wins := PullAllWins()
	for _, item := range wins {
		item.Upsert()
	}
}

//PullAllContracts gets all contracts in the cloud DB
func PullAllContracts() (contracts []model.Contract) {
	typeName := util.GetType(model.Contract{})
	url := fmt.Sprintf("%s?collection=%s", getDbSyncURL(), typeName)
	resp, err := util.Get(url)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error getting contacts during sync: %s", err))
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&contracts)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error decoding contract during sync: %s", err))
		return
	}
	return contracts
}

//PullAllWallets gets all wallets in the cloud DB
func PullAllWallets() (wallets []model.Wallet) {
	typeName := util.GetType(model.Wallet{})
	url := fmt.Sprintf("%s?collection=%s", getDbSyncURL(), typeName)
	resp, err := util.Get(url)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error getting wallets during sync: %s", err))
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&wallets)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error decoding wallets during sync: %s", err))
		return
	}
	return wallets
}

//PullAllBids gets all bids in the cloud DB
func PullAllBids() (bids []model.Bid) {
	typeName := util.GetType(model.Bid{})
	url := fmt.Sprintf("%s?collection=%s", getDbSyncURL(), typeName)
	resp, err := util.Get(url)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error getting bids during sync: %s", err))
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&bids)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error decoding bids during sync: %s", err))
		return
	}
	return bids
}

//PullAllWins gets all wins in the cloud DB
func PullAllWins() (wins []model.Win) {
	typeName := util.GetType(model.Win{})
	url := fmt.Sprintf("%s?collection=%s", getDbSyncURL(), typeName)
	resp, err := util.Get(url)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error getting wins during sync: %s", err))
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&wins)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error decoding wins during sync: %s", err))
		return
	}
	return wins
}