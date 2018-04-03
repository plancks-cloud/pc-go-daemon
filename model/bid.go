package model

import (
	"encoding/json"
	"fmt"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
	uuid "github.com/nu7hatch/gouuid"
	log "github.com/sirupsen/logrus"
)

//Bid object represents bid document in DB
type Bid struct {
	ID          string `json:"_id,omitempty" bson:"_id,omitempty"`
	ContractID  string `json:"contractId" bson:"contractId,omitempty"`
	Votes       int    `json:"votes" bson:"votes,omitempty"`
	FromAccount string `json:"fromAccount" bson:"fromAccount,omitempty"`
	Timestamp   int64  `json:"timestamp" bson:"timestamp,omitempty"`
	Signature   string `json:"signature" bson:"signature,omitempty"`
}

//BidSyncable is a wrapper for what get posted to the cloud
type BidSyncable struct {
	Collection string   `json:"collection" bson:"collection"`
	Index      string   `json:"index" bson:"index"`
	Indexes    []string `json:"indexes" bson:"indexes"`
	Rows       []Bid    `json:"rows" bson:"rows"`
}

//ToJSON converts object o json
func (bidSyncable BidSyncable) ToJSON() []byte {
	jsonBytes, jsonError := json.Marshal(bidSyncable)
	if jsonError != nil {
		log.Fatalln(fmt.Sprintf("Error converting bidSyncable to json: %s", jsonError))
	}
	return jsonBytes

}

//Push saves a bid to MongoDB
func (bid Bid) Push() error {
	if len(bid.ID) == 0 {
		u, _ := uuid.NewV4()
		bid.ID = u.String()
	}
	err := mongo.Push(bid)
	return err
}

//DbID returns the ID of the bid
// func (bid Bid) DbID() bson.ObjectId {
func (bid Bid) DbID() string {
	return bid.ID
}

//Upsert ..
func (bid Bid) Upsert() error {
	err := mongo.UpsertWithID(bid.ID, bid)
	return err
}
