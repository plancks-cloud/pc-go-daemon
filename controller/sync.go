package controller

import (
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/util"
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
	var body = model.WalletSyncable{"Wallet", "_id", nil, wallets}
	util.Post(getDbSyncURL(), body.ToJSON())

}

//PushAllContracts pushes all contracts to cloud
func PushAllContracts() {
	contracts := GetContract()
	var body = model.ContractSyncable{"Contract", "_id", nil, contracts}
	util.Post(getDbSyncURL(), body.ToJSON())

}

//PushAllBids pushes all bids to cloud
func PushAllBids() {
	bids := GetBid()
	var body = model.BidSyncable{"Bid", "_id", nil, bids}
	util.Post(getDbSyncURL(), body.ToJSON())

}

//PushAllWins pushes all wins to cloud
func PushAllWins() {
	wins := GetWin()
	var body = model.WinSyncable{"Win", "_id", nil, wins}
	util.Post(getDbSyncURL(), body.ToJSON())

}
