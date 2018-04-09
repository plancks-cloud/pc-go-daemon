package health

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"time"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller/db"
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
	log.Infoln(fmt.Sprintf("❄️  ScheduleHealthCheck: tick"))
	db.ReconServices()

}
