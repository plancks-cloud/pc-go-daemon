package model

//Contract represents a contract issues to run a container
type Contract struct {
	//Audit & admin
	ID        string `db:"_id" json:"_id"`
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
