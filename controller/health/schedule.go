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
			case <-time.After(model.ScheduledInterval * time.Second):
				waitingDoIt()
			}
		}
	}()

}

func waitingDoIt() {

	var wg sync.WaitGroup
	wg.Add(1)

	log.Infoln(fmt.Sprintf("⏰  Time to recon services"))
	healthCheckServices(wg)

	wg.Wait()

}
