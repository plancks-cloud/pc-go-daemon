package community

import (
	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller/db"
)

func considerWins() {

	contracts := db.GetContract()
	if len(contracts) == 0 {
		return
	}

	for _, contract := range contracts {
		wins := db.GetWinsByContractID(contract.ID)
		haveIWon, win := db.HaveIWonFromWins(wins)
		if !haveIWon {
			continue
		}
		if db.ServiceExistsByContractId(contract.ID) {
			continue
		}
		db.CreateServiceFromWin(&win)
	}

}
