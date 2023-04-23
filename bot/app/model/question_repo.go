package model

// Репозиторий вопросов к квизам
type QuestionRepository interface {

	// Создаёт новый вопрос
	Create(question Question) error

	// Производит поиск вопроса по идентификатору
	FindById(id int64) (Question, error)

	// Обновляет вопрос
	Update(question Question) error

	// Удаляет вопрос по его идентификатору
	Delete(id int64) error
}
