package model

import (
	"fmt"
	"strings"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
)

//Contract represents a contract issued to run a container
type Contract struct {
	//Audit & admin
	// ID        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	ID        string `json:"_id,omitempty" bson:"_id,omitempty"`
	Account   string
	Signature string
	Timestamp int64

	//Specification
	ImageAMD64       string
	Instances        int
	Replicas         int
	RequiredMBMemory int
	RequiredCPUCores int
	RunUntil         int64
	AllowSuicide     int64

	StartStrategy string
}

//Push saves a contract to MongoDB
func (contract Contract) Push() error {
	err := mongo.Push(contract)
	return err
}

//DbID returns the ID of the contract
// func (contract Contract) DbID() bson.ObjectId {
func (contract Contract) DbID() string {
	return contract.ID
}

//Upsert ..
func (contract Contract) Upsert() error {
	// err := mongo.Upsert(contract)
	err := mongo.UpsertWithID(contract.ID, contract)
	return err
}

//ServiceName returns the service name to be used by this contract
func (contract Contract) ServiceName() string {
	return fmt.Sprintf("service_%s", strings.Replace(contract.ID, "-", "", -1))
}
