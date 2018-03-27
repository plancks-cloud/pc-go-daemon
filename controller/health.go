package controller

import "git.amabanana.com/plancks-cloud/pc-go-daemon/model"

//HealthCheck performs a health check and returns the state of the system
func HealthCheck() model.MessageOK {
	return model.Ok(true)
}
