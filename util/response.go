package util

import (
	"net/http"
	"encoding/json"
	"fmt"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
)

func RespondWithJsonObject(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)


}

func RespondWithJsonOk(w http.ResponseWriter, v interface{}) {
	RespondWithJsonObject(w, Ok(true))

}

func RespondWithJsonError(w http.ResponseWriter, e error) {
	RespondWithJsonObject(w, OkMessage(false, e.Error()))

}

//MessageOK represents a successful http call
type MessageOK struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message,omitempty"`
}

func (message *MessageOK) String() string {
	if message.Ok == true {
		return "true"
	}
	return fmt.Sprintf("false, %s", message.Message)
}

//Ok returns a message object with the set state
func Ok(state bool) MessageOK {
	return MessageOK{Ok: state}
}

//OkMessage returns a message object with the set state
func OkMessage(state bool, message string) MessageOK {
	return MessageOK{Ok: state, Message: message}
}
