package controller

import (
	model "github.com/RullDeef/telegram-quiz-bot/model"
	"github.com/RullDeef/telegram-quiz-bot/service"
)

type AdminController struct {
	userService *service.UserService
	interactor  model.Interactor
}

func NewAdminController(
	userService *service.UserService,
	interactor model.Interactor,
) *AdminController {
	return &AdminController{
		userService: userService,
		interactor:  interactor,
	}
}

func (c *AdminController) CreateQuiz() {

}

func (c *AdminController) ViewMyQuizzes() {

}

func (c *AdminController) EditQuiz() {

}

func (c *AdminController) CommitEdit() {

}
