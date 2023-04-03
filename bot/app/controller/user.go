package controller

import (
	"fmt"
	"log"

	model "github.com/RullDeef/telegram-quiz-bot/model"
)

type UserController struct {
	userRepo   model.UserRepository
	interactor model.Interactor
}

func NewUserController(repo model.UserRepository, interactor model.Interactor) *UserController {
	return &UserController{
		userRepo:   repo,
		interactor: interactor,
	}
}

func (uc *UserController) ChangeNickname() {
	response := "Напишите, как мне Вас теперь называть?"
	uc.interactor.SendResponse(model.NewResponse(response))

	// wait for next message
	msg := <-uc.interactor.MessageChan()

	// TODO: store new nickname (better check before!)
	log.Printf("[storing nickname: \"%s\"]", msg.Text)

	response = "Ваш никнейм сохранен. Рад иметь с вами дело, %s."
	response = fmt.Sprintf(response, msg.Text)
	uc.interactor.SendResponse(model.NewResponse(response))
}

func (uc *UserController) ShowHelp() {
	responseText := "Для получения подробной информации используйте кнопки ниже."
	response := model.NewResponse(responseText)
	response.AddAction(1, "Правила игры")
	response.AddAction(2, "Статистика")
	uc.interactor.SendResponse(response)

	msg := <-uc.interactor.MessageChan()
	switch msg.ActionID() {
	case 1:
		uc.interactor.SendResponse(model.NewResponse(
			"Правила игры: бла бла бла.",
		))
	case 2:
		uc.interactor.SendResponse(model.NewResponse(
			"Ваша статистика: бла бла бла",
		))
	}
}
