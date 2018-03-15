package controller

import (
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
)

//DockerListServices lists all Docker services
func DockerListServices() []model.Service {
	var services []model.Service
	services = append(services, model.Service{})
	return services
}
