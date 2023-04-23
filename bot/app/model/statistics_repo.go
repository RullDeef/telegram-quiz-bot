package model

type StatisticsRepository interface {
	Create(Statistics) error
	FindByUserID(id int64) (Statistics, error)
	Update(Statistics) error
	Delete(Statistics) error
}
