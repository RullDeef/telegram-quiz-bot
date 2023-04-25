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
// Данный метод необходимо вызывать один раз для каждого пользователя после завершения квиза
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

// Обновляет счетчик правильных ответов
//
// Также обновляет среднее время ответа на вопрос.
// Данный метод необходимо вызывать в случае, если пользователь дал верный ответ на вопрос
func (ss *StatisticsService) SubmitCorrectAnswer(user model.User, answerTime time.Duration) error {
	ss.logger.
		WithFields(log.Fields{"user": user, "answerTime": answerTime}).
		Info("SubmitCorrectAnswer")
	stat, err := ss.statRepo.FindByUserID(user.ID)
	if err != nil {
		ss.logger.Error(err)
		return err
	}

	totalTime := stat.MeanQuestionReplyTime * float64(stat.TotalReplies)

	stat.TotalReplies += 1
	stat.CorrectReplies += 1
	stat.CorrectRepliesPercent = float64(stat.CorrectReplies) / float64(stat.TotalReplies)

	totalTime += answerTime.Seconds()
	stat.MeanQuestionReplyTime = totalTime / float64(stat.TotalReplies)

	err = ss.statRepo.Update(stat)
	if err != nil {
		ss.logger.Error(err)
	}
	return err
}

// Обновляет счетчик правильных ответов
//
// Также обновляет среднее время ответа на вопрос.
// Данный метод необходимо вызывать в случае, если пользователь дал первый ошибочный ответ на вопрос
func (ss *StatisticsService) SubmitWrongAnswer(user model.User, answerTime time.Duration) error {
	ss.logger.
		WithFields(log.Fields{"user": user, "answerTime": answerTime}).
		Info("SubmitCorrectAnswer")
	stat, err := ss.statRepo.FindByUserID(user.ID)
	if err != nil {
		ss.logger.Error(err)
		return err
	}

	totalTime := stat.MeanQuestionReplyTime * float64(stat.TotalReplies)

	stat.TotalReplies += 1
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
