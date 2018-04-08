package community

import (
	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller/db"
	log "github.com/sirupsen/logrus"
	"sync"
)

func considerContracts(wg sync.WaitGroup) {

	contracts := db.GetContract()
	for _, contract := range contracts {

		//Ignore ancient contracts
		if db.ExpiredContract(&contract) {
			continue
		}

		//Look at the wins
		wins := db.GetWinsByContractID(contract.ID)
		if len(wins) > 0 {
			isWon, win := db.HaveIWonFromWins(wins)
			if isWon {
				log.Infoln("🏆  I'm the winner of this contract %s", contract.ID)
				db.CreateServiceFromWin(&win)
			}

			//Stop processing if there is a win
			db.CheckForWinsNow(contract) // This needs to be channelled
			continue
		}

		//Did I bid
		if db.HaveIBidOnContract(contract.ID) {
			continue
		}
		db.CreateBidFromContract(contract)

	}

	wg.Done()
}
