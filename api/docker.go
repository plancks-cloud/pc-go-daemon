package api

import (
	"fmt"
	"html"
	"net/http"
)

//DockerListServices lists all docker services running
func DockerListServices(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello creator, %q", html.EscapeString(r.URL.Path))
}

//ForceSync forces a sync
func ForceSync(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello creator, %q", html.EscapeString(r.URL.Path))
}
