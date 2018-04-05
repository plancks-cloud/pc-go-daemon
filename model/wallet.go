package model

import (
	"encoding/json"
	"fmt"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
	"github.com/nu7hatch/gouuid"
	log "github.com/sirupsen/logrus"
)

//Wallet is the issuing party, as well as a node running a container
type Wallet struct {
	ID         string `json:"_id,omitempty" bson:"_id,omitempty"`
	PublicKey  string `json:"publicKey" bson:"publicKey,omitempty"`
	PrivateKey string `json:"privateKey" bson:"privateKey,omitempty"`
	Signature  string `json:"signature" bson:"signature,omitempty"`
}

//WalletSyncable is a wrapper for what get posted to the cloud
type WalletSyncable struct {
	Collection string   `json:"collection" bson:"collection"`
	Index      string   `json:"index" bson:"index"`
	Indexes    []string `json:"indexes" bson:"indexes"`
	Rows       []Wallet `json:"rows" bson:"rows"`
}

//ToJSON converts object o json
func (walletSyncable WalletSyncable) ToJSON() []byte {
	jsonBytes, jsonError := json.Marshal(walletSyncable)
	if jsonError != nil {
		log.Fatalln(fmt.Sprintf("Error converting walletSyncable to json: %s", jsonError))
	}
	return jsonBytes

}

//Push persists the wallet to the database
func (wallet Wallet) Push() (err error) {
	if len(wallet.ID) == 0 {
		u, _ := uuid.NewV4()
		wallet.ID = u.String()
	}
	err = mongo.Push(wallet)
	return
}

//GetSignature returns the wallet ID
func (wallet Wallet) GetSignature() string {
	return wallet.ID
}

//Upsert updates a document if it exists, otherwise inserts
func (wallet Wallet) Upsert() error {
	return mongo.UpsertWithID(wallet.ID, wallet)
}
