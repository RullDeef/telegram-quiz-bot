package orm

import (
	"fmt"

	"github.com/RullDeef/telegram-quiz-bot/model"
	"gorm.io/gorm"
)

type statEntity struct {
	UserID                uint `gorm:"primaryKey"`
	QuizzesCompleted      int
	MeanQuizCompleteTime  float64
	MeanQuestionReplyTime float64
	CorrectReplies        int
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

func (sr *ORMStatsRepository) Create(stat model.Statistics) error {
	entity := sr.statModelToEntity(stat)
	if err := sr.db.Create(&entity).Error; err != nil {
		return err
	}
	return nil
}

func (sr *ORMStatsRepository) FindByUserID(id int64) (model.Statistics, error) {
	var entity statEntity
	err := sr.db.First(&entity, "user_id = ?", uint(id)).Error
	if err != nil {
		return model.Statistics{}, fmt.Errorf(`statistics for user with id="%d" not found: %w`, id, err)
	}
	return sr.statEntityToModel(entity), nil
}

func (sr *ORMStatsRepository) Update(stat model.Statistics) error {
	entity := sr.statModelToEntity(stat)
	if err := sr.db.Updates(&entity).Error; err != nil {
		return fmt.Errorf(`failed to update stat for user with id="%d": %w`, entity.UserID, err)
	}
	return nil
}

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

func (sr *ORMStatsRepository) statModelToEntity(stat model.Statistics) statEntity {
	return statEntity{
		UserID:                uint(stat.UserID),
		QuizzesCompleted:      int(stat.QuizzesCompleted),
		MeanQuizCompleteTime:  stat.MeanQuizCompleteTime,
		MeanQuestionReplyTime: stat.MeanQuestionReplyTime,
		CorrectReplies:        int(stat.CorrectReplies),
		CorrectRepliesPercent: int(stat.CorrectRepliesPercent * 100),
	}
}

func (sr *ORMStatsRepository) statEntityToModel(stat statEntity) model.Statistics {
	return model.Statistics{
		UserID:                int64(stat.UserID),
		QuizzesCompleted:      uint(stat.QuizzesCompleted),
		MeanQuizCompleteTime:  stat.MeanQuizCompleteTime,
		MeanQuestionReplyTime: stat.MeanQuestionReplyTime,
		CorrectReplies:        uint(stat.CorrectReplies),
		CorrectRepliesPercent: float64(stat.CorrectRepliesPercent) / 100,
	}
}
