package controller

import (
	"fmt"
	"time"

	model "github.com/RullDeef/telegram-quiz-bot/model"
)

type SessionController struct {
	userRepo   model.UserRepository
	interactor model.Interactor
	state      *model.SessionState // maybe not needed here as field
}

func NewSessionController(
	userRepo model.UserRepository,
	interactor model.Interactor,
) *SessionController {
	return &SessionController{
		userRepo:   userRepo,
		interactor: interactor,
		state:      nil,
	}
}

func (c *SessionController) Run() {
	// TODO: populate users that want to play this quiz

	resp := model.NewResponse("Сбор участников квиза.")
	resp.AddAction(1, "Я участвую")
	c.interactor.SendResponse(resp)

	var users []*model.User
	timer := time.NewTimer(30 * time.Second)

outer:
	for {
		select {
		case msg := <-c.interactor.MessageChan():
			if msg.ActionID() == 1 {
				users = append(users, msg.Sender) // register user
			}
		case <-timer.C:
			c.interactor.SendResponse(model.NewResponse("Время вышло! Начинаем квиз!"))
			c.interactor.SendResponse(model.NewResponse(fmt.Sprintf("users: %v", users)))
			break outer
		}
	}
	time.Sleep(30 * time.Second)

	// TODO: make a quiz builder
	// quiz := model.Quiz{
	// 	Questions: []model.Question{{}, {}, {}},
	// }

	// с.state = model.NewSessionState(quiz, users)

	c.interactor.SendResponse(model.NewResponse("quiz started!"))
	time.Sleep(30 * time.Second)
	c.interactor.SendResponse(model.NewResponse("quiz ended! thanks for playing!"))
	// TODO one of two ways for session controller realization

	// imperative all-in-one-func scenario: (more control, preferred)
	//   1. shuffle quiz questions,
	//   2. iterate over them in a for loop,
	//   3. send question using interactor,
	//   4. use `select` to wait for each user guess or timeout.

	// state-machine-like: (less control)
	// methods below are called from tgBotPublisher,
	// this function just waits until quiz ends.
}

func (c *SessionController) EndQuiz() {

}

func (c *SessionController) PauseQuiz() {

}

func (c *SessionController) ResumeController() {

}

func (c *SessionController) SubmitAnswer(option int) {

}
