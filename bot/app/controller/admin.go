package controller

import (
	"fmt"
	"strconv"
	"strings"

	model "github.com/RullDeef/telegram-quiz-bot/model"
	log "github.com/sirupsen/logrus"
)

const (
	actionNextPage = iota
	actionChangeCorrectAnswer
	actionCommitChanges
	actionCancelChanges
)

const actionShift = 100

type AdminController struct {
	userService model.UserService
	quizService model.QuizService
	interactor  model.Interactor
	logger      *log.Logger
}

func NewAdminController(
	userService model.UserService,
	quizService model.QuizService,
	interactor model.Interactor,
	logger *log.Logger,
) *AdminController {
	return &AdminController{
		userService: userService,
		quizService: quizService,
		interactor:  interactor,
		logger:      logger,
	}
}

// Добавление вопроса в тематику квиза
func (ac *AdminController) CreateQuestion() {
	// 1. Спросить тематику
	ac.sendResponse("Выберите тематику квиза.")
	msg := <-ac.interactor.MessageChan()
	topic := msg.Text

	// 2. Спросить текст вопроса
	ac.sendResponse("Запишите формулировку вопроса.")
	msg = <-ac.interactor.MessageChan()
	question := msg.Text

	// 3. Спросить верный ответ
	ac.sendResponse("Введите верный ответ в текстовом виде.")
	msg = <-ac.interactor.MessageChan()
	correctAnswer := msg.Text

	// 4--n. Спросить дополнительный вариант ответа
	ac.sendResponse("Введите неверные ответы тремя сообщениями.")

	var wrongAnswers []string
	for i := 0; i < 3; i++ {
		msg = <-ac.interactor.MessageChan()
		wrongAnswers = append(wrongAnswers, msg.Text)
	}

	// n+1. Добавить вопрос в базу
	err := ac.commitQuestion(topic, question, correctAnswer, wrongAnswers)
	if err != nil {
		ac.sendResponse("Произошла непредвиденная ошибка. Вопрос не был добавлен.")
	} else {
		ac.sendResponse("Вопросы и ответы были успешно добавлены.")
	}
}

// Просмотр вопросов по тематике
func (ac *AdminController) ViewQuestions() {
	// 1. Спросить тематику
	ac.sendResponse("Выберите тематику квиза.")
	msg := <-ac.interactor.MessageChan()
	topic := msg.Text

	// 2. Получить список вопросов
	questions, err := ac.quizService.ViewQuestionsByTopic(topic)
	if err != nil {
		ac.logger.WithField("topic", topic).Error(err)
		panic(err)
	}

	if len(questions) == 0 {
		ac.sendResponse("Список вопросов для данной тематике пуст.")
		return
	}

	//Это для контроля страницы с выведенными вопросами
	currPage := 1
	nQuestionsPerPage := 10

	// 3. показать по 10 штук и кнопку "вперед"
dance:
	for {
		questionsPage, hasMorePages := paginate(questions, currPage, nQuestionsPerPage)

		resp := model.NewResponse(fmt.Sprintf("Список вопросов:\n%s", strings.Join(questionsPage, "\n")))
		if hasMorePages {
			resp.AddAction(actionNextPage, "Вперед")
		}
		ac.interactor.SendResponse(resp)

		if hasMorePages {
			msg = <-ac.interactor.MessageChan()
			if msg.IsButtonAction {
				currPage += 1
			} else {
				break dance
			}
		} else {
			break dance
		}
	}
}

func (ac *AdminController) EditQuestion() {
	ac.sendResponse("Введите идентификатор вопроса для редактирования.")

	msg := <-ac.interactor.MessageChan()
	questionID, err := strconv.ParseInt(msg.Text, 10, 64)
	if err != nil {
		ac.sendResponse("Некорректный идентификатор вопроса. Редактирование отменено.")
		return
	}

	question, err := ac.quizService.ViewQuestionByID(questionID)
	if err != nil {
		ac.sendResponse("Не найден вопрос с указанным идентификатором. Редактирование отменено.")
		return
	}

	for {
		answers := ""
		for i, ans := range question.Answers {
			correctSign := ""
			if ans.IsСorrect {
				correctSign = "*"
			}
			answers += fmt.Sprintf("%s%d. %s\n", correctSign, i+1, ans.Text)
		}
		resp := model.NewResponse(fmt.Sprintf(`Вопрос: <<%s>>
Варианты ответа:
%s
Редактируйте ответы c помощью кнопок ниже.
Для смены корректного варианта ответа нажмите "*".
Для сохранения нажмите ✔.
Для отмены изменений нажмите ✘.`, question.Text, answers))
		for i := range question.Answers {
			resp.AddAction(answerIndexToActionID(i), strconv.Itoa(i+1))
		}
		resp.AddAction(actionChangeCorrectAnswer, "*")
		resp.AddAction(actionCommitChanges, "✔")
		resp.AddAction(actionCancelChanges, "✘")
		ac.interactor.SendResponse(resp)

		msg := <-ac.interactor.MessageChan()

		// обработка сообщения от пользователя
		if msg.IsButtonAction {
			if msg.ActionID() == actionChangeCorrectAnswer {
				// смена корректного ответа
				resp := model.NewResponse("Выберите корректный вариант ответа.")
				for i := range question.Answers {
					resp.AddAction(answerIndexToActionID(i), strconv.Itoa(i+1))
				}
				ac.interactor.SendResponse(resp)

				msg := <-ac.interactor.MessageChan()
				answerIndex := actionIDToAnswerIndex(msg.ActionID())

				// установка корректного варианта ответа
				for i := range question.Answers {
					question.Answers[i].IsСorrect = false
				}
				question.Answers[answerIndex].IsСorrect = true
			} else if msg.ActionID() == actionCommitChanges {
				// сохранение изменений
				err := ac.quizService.UpdateQuestion(question)
				if err != nil {
					panic(err)
				}

				ac.sendResponse("Изменения успешно сохранены.")
				break
			} else if msg.ActionID() == actionCancelChanges {
				// отмена изменений
				ac.sendResponse("Изменения отменены.")
				break
			} else {
				// редактирование ответа
				answerIndex := actionIDToAnswerIndex(msg.ActionID())
				ac.sendResponse("Редактирование ответа <<%s>>.\nВведите новое описание:", question.Answers[answerIndex].Text)

				msg := <-ac.interactor.MessageChan()
				question.Answers[answerIndex].Text = msg.Text
			}
		} else {
			ac.sendResponse(`Используйте кнопки выше для редактирования вопроса.
Для сохранения нажмите ✔. Для отмены изменений нажмите ✘.`)
		}
	}
}

func (ac *AdminController) commitQuestion(topic, question, correctAnswer string, wrongAnswers []string) error {
	questionID, err := ac.quizService.AddQuestionToTopic(topic, question)
	if err == nil {
		err = ac.quizService.AddAnswer(questionID, correctAnswer, true)
	}
	for _, answer := range wrongAnswers {
		if err == nil {
			err = ac.quizService.AddAnswer(questionID, answer, false)
		}
	}
	return err
}

func answerIndexToActionID(index int) int64 {
	return int64(index) + actionShift
}

func actionIDToAnswerIndex(actionID int64) int {
	return int(actionID - actionShift)
}

// Отправляет ответ, одновременно логируя отправленное сообщение
func (ac *AdminController) sendResponse(format string, args ...interface{}) {
	msgText := fmt.Sprintf(format, args...)

	ac.logger.Info(msgText)
	ac.interactor.SendResponse(model.NewResponse(msgText))
}

// возвращает страницу page из данных data и флаг, есть ли еще данные
func paginate(data []string, page, itemsPerPage int) ([]string, bool) {
	i := (page - 1) * itemsPerPage
	j := page * itemsPerPage

	if j < len(data) {
		return data[i:j], true
	} else {
		return data[i:], false
	}
}
