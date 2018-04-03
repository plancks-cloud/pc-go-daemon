package util

import (
	"net/http"
	"encoding/json"
)

func RespondWithJson(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)


}