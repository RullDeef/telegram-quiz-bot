package model

// Модельная сущность Ответ
type Answer struct {
	// ID ответа
	ID uint64

	// Текст ответа
	Text string

	// Правильность ответа
	IsСorrect bool
}
