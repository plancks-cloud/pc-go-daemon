package model

//Bid object represents bid document in DB
type Bid struct {
	ID          string `json:"_id,omitempty" bson:"_id,omitempty"`
	ContractID  string `json:"contractId"`
	Votes       int
	FromAccount string
	Timestamp   int64
	Signature   string
}
