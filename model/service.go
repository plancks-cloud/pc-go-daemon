package model

import (
	"fmt"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/util"
	"github.com/docker/docker/api/types/swarm"
	"github.com/nu7hatch/gouuid"
	"vbom.ml/util/sortorder"
)

//Service represents a Docker service
type Service struct {
	ID             string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name           string `json:"name" bson:"name,omitempty"`
	Image          string `json:"image" bson:"image,omitempty"`
	HasWorked      bool   `json:"hasWorked" bson:"hasWorked,omitempty"`
	EffectiveDate  int64  `json:"effectiveDate" bson:"effectiveDate,omitempty"`
	Network        string `json:"network" bson:"network,omitempty"`
	HealthyManaged bool   `json:"healthyManaged" bson:"healthyManaged,omitempty"`
	Replicas       int    `json:"replicas" bson:"replicas,omitempty"`
	ContractID     string `json:"contractId" bson:"contractId"`
}

//Push saves a bid to Mongodb
func (service *Service) Push() (err error) {
	if len(service.ID) == 0 {
		u, _ := uuid.NewV4()
		service.ID = u.String()
	}
	err = mongo.Push(service)
	return
}

//Upsert updates a document if it exists, otherwise inserts
func (service *Service) Upsert() error {
	return mongo.UpsertWithID(service.ID, service)
}

//Expired checks if a service should still be running according to the contract it was created with
func (service *Service) Expired(contract *Contract) bool {
	now := util.MakeTimestamp()
	if contract.RunUntil == 0 {
		return false
	}
	return now > contract.RunUntil
}

//ServiceState models the current running state of a service
type ServiceState struct {
	ID               string
	Name             string
	ReplicasRunning  int
	ReplicasRequired uint64
}

//ServiceStateResult is a small struct for the client to visualize
type ServiceStateResult struct {
	Service      Service `json:"service"`
	ReplicasLive int     `json:"replicasLive"`
}

func (service *ServiceState) String() string {
	return fmt.Sprintf("ID: %s, Name: %s, Running: %d, Required: %d", service.ID, service.Name, service.ReplicasRunning, service.ReplicasRequired)
}

//ByName is a struct
type ByName []swarm.Service

func (n ByName) Len() int           { return len(n) }
func (n ByName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n ByName) Less(i, j int) bool { return sortorder.NaturalLess(n[i].Spec.Name, n[j].Spec.Name) }
