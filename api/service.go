package api

import (
	"encoding/json"
	"net/http"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller/db"
)

//GetServiceStateResult returns a service with its health
func GetServiceStateResult(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db.GetServiceStateResult())
}
