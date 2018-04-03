package util

import (
	"net/http"
	"encoding/json"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
)

func RespondWithJsonObject(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)


}

func RespondWithJsonOk(w http.ResponseWriter, v interface{}) {
	RespondWithJsonObject(w, model.Ok(true))

}

func RespondWithJsonError(w http.ResponseWriter, e error) {
	RespondWithJsonObject(w, model.OkMessage(false, e.Error()))

}