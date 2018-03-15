package api

import (
	"encoding/json"
	"net/http"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller"
)

//DockerListServices lists all docker services running
func DockerListServices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(controller.DockerListServices())
}
