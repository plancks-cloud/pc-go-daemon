package health

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

var Wakey = make(chan bool)

func ScheduleHealthCheck() {

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

	log.Debugln(fmt.Sprintf("⏰  Time to recon services"))
	healthCheckServices(wg)

	wg.Wait()

}
