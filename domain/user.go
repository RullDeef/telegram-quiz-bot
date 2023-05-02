package domain

import "context"

type User struct {

	Id		int64 	`json:"id"`
	Username	string 	`json:"name"`
	QuizId		string 	`json:"quiz"`
}

//go:generate mockery --name UserRepository
type UserRepository interface {
	GetById(ctx context.Context, id int64) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
}


