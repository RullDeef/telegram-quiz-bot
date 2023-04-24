package orm

import (
	"fmt"

	"github.com/RullDeef/telegram-quiz-bot/model"
	"gorm.io/gorm"
)

type answerEntity struct {
	ID         uint `gorm:"primaryKey"`
	Text       string
	IsCorrect  bool
	QuestionID uint
}

type questionEntity struct {
	ID      uint `gorm:"primaryKey"`
	Text    string
	Topic   string
	Answers []answerEntity `gorm:"foreignKey:QuestionID"`
}

type QuestionsRepository struct {
	db *gorm.DB
}

func NewQuestionsRepository(
	db *gorm.DB,
) *QuestionsRepository {
	return &QuestionsRepository{
		db: db,
	}
}

func (qr *QuestionsRepository) Create(question model.Question) (model.Question, error) {
	entity := questionModelToEntity(question)
	err := qr.db.Preload("Answers").Create(&entity).Error
	if err != nil {
		return model.Question{}, fmt.Errorf(`failed to create question "%s": %w`, question.Text, err)
	}
	return questionEntityToModel(entity), nil
}

func (qr *QuestionsRepository) FindByID(id int64) (model.Question, error) {
	var entity questionEntity
	err := qr.db.Preload("Answers").First(&entity, id).Error
	if err != nil {
		return model.Question{}, fmt.Errorf(`failed to find question with id=%d: %w`, id, err)
	}
	return questionEntityToModel(entity), err
}

func (qr *QuestionsRepository) FindByTopic(topic string) ([]model.Question, error) {
	var entities []questionEntity
	err := qr.db.Preload("Answers").Find(&entities, "topic = ?", topic).Error
	if err != nil {
		return nil, fmt.Errorf(`failed to find question with topic="%s": %w`, topic, err)
	}
	var questions []model.Question
	for _, entity := range entities {
		questions = append(questions, questionEntityToModel(entity))
	}
	return questions, nil
}

func (qr *QuestionsRepository) Update(q model.Question) error {
	entity := questionModelToEntity(q)
	err := qr.db.Preload("Answers").Updates(&entity).Error
	if err != nil {
		return fmt.Errorf(`failed to update question with id=%d: %w`, q.ID, err)
	}
	return nil
}

func (qr *QuestionsRepository) Delete(id int64) error {
	err := qr.db.Preload("Answers").Delete(&questionEntity{}, id).Error
	if err != nil {
		return fmt.Errorf(`failed to delete question with id=%d: %w`, id, err)
	}
	return nil
}

func (questionEntity) TableName() string {
	return "questions"
}

func (answerEntity) TableName() string {
	return "answers"
}

func questionModelToEntity(q model.Question) questionEntity {
	var answers []answerEntity
	for _, ans := range q.Answers {
		answers = append(answers, answerEntity{
			ID:         uint(ans.ID),
			Text:       ans.Text,
			IsCorrect:  ans.IsСorrect,
			QuestionID: uint(q.ID),
		})
	}

	return questionEntity{
		ID:      uint(q.ID),
		Text:    q.Text,
		Topic:   q.Topic,
		Answers: answers,
	}
}

func questionEntityToModel(entity questionEntity) model.Question {
	var answers []model.Answer
	for _, ans := range entity.Answers {
		answers = append(answers, model.Answer{
			ID:        uint64(ans.ID),
			Text:      ans.Text,
			IsСorrect: ans.IsCorrect,
		})
	}

	return model.Question{
		ID:      int64(entity.ID),
		Text:    entity.Text,
		Topic:   entity.Topic,
		Answers: answers,
	}
}
