package remote

import (
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/util/http"
)

func remoteGC() {
	go http.Get(model.DBGCURL)
}
