package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/util"

	log "github.com/sirupsen/logrus"
)

func DBPullDown() {
	log.Infoln("Inside the pulldown")
	contracts := DBSyncDownContract()
	for _, contract := range contracts {
		contract.Upsert()
	}
}

func DBSyncDownContract() (contracts []model.Contract) {
	resp, err := DBSyncDownRquest(model.Contract{})
	if err != nil {
		log.Errorln(fmt.Sprintf("Error getting contacts during sync: %s", err))
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&contracts)

	if err != nil {
		log.Errorln(fmt.Sprintf("Error decoding contract during sync: %s", err))
		return
	}

	return contracts
}

//DBSyncDownRquest ...
func DBSyncDownRquest(typeOf interface{}) (*http.Response, error) {
	typeName := util.GetType(typeOf)
	urlBase := "https://us-central1-plancks-cloud.cloudfunctions.net/pc-function-db-sync-v1"
	url := fmt.Sprintf("%s?collection=%s", urlBase, typeName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error creating request during sync: %s", err))
		return nil, err
	}
	client := &http.Client{}

	return client.Do(req)
}

//DBSyncDown object is coming through as []interface, not []contract
//TODO messed up
func DBSyncDown(typeOf interface{}, object interface{}) {

}
