package model

import (
	"encoding/json"
	"fmt"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/mem"
	"github.com/nu7hatch/gouuid"
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

//Push saves a bid to Mongodb
func (bid Bid) Push() (err error) {
	if len(bid.ID) == 0 {
		u, _ := uuid.NewV4()
		bid.ID = u.String()
	}
	err = mem.Push(bid)
	return
}

//DbID returns the ID of the bid
func (bid Bid) DbID() string {
	return bid.ID
}

//Upsert ..
func (bid Bid) Upsert() (err error) {
	err = mem.Push(bid)
	return
}

//ByVotes is a struct
type ByVotes []Bid

func (n ByVotes) Len() int           { return len(n) }
func (n ByVotes) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n ByVotes) Less(i, j int) bool { return n[i].Votes < n[j].Votes }
