package api

import (
	"encoding/json"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/util"
	"net/http"
)

//GetPrivatePublicKey generates a cryptographically secure key pair
func GetPrivatePublicKey(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pri, pub := util.GeneratePrivatePublicKeys()
	keyPair := model.KeyPair{PrivateKey: pri, PublicKey: pub}

	json.NewEncoder(w).Encode(keyPair)
}
