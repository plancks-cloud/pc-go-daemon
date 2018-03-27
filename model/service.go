package model

import (
	"fmt"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"

	"github.com/docker/docker/api/types/swarm"
	uuid "github.com/nu7hatch/gouuid"
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

//ServiceState models the current running state of a service
type ServiceState struct {
	ID               string
	Name             string
	ReplicasRunning  int
	ReplicasRequired uint64
}

func (service *ServiceState) String() string {
	return fmt.Sprintf("ID: %s, Name: %s, Running: %d, Required: %d", service.ID, service.Name, service.ReplicasRunning, service.ReplicasRequired)
}

//ByName is a struct
type ByName []swarm.Service

func (n ByName) Len() int           { return len(n) }
func (n ByName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n ByName) Less(i, j int) bool { return sortorder.NaturalLess(n[i].Spec.Name, n[j].Spec.Name) }

//Push saves the service object into MongoDB
func (service Service) Push() {
	if len(service.ID) == 0 {
		u, _ := uuid.NewV4()
		service.ID = u.String()
	}
	mongo.Push(service)
}
