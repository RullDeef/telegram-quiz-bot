package model

import "time"

type StatisticsService interface {
	CreateStatistics(user User) error
	GetStatistics(user User) (Statistics, error)
	SubmitQuizComplete(user User, totalQuizTime time.Duration) error
	SubmitAnswer(user User, isCorrect bool, answerTime time.Duration) error
	ResetStatistics(user User) error
}
