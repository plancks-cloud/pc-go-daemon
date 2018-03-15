package model

//MessageOK represents a successful http call
type MessageOK struct {
	Ok bool `json:"ok"`
}

func (message *MessageOK) String() string {
	if message.Ok == true {
		return "true"
	}
	return "false"
}

//OkMessage returns a message object with the set state
func OkMessage(state bool) MessageOK {
	return MessageOK{Ok: state}
}
