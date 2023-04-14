package orm

import (
	"github.com/RullDeef/telegram-quiz-bot/model"
	"gorm.io/gorm"
)

type QuestionsRepositoryStruct struct {
	Db     *gorm.DB
	LastId int
}

func (qr *QuestionsRepositoryStruct) Create(question model.Question) error {
	return qr.Db.Table("questions").Create(question).Error
}

func (qr *QuestionsRepositoryStruct) FindById(id int64) (model.Question, error) {
	var question model.Question
	result := qr.Db.Table("questions").Find(&question, id)

	return question, result.Error
}

func (qr *QuestionsRepositoryStruct) Update(question model.Question) error {
	return qr.Db.Table("questions").Where("id = ?", question.ID).Save(&question).Error
}

func (qr *QuestionsRepositoryStruct) Delete(id int64) error {
	return qr.Db.Table("questions").Delete(&model.Question{}, id).Error
}
