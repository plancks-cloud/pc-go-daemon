package controller

import (
	"fmt"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/util"
	"github.com/globalsign/mgo/bson"
	"github.com/nu7hatch/gouuid"
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

//CheckForWins announces winners where relavant
func CheckForWins(contract model.Contract) {
	bids := GetBid()
	if len(bids) == 0 {
		//No bids - no winner
		return
	}

	//TODO: better impl
	winnerVotes := -1
	winnerID := ""

	for _, element := range bids {
		if element.Votes > winnerVotes {
			winnerVotes = element.Votes
			winnerID = element.ID
		}
	}

	if winnerVotes != -1 {
		CreateWinFromContract(winnerID, contract)
	}

}

//CreateWinFromContract creates win
func CreateWinFromContract(winnerID string, contract model.Contract) {
	uuidString, _ := uuid.NewV4()
	win := model.Win{
		ID:            uuidString.String(),
		ContractID:    contract.ID,
		WinnerAccount: winnerID,
		Timestamp:     util.MakeTimestamp(),
		Signature:     winnerID}
	win.Upsert()
	CheckIfIWon(win)

}

//CheckIfIWon if I won will take the next steps if needed
func CheckIfIWon(win model.Win) {
	if model.SystemWallet.ID == win.WinnerAccount {
		CreateServiceFromWin(&win)
	}
}
