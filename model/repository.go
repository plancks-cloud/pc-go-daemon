package model

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

var (
	//SystemWallet is the wallet to use as the system
	SystemWallet *Wallet
)

//DBSyncURL is the endpoint for the function
const DBSyncURL = "https://us-central1-plancks-cloud.cloudfunctions.net/pc-function-db-sync-v1"

//InitRepo initialises the repository and it's variables.
func InitRepo() {
	initWallet()
}

func initWallet() {
	log.Debugln(fmt.Sprintf("Wallet: %s", GetEnvWallet()))

	//Instantiate wallet if not there - just so that there is an object in the DB
	walletName := GetEnvWallet()
	if len(walletName) == 0 {
		log.Fatalln("Could not find environment varable for wallet: WALLET - restart with your wallet")
	}
	wallet := Wallet{ID: walletName, PrivateKey: walletName, PublicKey: walletName, Signature: walletName}
	wallet.Upsert()

	SystemWallet = &wallet
}
