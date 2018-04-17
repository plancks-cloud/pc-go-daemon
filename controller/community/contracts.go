package community

import (
	"fmt"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller/db"
	log "github.com/sirupsen/logrus"
)

func considerContracts() {

	contracts := db.GetContract()
	for _, contract := range contracts {

		//Ignore ancient contracts
		if db.ExpiredContract(&contract) {
			continue
		}

		//Ignore cancelled contracts
		cancels := db.GetCancelContractsByContractID(contract.ID)
		if len(cancels) > 0 {
			continue
		}

		//Look at the wins
		wins := db.GetWinsByContractID(contract.ID)
		if len(wins) > 0 {
			isWon, win := db.HaveIWonFromWins(wins)
			if isWon {
				log.Debugln(fmt.Sprintf("🏆  I'm the winner of this contract %s", contract.ID))
				db.CreateServiceFromWin(&win)
			}
			continue
		}

		//Did I bid
		if !db.HaveIBidOnContract(contract.ID) {
			db.CreateBidFromContract(contract)
		}

		db.CheckForWinsNow(contract)

	}

}
