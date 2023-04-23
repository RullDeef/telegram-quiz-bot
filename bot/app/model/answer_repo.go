package model

// Не используется
type AnswerRepository interface {
	Create(answer Answer) error
	FindByAnswerId(id int64) (Answer, error)
	FindByQuestionId(id int64) ([]Answer, error)
	Update(answer Answer) error
	Delete(id int64) error
}
