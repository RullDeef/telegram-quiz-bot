package model

type Interactor interface {
	MessageChan() chan Message
	SendResponse(Response)
}
