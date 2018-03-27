package model

<<<<<<< HEAD
import "git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"

//Wallet represents a unique node
type Wallet struct {
	ID         string `db:"_id" json:"_id"`
	PrivateKey string
	PublicKey  string
}

//GetSignature will get sign an object
func (wallet *Wallet) GetSignature(obj interface{}) string {
	return wallet.ID
}

//Push saves a wallet to MongoDB
=======
import (
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
)

//Wallet is the issuing party, as well as a node running a container
type Wallet struct {
	ID         string `json:"id,omitempty" bson:"_id,omitempty"`
	PublicKey  string
	PrivateKey string
}

//Push persists the wallet to the database
>>>>>>> origin/master
func (wallet Wallet) Push() error {
	err := mongo.Push(wallet)
	return err
}
<<<<<<< HEAD
=======

//Signature returns the wallet ID
func (wallet Wallet) Signature() string {
	return wallet.ID
}

//Upsert updates a document if it exists, otherwise inserts
func (wallet Wallet) Upsert() error {
	return mongo.UpsertWithID(wallet.ID, wallet)
}
>>>>>>> origin/master
