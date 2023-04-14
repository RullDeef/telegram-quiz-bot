package model

type QuizRepository interface {
	Create(Quiz) (Quiz, error)
	FindAll() ([]Quiz, error)
	FindByID(id int64) (Quiz, error)
	FindByTopic(topic string) (Quiz, error)
	Update(Quiz) error
	Delete(Quiz) error
}
