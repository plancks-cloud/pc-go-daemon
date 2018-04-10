package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller/db"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/util"
	log "github.com/sirupsen/logrus"
)

//CreateWallet creates a new wallet
func CreateWallet(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var wallet model.Wallet
	err := decoder.Decode(&wallet)
	if err != nil {
		log.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", err))
		util.RespondWithJsonError(w, err)
		return
	}
	util.RespondWithJsonObject(w, db.CreateWallet(&wallet))

}

//SetCurrentWallet sets the currently used wallet
func SetCurrentWallet(w http.ResponseWriter, _ *http.Request) {
	util.RespondWithJsonObject(w, db.SetCurrentWallet())
}
