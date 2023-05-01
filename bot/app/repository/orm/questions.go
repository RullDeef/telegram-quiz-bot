package orm

import (
	"fmt"

	"github.com/RullDeef/telegram-quiz-bot/model"
	"gorm.io/gorm"
)

// БД сущность ответа на вопрос
type answerEntity struct {
	// ID ответа
	ID uint `gorm:"primaryKey"`

	// Текст ответа
	Text string

	// Корректность ответа
	IsCorrect bool

	// ID вопроса, к которому данный ответ относится
	QuestionID uint
}

// БД сущность вопроса
type questionEntity struct {
	// ID вопроса
	ID uint `gorm:"primaryKey"`

	// Текст вопроса
	Text string

	// Тематика, к которой относится данный вопрос
	Topic string

	// Перечень ответов
	Answers []answerEntity `gorm:"foreignKey:QuestionID"`
}

type QuestionsRepository struct {
	db *gorm.DB
}

// Создание сущности Вопросы
func NewQuestionsRepository(
	db *gorm.DB,
) *QuestionsRepository {
	return &QuestionsRepository{
		db: db,
	}
}

// Добавление вопроса в БД
//
//   - question - модель вопроса
//
// Возвращается созданный пользователь с установленным ID, а также ошибка выполнения.
// В случае успешного выполнения возвращается nil.
//
// Возможные ошибки:
//   - в случае ошибки создания возвращается в формате "failed to create question N"
func (qr *QuestionsRepository) Create(question model.Question) (model.Question, error) {
	entity := questionModelToEntity(question)
	err := qr.db.Preload("Answers").Create(&entity).Error
	if err != nil {
		return model.Question{}, fmt.Errorf(`failed to create question "%s": %w`, question.Text, err)
	}
	return questionEntityToModel(entity), nil
}

// Нахождение вопроса по его ID
//
//   - id - ID пользователя
//
// Возвращается модель вопроса, а также ошибка выполнения.
// В случае успешного выполнения возвращается nil.
//
// Возможные ошибки:
//   - в случае ошибки создания возвращается в формате "failed to find question with id=N"
func (qr *QuestionsRepository) FindByID(id int64) (model.Question, error) {
	var entity questionEntity
	err := qr.db.Preload("Answers").First(&entity, id).Error
	if err != nil {
		return model.Question{}, fmt.Errorf(`failed to find question with id=%d: %w`, id, err)
	}
	return questionEntityToModel(entity), err
}

// Нахождение вопроса по тематике квиза
//
//   - topic - тематика квиза
//
// Возвращаются все вопросы по данной тематике, а также ошибка выполнения.
// В случае успешного выполнения возвращается nil.
//
// Возможные ошибки:
//   - в случае ошибки создания вопроса возвращается в формате "failed to find question with topic=N"
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

func (qr *QuestionsRepository) GetAllTopics() ([]string, error) {
	rows, err := qr.db.Raw("select distinct topic from questions").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var topics []string
	for rows.Next() {
		var topic string
		if err := rows.Scan(&topic); err != nil {
			return nil, err
		}
		topics = append(topics, topic)
	}

	return topics, nil
}

// Обновление вопроса в БД
//
//   - q - модель вопроса
//
// В случае успешного выполнения возвращается nil.
// Возможные ошибки:
//   - в случае ошибки обновления вопроса возвращается в формате "failed to update question with id=N"
func (qr *QuestionsRepository) Update(q model.Question) error {
	entity := questionModelToEntity(q)
	err := qr.db.Preload("Answers").Updates(&entity).Error
	if err != nil {
		return fmt.Errorf(`failed to update question with id=%d: %w`, q.ID, err)
	}
	return nil
}

// Удаление вопроса в БД
//
//   - id - ID вопроса
//
// В случае успешного выполнения возвращается nil.
// Возможные ошибки:
//   - в случае ошибки удаления вопроса возвращается в формате "failed to delete question with id=N"
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

// Перевод модельной сущности вопроса в сущность БД
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

// Перевод БД сущности вопроса в модельную
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
