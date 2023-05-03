package controller

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	model "github.com/RullDeef/telegram-quiz-bot/model"
	log "github.com/sirupsen/logrus"
)

const (
	defaultGatherPlayerTimeout  = 30 * time.Second
	defaultWaitForAnswerTimeout = 30 * time.Second
)

const (
	pauseCommand  = "/пауза"
	resumeCommand = "/прод"
)

const confirmActionID = iota

type SessionController struct {
	userService          model.UserService
	statService          model.StatisticsService
	quizService          model.QuizService
	interactor           model.Interactor
	logger               *log.Logger
	state                *model.SessionState
	gatherPlayerTimeout  time.Duration
	waitForAnswerTimeout time.Duration
}

func NewSessionController(
	userService model.UserService,
	statService model.StatisticsService,
	quizService model.QuizService,
	interactor model.Interactor,
	logger *log.Logger,
) *SessionController {
	return &SessionController{
		userService:         userService,
		statService:         statService,
		quizService:         quizService,
		interactor:          interactor,
		logger:              logger,
		gatherPlayerTimeout: defaultGatherPlayerTimeout,
	}
}

// Устанавливает максимальное время ожидания при сборе игроков
func (c *SessionController) SetGatherPlayerTimeout(timeout time.Duration) {
	c.gatherPlayerTimeout = timeout
}

// Устанавливает максимальное время ожидания ответа на вопрос
func (c *SessionController) SetWaitForAnswerTimeout(timeout time.Duration) {
	c.waitForAnswerTimeout = timeout
}

// Запускает квиз
//
// Во время выполнения данного метода происходит следующее:
// 1. Выбираются участники для игры
// 2. Если участников меньше двух - квиз досрочно завершается
// 3. Создается случайный квиз
// 4. Каждый вопрос в квизе показывается пользователям
// 5. Собирается статистика с участников квиза
func (c *SessionController) Run() {
	defer c.panicRecoverer("Извините, произошла ошибка на сервере. Квиз остановлен.")

	users := c.gatherPlayers()
	if len(users) < 2 {
		return
	}

	quiz, err := c.quizService.CreateRandomQuiz()
	if err != nil {
		panic(err)
	}

	c.state = model.NewSessionState(quiz, users)
	startTime := time.Now()

	for _, question := range c.state.Quiz.Questions {
		c.state.CurrentQuestion = &question
		c.askQuestion(question)
	}

	quizDuration := time.Since(startTime)
	for _, user := range users {
		err := c.statService.SubmitQuizComplete(*user, quizDuration)
		if err != nil {
			c.logger.Error(err)
		}
	}

	c.sendResponse("Квиз завершен. Спасибо за участие!")
}

// Задаёт вопрос игрокам и ожидает ответов от них
func (c *SessionController) askQuestion(question model.Question) {
	answeredUsers := make(map[string]bool)

	c.showQuestion(question)
	timer := time.NewTimer(c.waitForAnswerTimeout)
	startTime := time.Now()

	for {
		select {
		case msg := <-c.interactor.MessageChan():
			if msg.Text == pauseCommand {
				c.PauseQuiz()
			} else if c.dispatchAnswerOption(msg, question, answeredUsers, startTime) {
				c.sendResponse("%s дает верный ответ!", msg.Sender.Nickname)
				return
			}
		case <-c.state.WaitForPause():
			c.waitForResume()
			timer = time.NewTimer(c.waitForAnswerTimeout)
		case <-timer.C:
			// time is up
			c.sendResponse("Никто не дал правильного ответа.")
			return
		}
	}
}

// Обрабатывает ответ на вопрос от пользователя
//
// В случае, если пользователь уже давал ответ на данный вопрос, ответ игнорируется.
//
// Возвращает true, если пользователь дал первый правильный ответ, иначе false.
func (c *SessionController) dispatchAnswerOption(
	msg model.Message,
	q model.Question,
	answeredUsers map[string]bool,
	questionStartTime time.Time,
) bool {
	opt, ok := extractAnswerOption(msg.Text)
	if !ok || opt <= 0 || opt > len(q.Answers) || answeredUsers[msg.Sender.TelegramID] {
		return false
	}

	ans := q.Answers[opt-1]
	err := c.statService.SubmitAnswer(*msg.Sender, ans.IsСorrect, time.Since(questionStartTime))
	if err != nil {
		c.logger.Error(err)
	}

	answeredUsers[msg.Sender.TelegramID] = true
	return ans.IsСorrect
}

// Досрочное звершение квиза
func (c *SessionController) EndQuiz() {
	if c.state == nil {
		c.logger.Warn("called EndQuiz when no quiz running (c.state == null)")
	} else {
		c.logger.Warn("TODO: EndQuiz")
	}
}

// Ставит квиз на паузу
func (c *SessionController) PauseQuiz() {
	if c.state == nil {
		c.logger.Warn("called PauseQuiz when no quiz running (c.state == null)")
	} else {
		c.state.Pause()
		c.sendResponse("Квиз приостановлен. Напишите /прод, чтобы возобновить.")
	}
}

// Возобновляет квиз
func (c *SessionController) ResumeQuiz() {
	if c.state == nil {
		c.logger.Warn("called ResumeQuiz when no quiz running (c.state == null)")
	} else {
		c.state.Resume()
		c.sendResponse("Квиз возобновлён.")
		c.showQuestion(*c.state.CurrentQuestion) // show last question again
	}
}

// Отображает переданный вопрос и варианты ответа к нему
func (c *SessionController) showQuestion(question model.Question) {
	var answers []string
	for i, ans := range question.Answers {
		answers = append(answers, fmt.Sprintf("%d. %s", i+1, ans.Text))
	}

	c.sendResponse("%s\n\n%s", question.Text, strings.Join(answers, "\n"))
}

// Собирает пользователей, которые будут участвовать в квизе
//
// Сбор осуществляется посредством отображения сообщения с кнопкой "Я участвую".
// Пользователь, нажавший кнопку, добавляется в список участников.
// По окончании сбора выводится информационное сообщение об участниках квиза.
func (c *SessionController) gatherPlayers() []*model.User {
	resp := model.NewResponse("Сбор участников квиза.")
	resp.AddAction(confirmActionID, "Я участвую")
	c.interactor.SendResponse(resp)

	var users []*model.User
	timer := time.NewTimer(c.gatherPlayerTimeout)

outer:
	for {
		select {
		case msg := <-c.interactor.MessageChan():
			if msg.ActionID() == confirmActionID {
				users, _ = appendUniqueUser(users, msg.Sender)
			}
		case <-timer.C:
			break outer
		}
	}

	c.informGatheringEnded(users)
	return users
}

// Оповещает об окончании сбора людей для игры
//
// В случае, если никто не участвует, выводится соответствующее сообщение.
//
// В случае, если участвует всего 1 человек, выводится сообщение о том,
// что одного человека недостаточно, чтобы играть в квиз.
//
// В остальных случаях выводится оповещение о начале игры и список участников.
func (c *SessionController) informGatheringEnded(users []*model.User) {
	switch len(users) {
	case 0:
		c.sendResponse("Никто не захотел участвовать в квизе.")
	case 1:
		c.sendResponse("Одного человека недостаточно для игры в квиз.")
	default:
		var usernames []string
		for _, u := range users {
			usernames = append(usernames, u.Nickname)
		}
		c.sendResponse("Начинаем квиз! Список участников:\n%s.", strings.Join(usernames, ",\n"))
	}
}

// Регистрирует команду для продолжения квиза и продолжает квиз
//
// Пока состояние квиза не изменится на "возобновлено", выполнение блоируется.
func (c *SessionController) waitForResume() {
	for {
		select {
		case msg := <-c.interactor.MessageChan():
			if msg.Text == resumeCommand {
				c.ResumeQuiz()
			}
		case <-c.state.WaitForResume():
			return
		}
	}
}

// Отлавливает проброшенные ошибки и оповещает пользователей о непредвиденных обстоятельствах
func (c *SessionController) panicRecoverer(recoverMessage string) {
	if err := recover(); err != nil {
		c.logger.Error(err)
		c.sendResponse(recoverMessage)
	}
}

// Отправляет ответ, одновременно логируя отправленное сообщение
func (c *SessionController) sendResponse(format string, args ...interface{}) {
	msgText := fmt.Sprintf(format, args...)

	c.logger.Info(msgText)
	c.interactor.SendResponse(model.NewResponse(msgText))
}

// Добавляет пользователя в список, если его там нет
//
// Возвращает список пользователей и флаг, показывающий, был ли добавлен пользователь
func appendUniqueUser(users []*model.User, user *model.User) ([]*model.User, bool) {
	for _, u := range users {
		if u.TelegramID == user.TelegramID {
			return users, false
		}
	}
	return append(users, user), true
}

// Считывает номер ответа из сообщения пользователя
//
// Возвращает номер ответа и флаг, показывающий успех считывания
func extractAnswerOption(message string) (int, bool) {
	opt, err := strconv.Atoi(message)
	return opt, err == nil
}
