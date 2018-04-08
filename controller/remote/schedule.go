package remote

import (
	"sync"
	"time"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
)


func ScheduleRemoteSync() {

	go func() {
		for {
			select {
			case <-model.DoorBellRemote:
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

	syncPushAll(wg)
	remoteGC()

	syncPullAll(wg)

	wg.Wait()

}
