package remote

import (
	"sync"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"time"
)

func ScheduleRemoteSync() {

	go func() {
		for {
			select {
			case <-time.After(model.ScheduledInterval * time.Second):
				waitingDoIt()
			case <-model.DoorBellRemote:
				waitingDoIt()
			}
		}
	}()

}

func waitingDoIt() {

	var wg sync.WaitGroup
	wg.Add(2)

	syncPushAll(wg)
	remoteGC()

	syncPullAll(wg)

	go func() {
		//Ping the community Go routine
		model.DoorBellCommunity <- true
	}()

	wg.Wait()

}
