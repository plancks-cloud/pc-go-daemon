package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
)

//PushAll gets all rows in DB and pushes to DB
func PushAll() {
	PushAllWallets()
}

//PushAllWallets pushes all wallets to cloud
func PushAllWallets() {

	var body = model.WalletSyncable{}

	wallets := GetWallet()
	body.Collection = "Wallet"
	body.Index = "_id"
	body.Rows = wallets

	url := "https://us-central1-plancks-cloud.cloudfunctions.net/pc-function-db-sync-v1"

	jsonString, jsonError := json.Marshal(body)
	if jsonError != nil {
		log.Errorln(fmt.Sprintf("Error converting wallets to json: %s", jsonError))
		panic(jsonError)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonString))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error doing http post for wallet: %s", err))
		panic(err)
	}
	defer resp.Body.Close()

	log.Infoln(fmt.Sprintf("Response status: %s", resp.Status))
	r, _ := ioutil.ReadAll(resp.Body)
	log.Infoln(fmt.Sprintf("Response body: %s", string(r)))

}
