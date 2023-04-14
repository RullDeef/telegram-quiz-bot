package model

type QuestionRepository interface {
	Create(question Question) error
	FindById(id int64) (Question, error)
	Update(question Question) error
	Delete(id int64) error
}
