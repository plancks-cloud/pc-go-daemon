package model

import (
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mem"
	"github.com/nu7hatch/gouuid"
)

//CancelContract represents a database row
type CancelContract struct {
	ID          string `json:"_id,omitempty" bson:"_id,omitempty"`
	ContractID  string `json:"contractId" bson:"contractId,omitempty"`
	FromAccount string `json:"fromAccount" bson:"fromAccount,omitempty"`
	Timestamp   int64  `json:"timestamp" bson:"timestamp,omitempty"`
	Signature   string `json:"signature" bson:"signature,omitempty"`
}

//Push saves a bid to Mongodb
func (item CancelContract) Push() (err error) {
	if len(item.ID) == 0 {
		u, _ := uuid.NewV4()
		item.ID = u.String()
	}
	err = mem.Push(item)
	return
}
