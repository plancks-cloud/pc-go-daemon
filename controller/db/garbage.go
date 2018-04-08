package db

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

const AncientAgeSeconds = 300

func LocalGC() {

	go func() {

		start := time.Now()
		contracts := GetContract()
		log.Debugln(fmt.Sprintf("⏰  Time to GC"))
		for _, item := range contracts {
			//Check if ancient
			log.Debugln(fmt.Sprintf("⏰  .. Checking %s", item.ID))

			//TODO: check for cancelled contracts
			if ExpiredContractBy(&item, AncientAgeSeconds) {
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
