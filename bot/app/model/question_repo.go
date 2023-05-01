package model

// Репозиторий вопросов к квизам
type QuestionRepository interface {

	// Создаёт новый вопрос
	Create(question Question) (Question, error)

	// Производит поиск вопроса по идентификатору
	FindByID(id int64) (Question, error)

	// Производит поиск вопросов по тематике
	FindByTopic(topic string) ([]Question, error)

	// Возвращает список уникальных тематик (("+ собранных по всем вопросам" надо написать короче))
	GetAllTopics() ([]string, error)

	// Обновляет вопрос
	Update(question Question) error

	// Удаляет вопрос по его идентификатору
	Delete(id int64) error
}
