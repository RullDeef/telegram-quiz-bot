package controller

import model "github.com/RullDeef/telegram-quiz-bot/model"

type AdminController struct {
	userRepo   model.UserRepository
	interactor model.Interactor
}

func NewAdminController(
	userRepo model.UserRepository,
	interactor model.Interactor,
) *AdminController {
	return &AdminController{
		userRepo:   userRepo,
		interactor: interactor,
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
