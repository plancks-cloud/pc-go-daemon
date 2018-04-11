package db

import (
	"fmt"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/mem"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"github.com/hashicorp/go-memdb"
	log "github.com/sirupsen/logrus"
)

const walletTable = "Wallet"

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

//GetWallet returns all wallets stored in the database
func GetWallet() (wallets []model.Wallet) {
	res, err := mem.GetAll(walletTable)
	return iteratorToManyWallets(res, err)
}

func iteratorToManyWallets(iterator memdb.ResultIterator, err error) (items []model.Wallet) {
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	more := true
	for more {
		next := iterator.Next()
		if next == nil {
			more = false
			continue
		}
		item := next.(model.Wallet)
		items = append(items, item)
	}
	return items

}
