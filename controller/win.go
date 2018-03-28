package controller

import (
	"fmt"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
)

//CreateWin saves a win
func CreateWin(item *model.Win) model.MessageOK {
	err := item.Push()
	if err != nil {
		log.Errorln(fmt.Sprintf("Error saving win: %s", err))
		return model.OkMessage(false, err.Error())
	}
	return model.Ok(true)
}

//GetWin returns all wins stored in the datastore
func GetWin() []model.Win {
	var wins []model.Win
	mongo.GetCollection(model.Win{}).Find(nil).All(&wins)
	return wins
}

//GetWinsByContractID returns all wins for a contract
func GetWinsByContractID(contractID string) []model.Win {
	var wins []model.Win
	mongo.GetCollection(model.Bid{}).Find(bson.M{"contractId": bson.ObjectIdHex(contractID)}).All(&wins)
	for _, row := range wins {
		log.Infoln(fmt.Sprintf("Item: %s", row.ID))
	}
	return wins
}

//GetOneWin returns a single win
func GetOneWin(id string) (model.Win, error) {
	var win model.Win
	err := mongo.GetCollection(&win).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&win)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error getting win: %s", err))
	}
	return win, err
}
