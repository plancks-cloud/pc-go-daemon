package model

import (
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
)

//Win object represents bid document in DB
type Win struct {
	ID            string `json:"_id,omitempty" bson:"_id,omitempty"`
	ContractID    string `json:"contractId"`
	WinnerAccount string
	timestamp     int64
	signature     string
}

//Push persists the win to the database
func (item Win) Push() error {
	err := mongo.Push(item)
	return err
}

//Upsert updates a document if it exists, otherwise inserts
func (item Win) Upsert() error {
	return mongo.UpsertWithID(item.ID, item)
}
