package orm

import (
	"github.com/RullDeef/telegram-quiz-bot/model"
	"gorm.io/gorm"
)

//New version
type AnswerRepositoryStruct struct {
	Db      *gorm.DB
	LastId  int64
	Quizzes []model.Quiz
}

func (qr *AnswerRepositoryStruct) Get(id int64) (model.Answer, error) {
	var answer model.Answer
	result := qr.Db.Find(&answer, id)

	return answer, result.Error
}

func (qr *AnswerRepositoryStruct) Create(answer model.Answer) error {
	return qr.Db.Table("answers").Create(answer).Error
}

func (qr *AnswerRepositoryStruct) Update(answer model.Answer) error {
	return qr.Db.Table("answers").Where("id = ?", answer.ID).Save(&answer).Error
}

func (qr *AnswerRepositoryStruct) Delete(id int64) error {
	return qr.Db.Table("answers").Delete(&model.Answer{}, id).Error
}
