package model

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"time"
)

var (
	//SystemWallet is the wallet to use as the system
	SystemWallet *Wallet

	//DoorBellCommunity can be used to start the Community checks in controller/community/schedule.go
	DoorBellCommunity chan bool
	//DoorBellHealth can be used to start the Health checks in controller/health/schedule.go
	DoorBellHealth chan bool
	//DoorBellRemote can be used to start the Remote database sync in controller/remote/schedule.go
	DoorBellRemote chan bool

	//StartupTime records when the system came up
	StartupTime = time.Now()
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

//SecondsBeforeWinDeclarer is how many seconds before this machine will think about who won
const SecondsBeforeWinDeclarer = ScheduledInterval * 2

//MaxBidMultiplier helps prevent too many votes
const MaxBidMultiplier = 1.5

//MaxBidConstant also helps prevent too many votes
const MaxBidConstant = 10

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
