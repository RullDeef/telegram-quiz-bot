package orm

import (
	"errors"

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
	result := ar.Db.Table("answers").Find(&answer, id)
	err := result.Error

	if result.RowsAffected == 0 {
		err = errors.New("null")
	}

	return answer, err
}

func (ar *AnswerRepositoryStruct) FindByQuestionId(id int64) ([]model.Answer, error) {
	var answer []model.Answer
	result := ar.Db.Table("answers").Where("question_id = ?", id).Find(&answer)
	err := result.Error

	if result.RowsAffected == 0 {
		err = errors.New("null")
	}

	return answer, err
}

func (ar *AnswerRepositoryStruct) Create(answer model.Answer) error {
	return ar.Db.Table("answers").Create(&answer).Error
}

func (ar *AnswerRepositoryStruct) Update(answer model.Answer) error {
	result := ar.Db.Table("answers").Where("id = ?", answer.ID).Updates(&answer)

	err := result.Error

	if result.RowsAffected == 0 {
		err = errors.New("null")
	}

	return err
}

func (ar *AnswerRepositoryStruct) Delete(id int64) error {
	result := ar.Db.Table("answers").Delete(&model.Answer{}, id)

	err := result.Error

	if result.RowsAffected == 0 {
		err = errors.New("null")
	}

	return err
}
