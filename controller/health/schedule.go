package health

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
)

func ScheduleHealthCheck() {

	go func() {
		for {
			select {
			case <-model.DoorBellHealth:
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

	log.Debugln(fmt.Sprintf("⏰  Time to recon services"))
	healthCheckServices(wg)

	wg.Wait()

}
