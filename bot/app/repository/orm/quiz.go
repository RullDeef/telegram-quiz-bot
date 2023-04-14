package orm

import (
	"fmt"
	"time"

	"github.com/RullDeef/telegram-quiz-bot/model"
	"gorm.io/gorm"
)

type questionEntity struct {
	ID     uint `gorm:"primaryKey"`
	QuizID int64
	Text   string
}

type quizEntity struct {
	ID        uint `gorm:"primaryKey"`
	Topic     string
	CreatorID uint
	Creator   userEntity       `gorm:"foreignKey:CreatorID"`
	Questions []questionEntity `gorm:"foreignKey:QuizID"`
	CreatedAt time.Time
}

type ORMQuizRepository struct {
	db *gorm.DB
}

func NewQuizRepo(db *gorm.DB) *ORMQuizRepository {
	return &ORMQuizRepository{
		db: db,
	}
}

func (qzr *ORMQuizRepository) Create(quiz model.Quiz) (model.Quiz, error) {
	entity := quizModelToEntity(quiz)
	err := qzr.db.Create(&entity).Error
	if err != nil {
		return model.Quiz{}, err
	}
	return quizEntityToModel(entity), nil
}

func (qzr *ORMQuizRepository) FindAll() ([]model.Quiz, error) {
	var entities []quizEntity
	err := qzr.db.Preload("Creator").Preload("Questions").Find(&entities).Error
	if err != nil {
		return nil, err
	}
	var quizzes []model.Quiz
	for _, ent := range entities {
		quizzes = append(quizzes, quizEntityToModel(ent))
	}
	return quizzes, nil
}

func (qzr *ORMQuizRepository) FindByID(id int64) (model.Quiz, error) {
	entity := quizEntity{
		ID: uint(id),
	}
	err := qzr.db.Preload("Creator").Preload("Questions").First(&entity).Error
	if err != nil {
		return model.Quiz{}, fmt.Errorf(`quiz with id="%d" not found^ %w`, id, err)
	}
	return quizEntityToModel(entity), nil
}

func (qzr *ORMQuizRepository) FindByTopic(topic string) (model.Quiz, error) {
	var entity quizEntity
	err := qzr.db.First(&entity, "topic = ?", topic).Error
	if err != nil {
		return model.Quiz{}, fmt.Errorf(`quiz with topic="%s" not found: %w`, topic, err)
	}
	return quizEntityToModel(entity), nil
}

func (qzr *ORMQuizRepository) Update(quiz model.Quiz) error {
	entity := quizModelToEntity(quiz)
	err := qzr.db.Preload("Questions").Updates(&entity).Error
	if err != nil {
		return fmt.Errorf(`failed to update quiz with id="%d": %w`, quiz.ID, err)
	}
	return nil
}

func (qzr *ORMQuizRepository) Delete(quiz model.Quiz) error {
	err := qzr.db.Delete(&quizEntity{}, quiz.ID).Error
	if err != nil {
		return fmt.Errorf(`failed to delete quiz with id="%d": %w`, quiz.ID, err)
	}
	return nil
}

func (quizEntity) TableName() string {
	return "quizzes"
}

func (questionEntity) TableName() string {
	return "questions"
}

func quizModelToEntity(quiz model.Quiz) quizEntity {
	ent := quizEntity{
		ID:        uint(quiz.ID),
		Topic:     quiz.Topic,
		CreatorID: uint(quiz.Creator.ID),
		Creator:   userModelToEntity(quiz.Creator),
	}

	for _, question := range quiz.Questions {
		ent.Questions = append(ent.Questions, questionEntity{
			ID:     uint(question.ID),
			QuizID: quiz.ID,
			Text:   question.Text,
		})
	}

	return ent
}

func quizEntityToModel(quiz quizEntity) model.Quiz {
	q := model.Quiz{
		ID:      int64(quiz.ID),
		Topic:   quiz.Topic,
		Creator: userEntityToModel(quiz.Creator),
	}

	for _, question := range quiz.Questions {
		q.Questions = append(q.Questions, model.Question{
			ID:   int64(question.ID),
			Text: question.Text,
		})
	}

	return q
}
