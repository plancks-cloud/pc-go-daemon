package api

import (
	"encoding/json"
	"fmt"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller/db"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//CreateCancelContract makes a contract void
func CreateCancelContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var cancel model.CancelContract
	err := decoder.Decode(&cancel)
	if err != nil {
		log.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", err))
		json.NewEncoder(w).Encode(model.OkMessage(false, err.Error()))
	}
	json.NewEncoder(w).Encode(db.CreateCancelContract(&cancel))

	go func() {
		//Ensure we kick off async processes
		model.DoorBellRemote <- true
	}()

}
