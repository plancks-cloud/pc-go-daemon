package model

//Win object represents bid document in DB
type Win struct {
	ID            string `json:"_id,omitempty" bson:"_id,omitempty"`
	ContractID    string `json:"contractId"`
	WinnerAccount string
	timestamp     int64
	signature     string
}
