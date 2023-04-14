package orm

import (
	"errors"

	"github.com/RullDeef/telegram-quiz-bot/model"
	"gorm.io/gorm"
)

type QuizRepositoryStruct struct {
	Db     *gorm.DB
	LastId int
}

func (qzr *QuizRepositoryStruct) Create(quiz model.QuizNew) error {
	return qzr.Db.Table("quizzes").Create(&quiz).Error
}

func (qzr *QuizRepositoryStruct) FindAll() ([]model.QuizNew, error) {
	var all_quizzes []model.QuizNew

	result := qzr.Db.Table("quizzes").Find(&all_quizzes)

	return all_quizzes, result.Error
}

func (qzr *QuizRepositoryStruct) FindByID(id int64) (model.QuizNew, error) {
	var quiz model.QuizNew

	result := qzr.Db.Table("quizzes").Find(&quiz, id)
	err := result.Error

	if result.RowsAffected == 0 {
		err = errors.New("null")
	}

	return quiz, err
}

func (qzr *QuizRepositoryStruct) FindByTopic(topic string) (model.QuizNew, error) {
	var topic_quizz model.QuizNew

	result := qzr.Db.Table("quizzes").Where("topic = ?", topic).Find(&topic_quizz)
	err := result.Error

	if result.RowsAffected == 0 {
		err = errors.New("null")
	}

	return topic_quizz, err
}

func (qzr *QuizRepositoryStruct) Update(quiz model.QuizNew) error {
	result := qzr.Db.Table("quizzes").Where("id = ?", quiz.ID).Updates(&quiz)
	err := result.Error

	if result.RowsAffected == 0 {
		err = errors.New("null")
	}

	return err
}

func (qzr *QuizRepositoryStruct) Delete(id int64) error {
	result := qzr.Db.Table("quizzes").Delete(&model.QuizNew{}, id)

	err := result.Error

	if result.RowsAffected == 0 {
		err = errors.New("null")
	}

	return err
}
