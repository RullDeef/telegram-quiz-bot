package model

// Модель Квиз
type Quiz struct {
	// ID квиза
	ID int64

	// Тематика квиза
	Topic string

	// Вопросы квиза
	Questions []Question
}
