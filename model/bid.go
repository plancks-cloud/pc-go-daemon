package model

import (
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
	uuid "github.com/nu7hatch/gouuid"
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
