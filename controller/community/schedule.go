package community

import (
	"sync"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"fmt"
	log "github.com/sirupsen/logrus"

	"time"
)

func ScheduleCommunityActivities() {

	go func() {
		for {
			log.Debugln(fmt.Sprintf("🍎️  ScheduleCommunityActivities"))
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
	log.Debugln(fmt.Sprintf("🍎️  ScheduleCommunityActivities: waitingDoIt"))

	var wg sync.WaitGroup
	wg.Add(2)

	considerContracts(wg)
	considerWins(wg)

	wg.Wait()

}
