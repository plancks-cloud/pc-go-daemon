package model

import "git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"

//Contract represents a contract issued to run a container
type Contract struct {
	//Audit & admin
	ID        string `db:"_id" json:"_id" bson:"_id"`
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

//Push saves a contract to MongoDB
func (contract Contract) Push() error {
	err := mongo.Push(contract)
	return err
}
