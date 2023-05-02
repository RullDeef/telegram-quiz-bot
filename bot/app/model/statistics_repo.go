package model

// Репозиторий для получения статистики пользователя
type StatisticsRepository interface {

	// Создает новый объект статистики
	Create(Statistics) error

	// Производит поиск по идентификатору пользователя
	FindByUserID(id int64) (Statistics, error)

	// Обновляет данные статистики
	Update(Statistics) error

	// Удаляет статистику
	Delete(Statistics) error
}
