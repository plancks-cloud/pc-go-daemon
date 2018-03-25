package model

import (
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
	"github.com/globalsign/mgo/bson"
)

//Contract represents a contract issued to run a container
type Contract struct {
	//Audit & admin
	ID        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Account   string
	Signature string
	Timestamp int64

	//Specification
	Images           map[string]string
	Instances        int
	RequiredMBMemory int
	RequiredCPUCores int
	RunUntil         int64
	AllowSuicide     int64

	StartStrategy string
}

const index = "_id"

//Push saves a contract to MongoDB
func (contract Contract) Push() error {
	err := mongo.Push(contract)
	return err
}
