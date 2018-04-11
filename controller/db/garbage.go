package db

import (
	"fmt"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	log "github.com/sirupsen/logrus"
	"time"
)

//LocalGC cleans up the local database
func LocalGC() {

	go func() {

		start := time.Now()
		contracts := GetContract()
		log.Debugln(fmt.Sprintf("⏰  Time to GC"))
		for _, item := range contracts {
			//Check if ancient
			log.Debugln(fmt.Sprintf("⏰  .. Checking %s", item.ID))

			//TODO: check for cancelled contracts
			if ExpiredContractBy(&item, model.AncientAgeSeconds) {
				log.Debugln(fmt.Sprintf("⏰  .. EXPIRED! %s", item.ID))
				//Remove
				DeleteContract(&item)
				DeleteBidsByContractID(item.ID)
				DeleteWinsByContractID(item.ID)
				DeleteServicesByContractID(item.ID)

			}
		}
		elapsed := time.Since(start)
		log.Infoln(fmt.Sprintf("⏰  Local GC took %s", elapsed))

	}()

}
