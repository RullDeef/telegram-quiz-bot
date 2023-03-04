package domain

import "context"

type Quiz struct {

	Id int64 `json:"id"`
}

//go:generate mockery --name QuizRepository
type QuizRepository interface {
	GetById(ctx context.Context, id int64) (Quiz, error)
}
