package model

type Answer struct {
	ID        uint64
	Q_ID      uint64
	Text      string
	IsCorrect bool
}
