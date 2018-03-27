package controller

import (
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
)

//SetCurrentWallet takes a wallet id, and marks it as the current wallet to use
func SetCurrentWallet() model.MessageOK {
	return model.Ok(true)
}

//CreateWallet takes a wallet id, and marks it as the current wallet to use
func CreateWallet() model.MessageOK {
	return model.Ok(true)
}
