package model

type Question struct {
	ID    int64
	Text  string
	Topic string
	Answers []Answer
}
