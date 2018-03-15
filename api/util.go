package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller"
)

//Ping perform a health check
func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(controller.HealthCheck())
}

//CorsHandler does cors check
func CorsHandler(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			fmt.Fprintf(w, "Options")
		} else {
			f.ServeHTTP(w, r)
		}
	}
}

//ForceSync forces a sync
func ForceSync(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(controller.ForceSync())
}
