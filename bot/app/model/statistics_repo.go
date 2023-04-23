package model

type StatisticsRepository interface {
	Create(Statistics) error
	FindByUserID(id uint64) (Statistics, error)
	Update(Statistics) error
	Delete(Statistics) error
}
