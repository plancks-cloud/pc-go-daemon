package controller

import (
	"fmt"
	"time"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/util"
	"github.com/globalsign/mgo/bson"
	"github.com/nu7hatch/gouuid"
	log "github.com/sirupsen/logrus"
	"sort"
)

//GetWin returns all wins stored in the datastore
func GetWin() (wins []model.Win) {
	mongo.GetCollection(model.Win{}).Find(nil).All(&wins)
	return wins
}

//GetWinsByContractID returns all wins for a contract
func GetWinsByContractID(contractID string) (wins []model.Win) {
	mongo.GetCollection(model.Win{}).Find(bson.M{"contractId": contractID}).All(&wins)
	return wins
}

//CheckForWinsLater announces winners where relevant
func CheckForWinsLater(contract model.Contract) {
	log.Infoln(fmt.Sprintf("> ⏰ ⏰ ⏰  Going to check for wins in two minutes: %s ", contract.ID))
	time.Sleep(2 * time.Minute)
	CheckForWinsNow(contract)

}

//CheckForWins announces winners where relevant
func CheckForWinsNow(contract model.Contract) {
	log.Infoln("win controller: CheckForWins")
	someTimeAfterContract := contract.Timestamp + (1000 * 60 * 2.5)
	now := util.MakeTimestamp()
	log.Infoln("WinController: Was comparing... %s and %d", someTimeAfterContract, now)
	log.Infoln("WinController: Now is: ", now)
	log.Infoln("WinController: Waitin:", someTimeAfterContract, now)

	//if now < someTimeAfterContract {
	//	log.Infoln(fmt.Sprintf("> Too early to find a winner. Stopping: %s ", contract.ID))
	//	return
	//}
	//log.Infoln(fmt.Sprintf("> Its been more than n minutes. We can announce a winner. ID: %s ", contract.ID))

	bids := GetBidsByContractID(contract.ID)
	if len(bids) == 0 {
		//No bids - no winner
		log.Infoln(fmt.Sprintf("> No bids found. For now, no winner on contract, ID: %s ", contract.ID))
		return
	}

	sort.Sort(model.ByVotes(bids))
	winnerCount := 0
	for winner := 0; winner < contract.Instances; winner++ {
		log.Infoln(fmt.Sprintf("> Going to say the winner was: %s", bids[winner].FromAccount))
		CreateWinFromContract(bids[winner].FromAccount, contract)
		winnerCount++
	}
	log.Infoln(fmt.Sprintf("> # of winners: %i", winnerCount))

	if winnerCount == 0 {
		log.Infoln(fmt.Sprintf("> This should never happen. No highest bid: %s ", contract.ID))
	}

}

//CreateWinFromContract creates win
func CreateWinFromContract(winnerID string, contract model.Contract) {
	log.Infoln("win controller: CreateWinFromContract")
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

//CallbackWinAsync checks an incoming DB row to see if it is interesting
func CallbackWinAsync(win model.Win) {

	//Check if expired first.
	contract, _  := GetOneContract(win.ContractID)
	//Should be there... if the win is there
	if win.Expired(&contract) {
		//Ignore
		return
	}

	go CheckIfIWon(win)
}

//CheckIfIWon if I won will take the next steps if needed
func CheckIfIWon(win model.Win) {
	if model.SystemWallet.ID == win.WinnerAccount {
		log.Infoln("🏆 🏆 🏆   I'm the winner of this contract %s", win.ContractID)
		CreateServiceFromWin(&win)
	}
}
