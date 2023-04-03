package model

type Response struct {
	// ToMessage *Message // not needed
	Text    string
	Actions []responseAction
}

type responseAction struct {
	ID   int64
	Text string
}

func NewResponse(text string) Response {
	return Response{
		Text:    text,
		Actions: nil,
	}
}

func (r *Response) AddAction(id int64, text string) {
	r.Actions = append(r.Actions, responseAction{
		ID:   id,
		Text: text,
	})
}
