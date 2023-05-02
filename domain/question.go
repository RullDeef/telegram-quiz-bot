package domain

import "context"

type Question struct {

	Id	int64 	`json:"id"`
	QuizId	int64	`json:"quiz"`
}

//go:generate mockery --name QuestionRepository
type QuestionRepository interface {
	GetById(ctx context.Context, id int64) (Question, error)
}
