package model

import (
	"encoding/json"
	"fmt"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
	uuid "github.com/nu7hatch/gouuid"
	log "github.com/sirupsen/logrus"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/util"
)

//Win object represents bid document in DB
type Win struct {
	ID            string `json:"_id,omitempty" bson:"_id,omitempty"`
	ContractID    string `json:"contractId" bson:"contractId,omitempty"`
	WinnerAccount string `json:"winnerAccount" bson:"winnerAccount,omitempty"`
	Timestamp     int64  `json:"timestamp" bson:"timestamp,omitempty"`
	Signature     string `json:"signature" bson:"signature,omitempty"`
}

//WinSyncable is a wrapper for what get posted to the cloud
type WinSyncable struct {
	Collection string   `json:"collection" bson:"collection"`
	Index      string   `json:"index" bson:"index"`
	Indexes    []string `json:"indexes" bson:"indexes"`
	Rows       []Win    `json:"rows" bson:"rows"`
}

//ToJSON converts object o json
func (winSyncable WinSyncable) ToJSON() []byte {
	jsonBytes, jsonError := json.Marshal(winSyncable)
	if jsonError != nil {
		log.Fatalln(fmt.Sprintf("Error converting winSyncable to json: %s", jsonError))
	}
	return jsonBytes

}

//Push persists the win to the database
func (win Win) Push() error {
	if len(win.ID) == 0 {
		u, _ := uuid.NewV4()
		win.ID = u.String()
	}
	err := mongo.Push(win)
	return err
}

//Upsert updates a document if it exists, otherwise inserts
func (win Win) Upsert() error {
	return mongo.UpsertWithID(win.ID, win)
}

//Expired checks if a service should still be running according to the contract it was created with
func (win *Win) Expired(contract *Contract) bool {
	now := util.MakeTimestamp()
	if contract.RunUntil == 0 {
		return false
	}
	return now > contract.RunUntil
}
