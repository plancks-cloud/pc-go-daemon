package controller

import (
	"fmt"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
)

//SetCurrentWallet takes a wallet id, and marks it as the current wallet to use
func SetCurrentWallet() model.MessageOK {
	return model.OkMessage(true)
}

//CreateWallet takes a wallet id, and marks it as the current wallet to use
func CreateWallet(wallet *model.Wallet) model.MessageOK {
	err := wallet.Push()
	if err != nil {
		log.Errorln(fmt.Sprintf("Error saving wallet: %s", err))
		return model.OkMessage(false)
	}
	return model.OkMessage(true)
}

//GetWallet returns all wallets stored in the datastore
func GetWallet() []model.Wallet {
	var wallets []model.Wallet
	mongo.GetCollection(model.Wallet{}).Find(nil).All(&wallets)
	for _, wallet := range wallets {
		log.Infoln(fmt.Sprintf("Wallet: %s", wallet.ID))
	}
	return wallets
}

//GetOneWallet returns a single contract
func GetOneWallet(id string) (model.Wallet, error) {
	var wallet model.Wallet
	err := mongo.GetCollection(&wallet).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&wallet)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error getting wallet: %s", err))
	}
	return wallet, err
}
