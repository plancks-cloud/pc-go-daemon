package health

import (
	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller/db"
	"sync"
)

//ReconServices sets up scheduler to recon docker services running
func healthCheckServices(wg sync.WaitGroup) {
	go func(wg sync.WaitGroup) {
		db.ReconServices()
		wg.Done()
	}(wg)
}
