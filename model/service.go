package model

import (
	"fmt"

	"github.com/docker/docker/api/types/swarm"
	"vbom.ml/util/sortorder"
)

//Service represents a Docker service
type Service struct {
	ID             string `db:"_id" json:"_id"`
	Name           string
	Image          string
	HasWorked      bool
	EffectiveDate  int64
	Network        string
	HealthyManaged bool
	Replicas       int
	ContractID     string
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
