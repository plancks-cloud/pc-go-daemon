package controller

import (
	"encoding/json"
	"fmt"
	"time"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/util"
	log "github.com/sirupsen/logrus"
)

//SyncDatabase sets up scheduler to sync up databases
func SyncDatabase() {
	go func() {
		//Wake up the function
		util.Options(model.DBSyncURL)
		util.Options(model.DBGCURL)

		for {
			//Sync and sleep
			log.Debugln(fmt.Sprintf("⏰  Time to sync database"))
			PullAll()
			garbageCollectAll()
			PushAll()
			garbageCollectRemote()
			time.Sleep(30 * time.Second)
		}
	}()
}

//ReconServicesNow runs the recon right now
func ReconServicesNow() {
	reconServices()
}

//ReconServices sets up scheduler to recon docker services running
func ReconServices() {
	go func() {
		time.Sleep(15 * time.Second)
		for {
			log.Debugln(fmt.Sprintf("⏰  Time to recon services"))
			reconServices()
			time.Sleep(30 * time.Second)
		}
	}()
}

//PushAll gets all rows in DB and pushes to DB
func PushAll() {
	PushAllWallets()
	PushAllContracts()
	PushAllBids()
	PushAllWins()
}

//PushAllWallets pushes all wallets to cloud
func PushAllWallets() {
	wallets := GetWallet()
	body := model.WalletSyncable{Collection: "Wallet", Index: "_id", Indexes: nil, Rows: wallets}
	util.Post(model.DBSyncURL, body.ToJSON())

}

//PushAllContracts pushes all contracts to cloud
func PushAllContracts() {
	contracts := GetContract()
	if contracts == nil || len(contracts) == 0 {
		return
	}
	body := model.ContractSyncable{Collection: "Contract", Index: "_id", Indexes: nil, Rows: contracts}
	util.Post(model.DBSyncURL, body.ToJSON())

}

//PushAllBids pushes all bids to cloud
func PushAllBids() {
	bids := GetBid()
	if bids == nil || len(bids) == 0 {
		return
	}
	body := model.BidSyncable{Collection: "Bid", Index: "_id", Indexes: nil, Rows: bids}
	util.Post(model.DBSyncURL, body.ToJSON())

}

//PushAllWins pushes all wins to cloud
func PushAllWins() {
	wins := GetWin()
	if wins == nil || len(wins) == 0 {
		return
	}
	body := model.WinSyncable{Collection: "Win", Index: "_id", Indexes: nil, Rows: wins}
	util.Post(model.DBSyncURL, body.ToJSON())

}

func garbageCollectAll() {
	contracts := GetContract()
	log.Debugln(fmt.Sprintf("⏰  Time to GC"))
	for _, item := range contracts {
		//Check if ancient
		log.Debugln(fmt.Sprintf("⏰  .. Checking %s", item.ID))

		//TODO: check for cancelled contracts
		if ExpiredContract(&item) {
			log.Debugln(fmt.Sprintf("⏰  .. EXPIRED! %s", item.ID))
			//Remove
			DeleteContract(&item)
			DeleteBidsByContractID(item.ID)
			DeleteWinsByContractID(item.ID)
			DeleteServicesByContractID(item.ID)

		}
	}
}

func garbageCollectRemote() {
	go util.Get(model.DBGCURL)
}

//PullAll gets all rows in cloud DB and puts them in local DB
func PullAll() {
	contracts := PullAllContracts()
	for _, contract := range contracts {
		if ContractExists(contract.ID) {
			//Ignore
			continue
		}
		contract.Upsert()
		CallbackContractAsync(contract, true)
	}
	wallets := PullAllWallets()
	for _, item := range wallets {
		item.Upsert()
	}
	bids := PullAllBids()
	for _, item := range bids {
		if ContractExists(item.ContractID) {
			item.Upsert()
		}
	}
	wins := PullAllWins()
	for _, item := range wins {
		if ContractExists(item.ContractID) {
			item.Upsert()
			CallbackWinAsync(item)
		}
	}
}

//PullAllContracts gets all contracts in the cloud DB
func PullAllContracts() (contracts []model.Contract) {
	typeName := util.GetType(model.Contract{})
	url := fmt.Sprintf("%s?collection=%s", model.DBSyncURL, typeName)
	resp, err := util.Get(url)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error getting contacts during sync: %s", err))
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&contracts)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error decoding contract during sync: %s", err))
		return
	}
	return
}

//PullAllWallets gets all wallets in the cloud DB
func PullAllWallets() (wallets []model.Wallet) {
	typeName := util.GetType(model.Wallet{})
	url := fmt.Sprintf("%s?collection=%s", model.DBSyncURL, typeName)
	resp, err := util.Get(url)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error getting wallets during sync: %s", err))
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&wallets)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error decoding wallets during sync: %s", err))
		return
	}
	return
}

//PullAllBids gets all bids in the cloud DB
func PullAllBids() (bids []model.Bid) {
	typeName := util.GetType(model.Bid{})
	url := fmt.Sprintf("%s?collection=%s", model.DBSyncURL, typeName)
	resp, err := util.Get(url)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error getting bids during sync: %s", err))
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&bids)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error decoding bids during sync: %s", err))
		return
	}
	return
}

//PullAllWins gets all wins in the cloud DB
func PullAllWins() (wins []model.Win) {
	typeName := util.GetType(model.Win{})
	url := fmt.Sprintf("%s?collection=%s", model.DBSyncURL, typeName)
	resp, err := util.Get(url)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error getting wins during sync: %s", err))
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&wins)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error decoding wins during sync: %s", err))
		return
	}
	return
}
