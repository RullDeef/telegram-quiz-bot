package orm

import (
	"testing"

	"github.com/RullDeef/telegram-quiz-bot/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestStatRepositoryInterface(t *testing.T) {
	var _ model.StatisticsRepository = &ORMStatsRepository{}
}

func TestStatRepository(t *testing.T) {
	dsn := "host=testdb user=postgres password=root port=5432 dbname=quizdb"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		t.Errorf("Get connection to db = %s; want nil", err)
		t.FailNow()
	}

	user_repo := NewUserRepo(db)
	stat_repo := NewStatisticsRepo(db)

	user := model.User{
		Nickname:   "mega user",
		TelegramID: "mega_user",
		Role:       "USER",
	}

	t.Run("create user", func(t *testing.T) {
		user, err = user_repo.Create(user)
		if err != nil {
			t.Errorf("failed to insert creator user")
			t.FailNow()
		}
	})

	t.Run("Create statistics", func(t *testing.T) {
		stat := model.Statistics{
			UserID: user.ID,
		}
		err = stat_repo.Create(stat)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("statistics created: %v", stat)
		}
	})

	t.Run("Find by user id", func(t *testing.T) {
		_, err = stat_repo.FindByUserID(user.ID)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Find by user id not found", func(t *testing.T) {
		_, err = stat_repo.FindByUserID(user.ID + 200)
		if err == nil {
			t.Errorf("found statistics; want 0")
		}
	})

	t.Run("Update", func(t *testing.T) {
		stat := model.Statistics{
			UserID:                user.ID,
			QuizzesCompleted:      20,
			MeanQuizCompleteTime:  631.3,
			MeanQuestionReplyTime: 2.4,
			CorrectReplies:        50,
			CorrectRepliesPercent: 42,
		}
		if err = stat_repo.Update(stat); err != nil {
			t.Error(err)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		err = stat_repo.Delete(model.Statistics{
			UserID: user.ID,
		})
		if err != nil {
			t.Error(err)
		}
	})
}
