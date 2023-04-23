package model

type Quiz struct {
	ID        int64
	Topic     string
	Questions []Question
}
