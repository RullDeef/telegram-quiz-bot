package controller

import (
	"fmt"
	"strings"

	model "github.com/RullDeef/telegram-quiz-bot/model"
	log "github.com/sirupsen/logrus"
)

const (
	actionNextPage = iota
)

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
		ac.logger.Error(err)
		ac.sendResponse("Произошла непредвиденная ошибка.")
		return
	}

	//Это для контроля страницы с выведенными вопросами
	currPage := 1
	nQuestionsPerPage := 10

	// 3. показать по 10 штук и кнопку "вперед"
dance:
	for {
		questionsPage := paginate(questions, currPage, nQuestionsPerPage)

		resp := model.NewResponse(fmt.Sprintf("Список вопросов:\n%s", strings.Join(questionsPage, "\n")))
		resp.AddAction(actionNextPage, "Вперед")

		msg = <-ac.interactor.MessageChan()
		if msg.IsButtonAction {
			currPage += 1
		} else {
			break dance
		}
	}
}

func (ac *AdminController) EditQuestion() {
	//1. Получение тематики может сразу id вопроса? A. оке
	// Тогда надо делать метод просмотра всех вопросов (не по тематике)
	//Сложна, ес чессна решить
	// А, сделать обертку над той функцией ViewQuestions
	// Сделать отдельно для тематики, отдельно для всех тематик срзу
	// можно

	// 2. Получение от пользователя п*здюлей
	//го проверим то что уже написали... => go

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

// Отправляет ответ, одновременно логируя отправленное сообщение
func (ac *AdminController) sendResponse(format string, args ...interface{}) {
	msgText := fmt.Sprintf(format, args...)

	ac.logger.Info(msgText)
	ac.interactor.SendResponse(model.NewResponse(msgText))
}

// возвращает страницу page из данных data
func paginate(data []string, page, itemsPerPage int) []string {
	i := (page - 1) * itemsPerPage
	j := page * itemsPerPage
	return data[i:j]
}
