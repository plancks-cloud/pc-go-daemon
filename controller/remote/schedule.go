package remote

import (
	"sync"
	"time"
)

var Wakey = make(chan bool)

func ScheduleRemoteSync() {

	go func() {

		select {
		case <-Wakey:
			waitingDoIt()
		case <-time.After(30 * time.Second):
			waitingDoIt()
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
