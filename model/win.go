package model

import (
	"encoding/json"
	"fmt"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/mem"
	"github.com/nu7hatch/gouuid"
	log "github.com/sirupsen/logrus"
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
func (win Win) Push() (err error) {
	if len(win.ID) == 0 {
		u, _ := uuid.NewV4()
		win.ID = u.String()
	}
	err = mem.Push(win)
	return
}

//Upsert updates a document if it exists, otherwise inserts
func (win Win) Upsert() error {
	return mem.Push(win)
}
