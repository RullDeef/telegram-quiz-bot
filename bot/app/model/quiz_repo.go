package model

// Интерфейс для получения квизов
type QuizRepository interface {

	// Создает новый квиз, возравщает ошибку при неудаче
	Create(Quiz) (Quiz, error)

	// Возвращает все квизы.
	FindAll() ([]Quiz, error)

	// Производит поиск квиза по его идентификатору
	FindByID(id int64) (Quiz, error)

	// Производит поиск квиза по тематике
	FindByTopic(topic string) (Quiz, error)

	// Обновляет данные о квизе, возвращает ошибку при неудаче
	Update(Quiz) error

	// Удаляет квиз по идентификатору, возвращает ошибку при неудаче
	Delete(Quiz) error
}
