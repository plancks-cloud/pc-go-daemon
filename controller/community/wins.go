package community

import (
	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller/db"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"time"
)

func considerWins() {

	//Check that I should be deciding who won
	duration := time.Since(model.StartupTime)
	secSinceStart := duration.Seconds()
	if secSinceStart < model.SecondsBeforeWinDeclarer {
		return
	}

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
