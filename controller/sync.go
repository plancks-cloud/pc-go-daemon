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
}

//PushAllWallets pushes all wallets to cloud
func PushAllWallets() {
	wallets := GetWallet()
	var body = model.WalletSyncable{"Wallet", "_id", nil, wallets}
	util.Post(getDbSyncURL(), body.ToJson())

}
