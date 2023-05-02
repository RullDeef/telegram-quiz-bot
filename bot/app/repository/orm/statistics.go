package orm

import (
	"fmt"

	"github.com/RullDeef/telegram-quiz-bot/model"
	"gorm.io/gorm"
)

// БД сущность статистики
type statEntity struct {
	// Идентификатор полльзователя
	UserID uint `gorm:"primaryKey"`

	// Количество пройденных квизов
	QuizzesCompleted int

	// Среднее время прохождения квиза в секундах
	MeanQuizCompleteTime float64

	// Среднее время ответа на вопрос в секундах
	MeanQuestionReplyTime float64

	// Общее количество ответов
	TotalReplies int

	// Количество верных ответов
	CorrectReplies int

	// Процент верных  ответов (от 0 до 100)
	CorrectRepliesPercent int
}

type ORMStatsRepository struct {
	db *gorm.DB
}

func NewStatisticsRepo(db *gorm.DB) *ORMStatsRepository {
	return &ORMStatsRepository{
		db: db,
	}
}

// Создание статистики пользователя
//
//   - stat - модель статистики
//
// Возможные ошибки:
//   - Объект статистики для данного пользователя уже существует
func (sr *ORMStatsRepository) Create(stat model.Statistics) error {
	entity := sr.statModelToEntity(stat)
	if err := sr.db.Create(&entity).Error; err != nil {
		return err
	}
	return nil
}

// Поиск статистики пользователя по его идентификатору
//
//   - id - идентификатор пользователя
//
// Возвращается объект статистики, соответствующий данному пользователю.
// Возможные ошибки:
//   - Отсутсвие объекта статистики в базе
func (sr *ORMStatsRepository) FindByUserID(id int64) (model.Statistics, error) {
	var entity statEntity
	err := sr.db.First(&entity, "user_id = ?", uint(id)).Error
	if err != nil {
		return model.Statistics{}, fmt.Errorf(`statistics for user with id="%d" not found: %w`, id, err)
	}
	return sr.statEntityToModel(entity), nil
}

// Обновление объекта статистики.
// Обновление осуществляется по идентификатору пользователя
//
//   - stat - обновлённый объект статистики
//
// Возвращает nil в случае успеха. Возможные ошибки:
//   - Статистика с идентификатором stat.UserID не существует в базе
func (sr *ORMStatsRepository) Update(stat model.Statistics) error {
	entity := sr.statModelToEntity(stat)
	if err := sr.db.Updates(&entity).Error; err != nil {
		return fmt.Errorf(`failed to update stat for user with id="%d": %w`, entity.UserID, err)
	}
	return nil
}

// Удаление объекта статистики.
//
//   - stat - удаляемый объект статистики
//
// Возращает nil в случае успеха. Возможные ошибки:
//   - Внутренние ошибки базы данных
func (sr *ORMStatsRepository) Delete(stat model.Statistics) error {
	entity := sr.statModelToEntity(stat)
	if err := sr.db.Delete(&entity).Error; err != nil {
		return fmt.Errorf(`failed to delete stat with user id="%d": %w`, entity.UserID, err)
	}
	return nil
}

func (statEntity) TableName() string {
	return "statistics"
}

// Перевод модельной сущности статистики в сущность БД
func (sr *ORMStatsRepository) statModelToEntity(stat model.Statistics) statEntity {
	return statEntity{
		UserID:                uint(stat.UserID),
		QuizzesCompleted:      int(stat.QuizzesCompleted),
		MeanQuizCompleteTime:  stat.MeanQuizCompleteTime,
		MeanQuestionReplyTime: stat.MeanQuestionReplyTime,
		TotalReplies:          int(stat.TotalReplies),
		CorrectReplies:        int(stat.CorrectReplies),
		CorrectRepliesPercent: int(stat.CorrectRepliesPercent * 100),
	}
}

// Перевод сущности БД статистики в модельную сущность
func (sr *ORMStatsRepository) statEntityToModel(stat statEntity) model.Statistics {
	return model.Statistics{
		UserID:                int64(stat.UserID),
		QuizzesCompleted:      uint(stat.QuizzesCompleted),
		MeanQuizCompleteTime:  stat.MeanQuizCompleteTime,
		MeanQuestionReplyTime: stat.MeanQuestionReplyTime,
		TotalReplies:          uint(stat.TotalReplies),
		CorrectReplies:        uint(stat.CorrectReplies),
		CorrectRepliesPercent: float64(stat.CorrectRepliesPercent) / 100,
	}
}
