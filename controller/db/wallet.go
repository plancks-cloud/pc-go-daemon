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
func GetWallet() chan model.Wallet {
	res, err := mem.GetAll(walletTable)
	return iteratorToManyWallets(res, err)
}

func iteratorToManyWallets(iterator memdb.ResultIterator, err error) (res chan model.Wallet) {
	c := mem.IteratorToChannel(iterator, err)
	go func() {
		for i := range c {
			item := i.(model.Wallet)
			res <- item
		}
		close(c)

	}()
	return res

}
