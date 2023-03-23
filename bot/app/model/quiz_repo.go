package model

type QuizRepository interface {
	Create(Quiz)
	FindByID(id uint64)
	FindByTopic(topic string)
	Update(Quiz)
	Delete(Quiz)
}
