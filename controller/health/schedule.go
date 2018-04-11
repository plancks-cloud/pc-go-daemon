package health

import (
	"fmt"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller/db"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	log "github.com/sirupsen/logrus"
	"time"
)

func ScheduleHealthCheck() {

	go func() {
		for {
			log.Debugln(fmt.Sprintf("🍎️  ScheduleHealthCheck"))
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
	log.Infoln(fmt.Sprintf("🍎️  ScheduleHealthCheck: tick"))
	db.ReconServices()

}
