package controller

import (
	"fmt"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
)

//CreateBid creates a new contract
func CreateBid(bid *model.Bid) model.MessageOK {
	err := bid.Push()
	if err != nil {
		log.Errorln(fmt.Sprintf("Error saving bid: %s", err))
		return model.OkMessage(false, err.Error())
	}
	return model.Ok(true)
}

//GetBid returns all contracts stored in the datastore
func GetBid() []model.Bid {
	var bids []model.Bid
	mongo.GetCollection(model.Bid{}).Find(nil).All(&bids)
	for _, bid := range bids {
		log.Infoln(fmt.Sprintf("Bid: %s", bid.ID))
	}
	return bids
}

//GetBidsByContractID returns all bids for a contract
func GetBidsByContractID(contractID string) []model.Bid {
	var bids []model.Bid
	mongo.GetCollection(model.Bid{}).Find(bson.M{"contractId": bson.ObjectIdHex(contractID)}).All(&bids)
	for _, bid := range bids {
		log.Infoln(fmt.Sprintf("Bid: %s", bid.ID))
	}
	return bids
}

//GetOneBid returns a single contract
func GetOneBid(id string) (model.Bid, error) {
	var bid model.Bid
	err := mongo.GetCollection(&bid).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&bid)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error getting bid: %s", err))
	}
	return bid, err
}

//UpdateBid upserts the given bid
func UpdateBid(bid *model.Bid) error {
	err := bid.Upsert()
	return err
}
