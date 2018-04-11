package model

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

var (
	//SystemWallet is the wallet to use as the system
	SystemWallet *Wallet

	DoorBellCommunity chan bool
	DoorBellHealth    chan bool
	DoorBellRemote    chan bool
)

//DBSyncURL is the endpoint for the function
const DBSyncURL = "https://us-central1-plancks-cloud.cloudfunctions.net/pc-function-db-sync-v1"

//DBGCURL is the endpoint for the function that performs GC remotely
const DBGCURL = "https://us-central1-plancks-cloud.cloudfunctions.net/pc-function-db-gc-v1"

//ScheduledInterval is how often some things run
const ScheduledInterval = 25

//AncientAgeSeconds is how long before a row can be GCd
const AncientAgeSeconds = 300

//WinnerAgeSeconds is how many seconds before a winner can be declared
const WinnerAgeSeconds = 70

//InitRepo initialises the repository and it's variables.
func InitRepo() {
	initWallet()
}

func initWallet() {
	log.Debugln(fmt.Sprintf("Wallet: %s", GetEnvWallet()))

	//Instantiate wallet if not there - just so that there is an object in the DB
	walletName := GetEnvWallet()
	if len(walletName) == 0 {
		log.Fatalln("Could not find environment variable for wallet: WALLET - restart with your wallet")
	}
	wallet := Wallet{ID: walletName, PrivateKey: walletName, PublicKey: walletName, Signature: walletName}
	wallet.Upsert()

	SystemWallet = &wallet
}
