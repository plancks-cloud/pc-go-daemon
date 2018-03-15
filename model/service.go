package model

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
