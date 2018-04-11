package community

import (
	"fmt"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	log "github.com/sirupsen/logrus"

	"time"
)

//ScheduleCommunityActivities checks for bids and wins
func ScheduleCommunityActivities() {

	go func() {
		for {
			log.Debugln(fmt.Sprintf("🏟️️  ScheduleCommunityActivities"))
			select {
			case <-time.After(model.ScheduledInterval * time.Second):
				waitingDoIt()
			case <-model.DoorBellCommunity:
				waitingDoIt()
			}
		}
	}()

}

func waitingDoIt() {
	log.Infoln(fmt.Sprintf("🏟️️  ScheduleCommunityActivities: tick"))

	considerContracts()
	considerWins()

	//Ping the healing GR
	go func() {
		model.DoorBellHealth <- true
	}()

}
