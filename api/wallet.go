package api

import (
	"fmt"
	"html"
	"net/http"
)

//CreateWallet creates a new wallet
func CreateWallet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello creator, %q", html.EscapeString(r.URL.Path))
}

//SetCurrentWallet sets the currently used wallet
func SetCurrentWallet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello creator, %q", html.EscapeString(r.URL.Path))
}
