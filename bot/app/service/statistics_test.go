package service

import (
	"testing"
	"time"

	"github.com/RullDeef/telegram-quiz-bot/model"
	mem_repo "github.com/RullDeef/telegram-quiz-bot/repository/mem"
	log "github.com/sirupsen/logrus"
)

func TestStatisticsService(t *testing.T) {
	userRepo := mem_repo.NewUserRepository()

	user := model.User{
		Nickname:   "petya",
		TelegramID: "petyatg",
		Role:       model.UserRoleUser,
	}
	user, err := userRepo.Create(user)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	correctAnswer := true
	wrongAnswer := false

	statRepo := mem_repo.NewStatisticsRepository()
	logger := log.New()

	statService := NewStatisticsService(userRepo, statRepo, logger)

	t.Run("Create statistics", func(t *testing.T) {
		err = statService.CreateStatistics(user)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Submit wrong answer", func(t *testing.T) {
		err = statService.SubmitAnswer(user, wrongAnswer, time.Minute)
		if err != nil {
			t.Error(err)
		}

		stat, err := statRepo.FindByUserID(user.ID)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		if stat.CorrectReplies != 0 {
			t.Errorf(`expected CorrectReplies=0, got %d`, stat.CorrectReplies)
		}

		if stat.MeanQuestionReplyTime != 60.0 {
			t.Errorf(`expected MeanQuestionReplyTime=60.0, got %.1f`, stat.MeanQuestionReplyTime)
		}
	})

	t.Run("Submit correct answer", func(t *testing.T) {
		err = statService.SubmitAnswer(user, correctAnswer, time.Second)
		if err != nil {
			t.Error(err)
		}

		stat, err := statRepo.FindByUserID(user.ID)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		if stat.CorrectReplies != 1 {
			t.Errorf(`expected CorrectReplies=1, got %d`, stat.CorrectReplies)
		}

		if stat.MeanQuestionReplyTime != 30.5 {
			t.Errorf(`expected MeanQuestionReplyTime=30.5, got %.1f`, stat.MeanQuestionReplyTime)
		}
	})

	t.Run("Submit quiz completed", func(t *testing.T) {
		err = statService.SubmitQuizComplete(user, 3*time.Minute)
		if err != nil {
			t.Error(err)
		}

		stat, err := statRepo.FindByUserID(user.ID)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		if stat.QuizzesCompleted != 1 {
			t.Errorf(`expected QuizzesCompleted=1, got %d`, stat.QuizzesCompleted)
		}

		if stat.MeanQuizCompleteTime != 180.0 {
			t.Errorf(`expected MeanQuizCompleteTime=180.0, got %.1f`, stat.MeanQuizCompleteTime)
		}
	})
}
