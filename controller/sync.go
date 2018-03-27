package controller

import (
	"bytes"
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
	wallets := GetWallet()
	var body = model.WalletSyncable{"Wallet", "_id", nil, wallets}
	jsonBytes := body.ToJson()
	Post(jsonBytes)

}

//Post method does a simple post
func Post(jsonBytes []byte) {
	url := "https://us-central1-plancks-cloud.cloudfunctions.net/pc-function-db-sync-v1"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
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
