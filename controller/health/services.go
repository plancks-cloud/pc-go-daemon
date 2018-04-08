package health

import (
	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller/db"
)

//ReconServices sets up scheduler to recon docker services running
func healthCheckServices() {
	db.ReconServices()
}
