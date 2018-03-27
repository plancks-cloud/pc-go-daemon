package model

import (
	"fmt"
	"strings"

	"github.com/nu7hatch/gouuid"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
)

//Contract represents a contract issued to run a container
type Contract struct {
	//Audit & admin
	// ID        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	ID        string `json:"_id,omitempty" bson:"_id,omitempty"`
	Account   string `json:"account" bson:"account,omitempty"`
	Signature string `json:"signature" bson:"signature,omitempty"`
	Timestamp int64  `json:"timestamp" bson:"timestamp,omitempty"`

	//Specification
	ImageAMD64       string `json:"imageAMD64" bson:"imageAMD64,omitempty"`
	Instances        int    `json:"instances" bson:"instances,omitempty"`
	Replicas         int    `json:"replicas" bson:"replicas,omitempty"`
	RequiredMBMemory int    `json:"requiredMBMemory" bson:"requiredMBMemory,omitempty"`
	RequiredCPUCores int    `json:"requiredCPUCores" bson:"requiredCPUCores,omitempty"`
	RunUntil         int64  `json:"runUntil" bson:"runUntil,omitempty"`
	AllowSuicide     int64  `json:"allowSuicide" bson:"allowSuicide,omitempty"`
	StartStrategy    string `json:"startStrategy" bson:"startStrategy,omitempty"`
	ServiceName      string `json:"serviceName" bson:"serviceName,omitempty"`
}

//Push saves a contract to MongoDB
func (contract Contract) Push() error {
	if len(contract.ID) == 0 {
		u, _ := uuid.NewV4()
		contract.ID = u.String()
	}
	if len(contract.ServiceName) == 0 {
		contract.ServiceName = contract.GetServiceName()
	}
	err := mongo.Push(contract)
	return err
}

//DbID returns the ID of the contract
func (contract Contract) DbID() string {
	return contract.ID
}

//Upsert ..
func (contract Contract) Upsert() error {
	err := mongo.UpsertWithID(contract.ID, contract)
	return err
}

//GetServiceName returns the service name to be used by this contract
func (contract Contract) GetServiceName() string {
	return fmt.Sprintf("service_%s", strings.Replace(contract.ID, "-", "", -1))
}
