package remote

import (
	"encoding/json"
	"fmt"
	"time"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller/db"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/util"
	log "github.com/sirupsen/logrus"
	"sync"
)

func syncPullAll(outerWaitGroup *sync.WaitGroup) {

	go func(outerWaitGroup *sync.WaitGroup) {

		var wg sync.WaitGroup
		wg.Add(4)

		startMethod := time.Now()
		go pullAndStoreAllContracts(&wg)
		go pullAndStoreAllWallets(&wg)
		go pullAndStoreAllBids(&wg)
		go pullAndStoreAllWins(&wg)

		wg.Wait()

		elapsed := time.Since(startMethod)
		log.Infoln(fmt.Sprintf("⏰  PullAll took: %s", elapsed))

		//Kick off a local GC
		db.LocalGC()

		outerWaitGroup.Done()

	}(outerWaitGroup)

}

func pullAndStoreAllContracts(wg *sync.WaitGroup) {
	start := time.Now()
	contracts := pullAllContracts()
	for _, contract := range contracts {
		if db.ContractExists(contract.ID) {
			//Ignore
			continue
		}
		contract.Upsert()
	}
	elapsed := time.Since(start)
	log.Infoln(fmt.Sprintf("⏰  PullAll-contracts took: %s", elapsed))
	wg.Done()
}
func pullAllContracts() (contracts []model.Contract) {
	typeName := util.GetType(model.Contract{})
	url := fmt.Sprintf("%s?collection=%s", model.DBSyncURL, typeName)
	resp, err := util.Get(url)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error getting contacts during remote: %s", err))
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&contracts)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error decoding contract during remote: %s", err))
		return
	}
	return
}

func pullAndStoreAllWallets(wg *sync.WaitGroup) {
	start := time.Now()
	wallets := pullAllWallets()
	for _, item := range wallets {
		item.Upsert()
	}
	elapsed := time.Since(start)
	log.Infoln(fmt.Sprintf("⏰  PullAll-wallets took: %s", elapsed))
	wg.Done()
}
func pullAllWallets() (wallets []model.Wallet) {
	typeName := util.GetType(model.Wallet{})
	url := fmt.Sprintf("%s?collection=%s", model.DBSyncURL, typeName)
	resp, err := util.Get(url)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error getting wallets during remote: %s", err))
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&wallets)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error decoding wallets during remote: %s", err))
		return
	}
	return
}

func pullAndStoreAllBids(wg *sync.WaitGroup) {
	start := time.Now()
	bids := pullAllBids()
	for _, item := range bids {
		if db.ContractExists(item.ContractID) {
			item.Upsert()
		}
	}
	elapsed := time.Since(start)
	log.Infoln(fmt.Sprintf("⏰  PullAll-bids took: %s", elapsed))
	wg.Done()
}
func pullAllBids() (bids []model.Bid) {
	typeName := util.GetType(model.Bid{})
	url := fmt.Sprintf("%s?collection=%s", model.DBSyncURL, typeName)
	resp, err := util.Get(url)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error getting bids during remote: %s", err))
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&bids)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error decoding bids during remote: %s", err))
		return
	}
	return
}

func pullAndStoreAllWins(wg *sync.WaitGroup) {
	start := time.Now()
	wins := pullAllWins()
	for _, item := range wins {
		if db.ContractExists(item.ContractID) {
			item.Upsert()
		}
	}
	elapsed := time.Since(start)
	log.Infoln(fmt.Sprintf("⏰  PullAll-wins took: %s", elapsed))
	wg.Done()

}
func pullAllWins() (wins []model.Win) {
	typeName := util.GetType(model.Win{})
	url := fmt.Sprintf("%s?collection=%s", model.DBSyncURL, typeName)
	resp, err := util.Get(url)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error getting wins during remote: %s", err))
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&wins)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error decoding wins during remote: %s", err))
		return
	}
	return
}
