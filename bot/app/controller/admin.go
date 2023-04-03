package controller

import model "github.com/RullDeef/telegram-quiz-bot/model"

type AdminController struct {
	quizRepo   model.QuizRepository
	userRepo   model.UserRepository
	interactor model.Interactor
}

func NewAdminController(
	quizRepo model.QuizRepository,
	userRepo model.UserRepository,
	interactor model.Interactor,
) *AdminController {
	return &AdminController{
		quizRepo:   quizRepo,
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
