package model

// Модельная сущность Вопроса
type Question struct {
	// ID вопроса
	ID int64

	// Текст вопроса
	Text string

	// Тематика вопроса
	Topic string

	// Ответы на вопрос.
	// Должна содержать максимум 1 ответ с IsCorrect = true.
	Answers []Answer
}
