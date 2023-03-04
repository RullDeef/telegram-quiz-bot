package domain

import "context"

type Stats struct {

	Id	int64 	`json:"id"`
	UserId	int64	`json:"user"`
}

//go:generate mockery --name StatsRepository
type StatsRepository interface {
	GetById(ctx context.Context, id int64) (Stats, error)
}
