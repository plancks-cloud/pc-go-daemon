package remote

import (
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/util"
)

//Init warms up the remote functions
func Init() {
	go util.Options(model.DBSyncURL)
	go util.Options(model.DBGCURL)

}
