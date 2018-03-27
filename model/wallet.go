package model

import (
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
	uuid "github.com/nu7hatch/gouuid"
)

//Wallet is the issuing party, as well as a node running a container
type Wallet struct {
	ID         string `json:"id,omitempty" bson:"_id,omitempty"`
	PublicKey  string
	PrivateKey string
}

//Push persists the wallet to the database
func (wallet Wallet) Push() error {
	if len(wallet.ID) == 0 {
		u, _ := uuid.NewV4()
		wallet.ID = u.String()
	}
	err := mongo.Push(wallet)
	return err
}

//Signature returns the wallet ID
func (wallet Wallet) Signature() string {
	return wallet.ID
}

//Upsert updates a document if it exists, otherwise inserts
func (wallet Wallet) Upsert() error {
	return mongo.UpsertWithID(wallet.ID, wallet)
}
