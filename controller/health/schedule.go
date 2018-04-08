package health

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"time"
)

func ScheduleHealthCheck() {

	go func() {
		for {
			log.Debugln(fmt.Sprintf("❄️  ScheduleHealthCheck"))
			select {
			case <-time.After(model.ScheduledInterval * time.Second):
				waitingDoIt()
			case <-model.DoorBellHealth:
				waitingDoIt()
			}
		}
	}()

}

func waitingDoIt() {
	log.Debugln(fmt.Sprintf("❄️  ScheduleHealthCheck: waitingDoIt"))
	healthCheckServices()

}
