package orm

import (
	"github.com/RullDeef/telegram-quiz-bot/model"
	"gorm.io/gorm"
)

type QuizRepositoryStruct struct {
	Db     *gorm.DB
	LastId int
}

func (qzr *QuizRepositoryStruct) Create(quiz model.Quiz) error {
	return qzr.Db.Table("quizzes").Create(&quiz).Error
}

func (qzr *QuizRepositoryStruct) FindAll() ([]model.Quiz, error) {
	var all_quizzes []model.Quiz

	result := qzr.Db.Table("quizzes").Find(&all_quizzes)

	return all_quizzes, result.Error
}

func (qzr *QuizRepositoryStruct) FindByID(id int64) (model.Quiz, error) {
	var quiz model.Quiz

	result := qzr.Db.Table("quizzes").Find(&quiz, id)

	return quiz, result.Error
}

func (qzr *QuizRepositoryStruct) FindByTopic(topic string) (model.Quiz, error) {
	var topic_quizz model.Quiz

	result := qzr.Db.Table("quizzes").Where("topic = ?", topic).Find(&topic_quizz)

	return topic_quizz, result.Error
}

func (qzr *QuizRepositoryStruct) Update(quiz model.Quiz) error {
	return qzr.Db.Table("quizzes").Where("id = ?", quiz.ID).Save(&quiz).Error
}

func (qzr *QuizRepositoryStruct) Delete(id int64) error {
	return qzr.Db.Table("quizzes").Delete(&model.Quiz{}, id).Error
}
