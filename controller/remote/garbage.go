package remote

import (
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/util"
)

func remoteGC() {
	go util.Get(model.DBGCURL)
}
