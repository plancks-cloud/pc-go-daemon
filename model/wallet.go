package model

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
func (wallet Wallet) Push() error {
	err := mongo.Push(wallet)
	return err
}
