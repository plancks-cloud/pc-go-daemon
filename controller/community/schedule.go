package community

import (
	"sync"
	"time"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
)

func ScheduleCommunityActivities() {

	go func() {
		for {
			select {
			case <-model.DoorBellCommunity:
				waitingDoIt()
			case <-time.After(30 * time.Second):
				waitingDoIt()
			}
		}
	}()

}

func waitingDoIt() {

	var wg sync.WaitGroup
	wg.Add(2)

	considerContracts(wg)
	considerWins(wg)

	wg.Wait()

}
