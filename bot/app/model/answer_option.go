package model

type Answer struct {
	ID          uint64
	Question_ID uint64
	Text        string
	Is_correct  bool
}
