package model

type StatisticsRepository interface {
	Create(User)
	FindByID(id uint64)
	FindByUserID(id uint64)
	Update(Statistics)
	Delete(Statistics)
}
