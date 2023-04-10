package model

type Quiz struct {
	ID        uint64
	Topic     string
	Questions []Question
}
