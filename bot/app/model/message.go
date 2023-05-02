package model

import "time"

const INVALID_BUTTON_ACTION_ID = -1

type Message struct {
	Sender         *User
	ChatID         int64
	Text           string
	ReceiveTime    time.Time
	IsPrivate      bool
	IsButtonAction bool
	ButtonActionID int64
}

func (m *Message) ActionID() int64 {
	if m.IsButtonAction {
		return m.ButtonActionID
	} else {
		return INVALID_BUTTON_ACTION_ID
	}
}
