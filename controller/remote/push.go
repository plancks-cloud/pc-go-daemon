package remote

import (
	"fmt"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller/db"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/util"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

func syncPushAll(outerWaitGroup sync.WaitGroup) {

	go func(outerWaitGroup sync.WaitGroup) {
		start := time.Now()

		var wg sync.WaitGroup
		wg.Add(4)

		go pushAllWallets(wg)
		go pushAllContracts(wg)
		go pushAllBids(wg)
		go pushAllWins(wg)

		wg.Wait()

		elapsed := time.Since(start)
		log.Infoln(fmt.Sprintf("⏰  PushAll took: %s", elapsed))

		outerWaitGroup.Done()

	}(outerWaitGroup)

}

//PushAllWallets pushes all wallets to cloud
func pushAllWallets(wg sync.WaitGroup) {
	wallets := db.GetWallet()
	body := model.WalletSyncable{Collection: "Wallet", Index: "_id", Indexes: nil, Rows: wallets}
	util.Post(model.DBSyncURL, body.ToJSON())

	wg.Done()

}

//PushAllContracts pushes all contracts to cloud
func pushAllContracts(wg sync.WaitGroup) {
	contracts := db.GetContract()
	if contracts == nil || len(contracts) == 0 {
		return
	}
	body := model.ContractSyncable{Collection: "Contract", Index: "_id", Indexes: nil, Rows: contracts}
	util.Post(model.DBSyncURL, body.ToJSON())

	wg.Done()

}

//PushAllBids pushes all bids to cloud
func pushAllBids(wg sync.WaitGroup) {
	bids := db.GetBid()
	if bids == nil || len(bids) == 0 {
		return
	}
	body := model.BidSyncable{Collection: "Bid", Index: "_id", Indexes: nil, Rows: bids}
	util.Post(model.DBSyncURL, body.ToJSON())

	wg.Done()

}

//PushAllWins pushes all wins to cloud
func pushAllWins(wg sync.WaitGroup) {
	wins := db.GetWin()
	if wins == nil || len(wins) == 0 {
		return
	}
	body := model.WinSyncable{Collection: "Win", Index: "_id", Indexes: nil, Rows: wins}
	util.Post(model.DBSyncURL, body.ToJSON())

	wg.Done()

}
