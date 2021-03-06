package api

import (
	"encoding/json"
	"net/http"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller/db"
)

//DockerListServices lists all docker services in DB
func DockerListServices(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db.DockerListServices())
}

//DockerListRunningServices lists all docker services running
func DockerListRunningServices(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db.DockerListRunningServices())
}
