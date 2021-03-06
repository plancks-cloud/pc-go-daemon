package model

import (
	"encoding/json"
	"fmt"
	"strings"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/mem"
	"github.com/nu7hatch/gouuid"
	log "github.com/sirupsen/logrus"
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
	SecondsToLive    int64  `json:"secondsToLive" bson:"secondsToLive,omitempty"`
	AllowSuicide     bool   `json:"allowSuicide" bson:"allowSuicide,omitempty"`
	StartStrategy    string `json:"startStrategy" bson:"startStrategy,omitempty"`
	ServiceName      string `json:"serviceName" bson:"serviceName,omitempty"`
}

//ContractSyncable holds info for syncing with the cloud
type ContractSyncable struct {
	Collection string     `json:"collection" bson:"collection"`
	Index      string     `json:"index" bson:"index"`
	Indexes    []string   `json:"indexes" bson:"indexes"`
	Rows       []Contract `json:"rows" bson:"rows"`
}

//ContractResult helps the client see the state of a contract
type ContractResult struct {
	Contract        `json:"contract"`
	Bids            []Bid            `json:"bids"`
	Wins            []Win            `json:"wins"`
	CancelContracts []CancelContract `json:"cancelContracts"`
}

//ToJSON converts an object to json
func (contractSyncable ContractSyncable) ToJSON() []byte {
	jsonBytes, jsonError := json.Marshal(contractSyncable)
	if jsonError != nil {
		log.Fatalln(fmt.Sprintf("Error converting contractSyncable to json: %s", jsonError))
	}
	return jsonBytes

}

//Push saves a contract to Mongodb
func (contract Contract) Push() (err error) {
	if len(contract.ID) == 0 {
		u, _ := uuid.NewV4()
		contract.ID = u.String()
	}
	if len(contract.ServiceName) == 0 {
		contract.ServiceName = contract.GetServiceName()
	}
	err = mem.Push(contract)
	return
}

//DbID returns the ID of the contract
func (contract Contract) DbID() string {
	return contract.ID
}

//Upsert ..
func (contract Contract) Upsert() error {
	err := mem.Push(contract)
	return err
}

//GetServiceName returns the service name to be used by this contract
func (contract Contract) GetServiceName() string {
	return fmt.Sprintf("service_%s", strings.Replace(contract.ID, "-", "", -1))
}
