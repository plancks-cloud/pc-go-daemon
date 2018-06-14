package remote

import (
	"fmt"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller/db"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"sync"
	"time"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/util/http"
	log "github.com/sirupsen/logrus"
)

func syncPushAll(outerWaitGroup *sync.WaitGroup) {

	go func(outerWaitGroup *sync.WaitGroup) {
		start := time.Now()

		var wg sync.WaitGroup
		wg.Add(4)

		go pushAllWallets(&wg)
		go pushAllContracts(&wg)
		go pushAllBids(&wg)
		go pushAllWins(&wg)

		wg.Wait()

		elapsed := time.Since(start)
		log.Infoln(fmt.Sprintf("⏰  PushAll took: %s", elapsed))

		outerWaitGroup.Done()

	}(outerWaitGroup)

}

//PushAllWallets pushes all wallets to cloud
func pushAllWallets(wg *sync.WaitGroup) {
	start := time.Now()
	ch := db.GetWallet()
	var wallets []model.Wallet
	for _, w := range ch {
		wallets = append(wallets, w)
	}
	if len(wallets) > 0 {
		body := model.WalletSyncable{Collection: "Wallet", Index: "_id", Indexes: nil, Rows: wallets}
		http.Post(model.DBSyncURL, body.ToJSON())
	}

	elapsed := time.Since(start)
	log.Debugln(fmt.Sprintf("⏰  PushAll-wallets took: %s", elapsed))
	wg.Done()

}

//PushAllContracts pushes all contracts to cloud
func pushAllContracts(wg *sync.WaitGroup) {
	start := time.Now()
	contracts := db.GetContract()
	if len(contracts) > 0 {
		body := model.ContractSyncable{Collection: "Contract", Index: "_id", Indexes: nil, Rows: contracts}
		http.Post(model.DBSyncURL, body.ToJSON())
	}

	elapsed := time.Since(start)
	log.Debugln(fmt.Sprintf("⏰  PushAll-contracts took: %s", elapsed))
	wg.Done()

}

//PushAllBids pushes all bids to cloud
func pushAllBids(wg *sync.WaitGroup) {
	start := time.Now()
	bids := db.GetBid()
	if len(bids) > 0 {
		body := model.BidSyncable{Collection: "Bid", Index: "_id", Indexes: nil, Rows: bids}
		http.Post(model.DBSyncURL, body.ToJSON())
	}

	elapsed := time.Since(start)
	log.Debugln(fmt.Sprintf("⏰  PushAll-bids took: %s", elapsed))
	wg.Done()

}

//PushAllWins pushes all wins to cloud
func pushAllWins(wg *sync.WaitGroup) {
	start := time.Now()
	wins := db.GetWin()
	if len(wins) > 0 {
		body := model.WinSyncable{Collection: "Win", Index: "_id", Indexes: nil, Rows: wins}
		http.Post(model.DBSyncURL, body.ToJSON())
	}

	elapsed := time.Since(start)
	log.Debugln(fmt.Sprintf("⏰  PushAll-wins took: %s", elapsed))
	wg.Done()

}
