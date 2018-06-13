package remote

import (
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/util/http"
)

//Init warms up the remote functions
func Init() {
	go http.Options(model.DBSyncURL)
	go http.Options(model.DBGCURL)

}
