package repo

import (
	"fmt"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	log "github.com/sirupsen/logrus"
)

var (
	//Wallet is the wallet to use as the system
	Wallet *model.Wallet
)

//Init initialises the repository and it's variables.
func Init() {
	initWallet()
}

func initWallet() {
	log.Info(fmt.Sprintf("Wallet: %s", model.GetEnvWallet()))

	//Instantiate wallet if not there - just so that there is an object in the DB
	walletName := model.GetEnvWallet()
	if len(walletName) == 0 {
		log.Fatalln("Could not find environment varable for wallet: WALLET - restart with your wallet")
	}
	wallet := model.Wallet{ID: walletName, PrivateKey: walletName, PublicKey: walletName, Signature: walletName}
	wallet.Upsert()

	Wallet = &wallet
}
