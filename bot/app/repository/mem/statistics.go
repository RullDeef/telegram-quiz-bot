package mem_repo

import (
	"errors"
	"fmt"

	"github.com/RullDeef/telegram-quiz-bot/model"
)

type StatisticsRepository struct {
	lastId int
	stats  []model.Statistics
}

func NewStatisticsRepository() *StatisticsRepository {
	return &StatisticsRepository{
		lastId: 1,
		stats:  nil,
	}
}

// Предполагается, что идентификатор статистики будет совпадать с
// идентификатором пользователя, для которого она создается
func (ur *StatisticsRepository) Create(stat model.Statistics) error {
	ur.stats = append(ur.stats, stat)
	return nil
}

func (ur *StatisticsRepository) FindByUserID(id int64) (model.Statistics, error) {
	for _, s := range ur.stats {
		if s.UserID == id {
			return s, nil
		}
	}
	return model.Statistics{}, errors.New("not found")
}

func (ur *StatisticsRepository) Update(stat model.Statistics) error {
	for i, s := range ur.stats {
		if s.UserID == stat.UserID {
			ur.stats[i] = stat
			return nil
		}
	}
	return errors.New("not found")
}

func (ur *StatisticsRepository) Delete(stat model.Statistics) error {
	for i, s := range ur.stats {
		if s.UserID == stat.UserID {
			ur.stats = append(ur.stats[:i], ur.stats[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf(`statistics with UserID="%d" not found`, stat.UserID)
}
