package orm

import (
	"github.com/RullDeef/telegram-quiz-bot/model"
	"gorm.io/gorm"
)

//New version
type AnswerRepositoryStruct struct {
	Db     *gorm.DB
	LastId int64
}

func (ar *AnswerRepositoryStruct) FindByAnswerId(id int64) (model.Answer, error) {
	var answer model.Answer
	result := ar.Db.Find(&answer, id)

	return answer, result.Error
}

func (ar *AnswerRepositoryStruct) Create(answer model.Answer) error {
	return ar.Db.Table("answers").Create(answer).Error
}

func (ar *AnswerRepositoryStruct) Update(answer model.Answer) error {
	return ar.Db.Table("answers").Where("id = ?", answer.ID).Save(&answer).Error
}

func (ar *AnswerRepositoryStruct) Delete(id int64) error {
	return ar.Db.Table("answers").Delete(&model.Answer{}, id).Error
}
