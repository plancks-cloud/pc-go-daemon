package db

import (
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
)

//ForceSync syncs databases
func ForceSync() model.MessageOK {
	return model.Ok(true)
}
