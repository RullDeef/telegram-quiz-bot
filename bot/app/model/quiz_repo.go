package model

type QuizRepository interface {
	Create(Quiz) Quiz
	FindAll() []Quiz
	FindByID(id int64) (Quiz, error)
	FindByTopic(topic string) (Quiz, error)
	Update(Quiz) error
	Delete(Quiz)
}

type QuizRepositoryNew interface {
	Create(Quiz) error
	FindAll() []Quiz
	FindByID(id int64) (Quiz, error)
	FindByTopic(topic string) (Quiz, error)
	Update(Quiz) error
	Delete(Quiz)
}
