package community

import (
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"fmt"
	log "github.com/sirupsen/logrus"

	"time"
)

func ScheduleCommunityActivities() {

	go func() {
		for {
			log.Infoln(fmt.Sprintf("🍎️  ScheduleCommunityActivities"))
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
	log.Infoln(fmt.Sprintf("🍎️  ScheduleCommunityActivities: waitingDoIt"))

	considerContracts()
	considerWins()

}
