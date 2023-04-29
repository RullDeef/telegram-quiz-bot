package service

import (
	"time"

	"github.com/RullDeef/telegram-quiz-bot/model"
	log "github.com/sirupsen/logrus"
)

type StatisticsService struct {
	userRepo model.UserRepository
	statRepo model.StatisticsRepository
	logger   *log.Logger
}

func NewStatisticsService(
	userRepo model.UserRepository,
	statRepo model.StatisticsRepository,
	logger *log.Logger,
) *StatisticsService {
	return &StatisticsService{
		userRepo: userRepo,
		statRepo: statRepo,
		logger:   logger,
	}
}

// Создает объект статистики через репозиторий
func (ss *StatisticsService) CreateStatistics(user model.User) error {
	ss.logger.WithField("user", user).Info("CreateStatistics")
	stat := model.Statistics{UserID: user.ID}
	err := ss.statRepo.Create(stat)
	if err != nil {
		ss.logger.Error(err)
	}
	return err
}

// Получает статистику пользователя из репозитория
func (ss *StatisticsService) GetStatistics(user model.User) (model.Statistics, error) {
	ss.logger.WithField("user", user).Info("GetStatistics")
	stat, err := ss.statRepo.FindByUserID(user.ID)
	if err != nil {
		ss.logger.Error(err)
	}
	return stat, err
}

// Увеличивает счетчик пройденных квизов для данного пользователя
//
// Также обновляет среднее время прохождения квиза.
// Данный метод необходимо вызывать один раз для каждого пользователя после завершения квиза.
//
// Возвращает ошибку в следующих случаях:
//   - объект статистики для данного пользователя не найден
//   - ошибка при записи данных
func (ss *StatisticsService) SubmitQuizComplete(user model.User, totalQuizTime time.Duration) error {
	ss.logger.
		WithFields(log.Fields{"user": user, "totalQuizTime": totalQuizTime}).
		Info("SubmitQuizComplete")
	stat, err := ss.statRepo.FindByUserID(user.ID)
	if err != nil {
		ss.logger.Error(err)
		return err
	}

	totalTime := stat.MeanQuizCompleteTime * float64(stat.QuizzesCompleted)
	stat.QuizzesCompleted += 1
	stat.MeanQuizCompleteTime = (totalQuizTime.Seconds() + totalTime) / float64(stat.QuizzesCompleted)

	err = ss.statRepo.Update(stat)
	if err != nil {
		ss.logger.Error(err)
	}
	return err
}

// Обновляет счетчик правильных ответов и среднее время ответа на вопрос
//
// Данный метод необходимо вызывать после того, как пользователь дал первый ответ на вопрос.
//
// Возвращает ошибку в следующих случаях:
//   - объект статистики для данного пользователя не найден
//   - ошибка при записи данных
func (ss *StatisticsService) SubmitAnswer(user model.User, isCorrect bool, answerTime time.Duration) error {
	ss.logger.
		WithFields(log.Fields{"user": user, "isCorrect": isCorrect, "answerTime": answerTime}).
		Info("SubmitAnswer")
	stat, err := ss.statRepo.FindByUserID(user.ID)
	if err != nil {
		ss.logger.Error(err)
		return err
	}

	totalTime := stat.MeanQuestionReplyTime * float64(stat.TotalReplies)

	stat.TotalReplies += 1
	if isCorrect {
		stat.CorrectReplies += 1
	}
	stat.CorrectRepliesPercent = float64(stat.CorrectReplies) / float64(stat.TotalReplies)

	totalTime += answerTime.Seconds()
	stat.MeanQuestionReplyTime = totalTime / float64(stat.TotalReplies)

	err = ss.statRepo.Update(stat)
	if err != nil {
		ss.logger.Error(err)
	}
	return err
}

// Сбрасывает статистику пользователя в значения по-умолчанию (нули)
//
// Возвращает ошибку в следующих случаях:
//   - объект статистики для данного пользователя не найден
//   - ошибка при записи данных
func (ss *StatisticsService) ResetStatistics(user model.User) error {
	ss.logger.
		WithFields(log.Fields{"user": user}).
		Info("ResetStatistics")
	stat, err := ss.statRepo.FindByUserID(user.ID)
	if err != nil {
		ss.logger.Error(err)
		return err
	}

	stat = model.Statistics{UserID: stat.UserID}

	err = ss.statRepo.Update(stat)
	if err != nil {
		ss.logger.Error(err)
	}
	return err
}
