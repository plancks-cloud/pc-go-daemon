package community

import (
	"sync"
	"time"
)

var Wakey = make(chan bool)

func ScheduleCommunityActivities() {

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

	considerContracts(wg)
	considerWins(wg)

	wg.Wait()

}
