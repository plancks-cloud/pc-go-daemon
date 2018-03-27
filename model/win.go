package model

import (
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
	uuid "github.com/nu7hatch/gouuid"
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
