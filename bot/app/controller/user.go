package controller

import (
	"fmt"
	"time"

	model "github.com/RullDeef/telegram-quiz-bot/model"
	log "github.com/sirupsen/logrus"
)

const (
	defaultWaitDuration = time.Duration(30) * time.Second

	actionGameRules = iota
	actionStatistics
)

type UserController struct {
	userRepo   model.UserRepository
	statRepo   model.StatisticsRepository
	interactor model.Interactor
	logger     *log.Logger
}

func NewUserController(
	userRepo model.UserRepository,
	statRepo model.StatisticsRepository,
	interactor model.Interactor,
	logger *log.Logger,
) *UserController {
	return &UserController{
		userRepo:   userRepo,
		statRepo:   statRepo,
		interactor: interactor,
		logger:     logger,
	}
}

// Изменяет имя пользователя
//
// Соответствует команде `/ник`
func (uc *UserController) ChangeNickname() {
	uc.sendResponse("Напишите, как мне Вас теперь называть?")

	for {
		msg, err := uc.waitForNextMessageWithTimeout(defaultWaitDuration)
		if err != nil {
			response := "Время для смены никнейма вышло. Для смены никнейма повторно используйте команду /ник."
			uc.sendResponse(response)
			break
		}

		if isNicknameValid(msg.Text) {
			uc.sendResponse("Некорректный никнейм, выберите другой.")
			continue
		}

		msg.Sender.Nickname = msg.Text
		err = uc.userRepo.Update(*msg.Sender)
		if err != nil {
			uc.sendResponse("Не удалось обновить никнейм. Попробуйте снова через какое-то время.")
		} else {
			response := "Ваш новый никнейм сохранен. Рад иметь с вами дело, %s."
			uc.sendResponse(response, msg.Text)
		}
		break
	}
}

// Показывает справку с правилами игры и со статистикой
//
// Соответсвует команде `/помощь`
func (uc *UserController) ShowHelp() {
	responseText := "Для получения подробной информации используйте кнопки ниже."
	response := model.NewResponse(responseText)
	response.AddAction(actionGameRules, "Правила игры")
	response.AddAction(actionStatistics, "Статистика")
	uc.interactor.SendResponse(response)

	for {
		msg, err := uc.waitForNextMessageWithTimeout(defaultWaitDuration)
		if err != nil {
			break
		}

		switch msg.ActionID() {
		case actionGameRules:
			uc.showGameRules()
		case actionStatistics:
			uc.showStatisticsForUser(msg.Sender.ID)
		}
	}
}

// Показывает правила игры
//
// Соответствует команде `/правила`
func (uc *UserController) showGameRules() {
	uc.sendResponse(`Правила игры:

Тут должны быть правила, но их пока нет.`)
}

// Выводит статистику конкретного пользователя
//
// Соответствует команде `/статистика`
func (uc *UserController) showStatisticsForUser(userID int64) {
	stat, err := uc.statRepo.FindByUserID(userID)
	if err != nil {
		// must never happen - statistics must me created with user at the same time
		uc.logger.
			WithField("userID", userID).
			Error("failed to find statistics")
		uc.sendResponse("К сожалению, произошла ошибка. Попробуйте повторить ваш запрос позже.")
	} else {
		uc.sendResponse(`Ваша статистика:

Количество игр: %d.
Среднее время одной игры: %.1f сек.
Среднее время ответа: %.1f сек.
Верных ответов: %d (%d %%).`,
			stat.QuizzesCompleted,
			stat.MeanQuizCompleteTime,
			stat.MeanQuestionReplyTime,
			stat.CorrectReplies,
			int(stat.CorrectRepliesPercent*100))
	}
}

// Отправляет ответ, одновременно логируя отправленное сообщение
func (uc *UserController) sendResponse(format string, args ...interface{}) {
	msgText := fmt.Sprintf(format, args...)

	uc.logger.Info(msgText)
	uc.interactor.SendResponse(model.NewResponse(msgText))
}

// Ожидает получения следующего сообщения от пользователя
//
// По истечении заданного времени возвращает ошибку
func (uc *UserController) waitForNextMessageWithTimeout(timeout time.Duration) (model.Message, error) {
	select {
	case <-time.After(timeout):
		uc.logger.WithField("duration", timeout).Warn("timeout")
		return model.Message{}, fmt.Errorf("timed out")
	case msg := <-uc.interactor.MessageChan():
		uc.logger.
			WithField("senderID", msg.Sender.ID).
			WithField("nickname", msg.Sender.Nickname).
			WithField("msg", msg.Text).
			Info(`got message`)
		return msg, nil
	}
}

// Проверяет корректность ника пользователя
func isNicknameValid(nickname string) bool {
	return len(nickname) > 0
}
