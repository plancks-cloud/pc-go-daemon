package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//RespondWithJsonObject is a generic json returning utility method
func RespondWithJsonObject(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)

}

//RespondWithJsonOk writes ok back to the http response
func RespondWithJsonOk(w http.ResponseWriter, v interface{}) {
	RespondWithJsonObject(w, Ok(true))

}

//RespondWithJsonError sends back an error
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
