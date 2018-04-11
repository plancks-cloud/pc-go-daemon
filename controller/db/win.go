package db

import (
	"fmt"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/util"
	"github.com/globalsign/mgo/bson"
	"github.com/nu7hatch/gouuid"
	log "github.com/sirupsen/logrus"
	"sort"
)

//GetWin returns all wins stored in the Datastore
func GetWin() (wins []model.Win) {
	mongo.GetCollection(model.Win{}).Find(nil).All(&wins)
	return
}

//GetWinsByContractID returns all wins for a contract
func GetWinsByContractID(contractID string) (wins []model.Win) {
	mongo.GetCollection(model.Win{}).Find(bson.M{"contractId": contractID}).All(&wins)
	return
}

//CheckForWinsNow announces winners where relevant
func CheckForWinsNow(contract model.Contract) {
	log.Debugln("win controller: CheckForWins")
	ripeTime := contract.Timestamp + (1000 * model.WinnerAgeSeconds)
	now := util.MakeTimestamp()

	//If now is before the time we need
	if now < ripeTime {
		log.Debugln(fmt.Sprintf("> Too early to find a winner. Stopping: %s ", contract.ID))
		return
	}
	log.Debugln(fmt.Sprintf("> Its been more than n minutes. We can announce a winner. ID: %s ", contract.ID))

	winsForContract := GetWinsByContractID(contract.ID)
	if len(winsForContract) > 0 {
		//Wins have already been declared..
		log.Debugln(fmt.Sprintf("> Looks like there are wins for this contract already. Stopping: %s ", contract.ID))
		return
	}
	log.Infoln(fmt.Sprintf("> No wins and is ripe.. Will decide winner!. ID: %s ", contract.ID))

	bids := GetBidsByContractID(contract.ID)
	if len(bids) == 0 {
		//No bids - no winner
		log.Debugln(fmt.Sprintf("> No bids found. For now, no winner on contract, ID: %s ", contract.ID))
		return
	}
	log.Infoln(fmt.Sprintf("> # of votes ID: %d ", len(bids)))

	sort.Sort(model.ByVotes(bids))
	winnerCount := 0
	for winner := 0; winner < contract.Instances; winner++ {

		//Ensure there are enough bids
		if winner >= len(bids) {
			//Out of bounds
			continue
		}
		log.Infoln(fmt.Sprintf("> Going to say the winner was: %s", bids[winner].FromAccount))
		CreateWinFromContract(bids[winner].FromAccount, contract)
		winnerCount++
	}
	log.Infoln(fmt.Sprintf("> # of winners: %d ", winnerCount))

	if winnerCount == 0 {
		log.Error(fmt.Sprintf("> This should never happen. No highest bid: %s ", contract.ID))
	}

}

//CreateWinFromContract creates win
func CreateWinFromContract(winnerID string, contract model.Contract) {
	log.Debugln("win controller: CreateWinFromContract")
	uuidString, _ := uuid.NewV4()
	win := model.Win{
		ID:            uuidString.String(),
		ContractID:    contract.ID,
		WinnerAccount: winnerID,
		Timestamp:     util.MakeTimestamp(),
		Signature:     model.SystemWallet.GetSignature()}
	win.Upsert()
	CheckIfIWon(win)

}

func HaveIWonFromWins(wins []model.Win) (bool, model.Win) {
	for _, win := range wins {
		if HaveIWonFromWin(win) {
			return true, win
		}
	}
	return false, model.Win{}
}
func HaveIWonFromWin(win model.Win) bool {
	return model.SystemWallet.ID == win.WinnerAccount

}

//CheckIfIWon if I won will take the next steps if needed
func CheckIfIWon(win model.Win) {
	if HaveIWonFromWin(win) {
		log.Debugln(fmt.Sprintf("🏆  I'm the winner of this contract %s", win.ContractID))
		CreateServiceFromWin(&win)
	}
}

func DeleteWinsByContractID(id string) {
	win := model.Win{}
	_, err := mongo.GetCollection(&win).RemoveAll(bson.M{"contractId": id})
	if err != nil {
		log.Errorln(fmt.Sprintf("Error deleting wins by contractId: %s", err))
	}

}
