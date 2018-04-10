package db

import (
	"fmt"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
	log "github.com/sirupsen/logrus"
)

//SetCurrentWallet takes a wallet id, and marks it as the current wallet to use
func SetCurrentWallet() model.MessageOK {
	return model.Ok(true)
}

//CreateWallet takes a wallet id, and marks it as the current wallet to use
func CreateWallet(wallet *model.Wallet) model.MessageOK {
	err := wallet.Push()
	if err != nil {
		log.Errorln(fmt.Sprintf("Error saving wallet: %s", err))
		return model.OkMessage(false, err.Error())
	}
	return model.Ok(true)
}

//GetWallet returns all wallets stored in the datastore
func GetWallet() (wallets []model.Wallet) {
	mongo.GetCollection(model.Wallet{}).Find(nil).All(&wallets)
	return
}
