package controller

import (
	"fmt"
	"strings"
	"testing"

	"github.com/RullDeef/telegram-quiz-bot/mock_model"
	"github.com/RullDeef/telegram-quiz-bot/mockinteractor"
	"github.com/RullDeef/telegram-quiz-bot/model"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
)

func TestCreateQuestion(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserService := mock_model.NewMockUserService(mockCtrl)
	mockQuizService := mock_model.NewMockQuizService(mockCtrl)
	mockInteractor := mockinteractor.New()
	defer mockInteractor.Dispose()
	defer mockInteractor.AssertErrors(t)

	controller := NewAdminController(mockUserService, mockQuizService, mockInteractor, &logrus.Logger{})

	t.Run("should create valid question", func(t *testing.T) {
		user := &model.User{
			Nickname:   "pepe",
			TelegramID: "pepega",
			Role:       model.UserRoleAdmin,
		}
		var questionID int64 = 404
		answers := []string{"correct answer", "wrong answer 1", "wrong answer 2", "wrong answer 3"}

		mockInteractor.Expect("Выберите тематику квиза.")
		mockInteractor.SlipMessages(user, "test topic")
		mockInteractor.Expect("Запишите формулировку вопроса.")
		mockInteractor.SlipMessages(user, "test question?")
		mockInteractor.Expect("Введите верный ответ в текстовом виде.")
		mockInteractor.SlipMessages(user, answers[0])
		mockInteractor.Expect("Введите неверные ответы тремя сообщениями.")
		mockInteractor.SlipMessages(user, answers[1:]...)
		mockInteractor.Expect("Вопросы и ответы были успешно добавлены.")

		mockQuizService.EXPECT().
			AddQuestionToTopic("test topic", "test question?").
			Return(questionID, nil).
			Times(1)

		for i, ans := range answers {
			mockQuizService.EXPECT().
				AddAnswer(questionID, ans, i == 0).
				Return(nil).
				Times(1)
		}

		controller.CreateQuestion()
	})
}

func TestViewQuestions(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserService := mock_model.NewMockUserService(mockCtrl)
	mockQuizService := mock_model.NewMockQuizService(mockCtrl)
	mockInteractor := mockinteractor.New()
	defer mockInteractor.Dispose()
	defer mockInteractor.AssertErrors(t)

	controller := NewAdminController(mockUserService, mockQuizService, mockInteractor, &logrus.Logger{})

	user := &model.User{
		Nickname:   "pepe",
		TelegramID: "pepega",
		Role:       model.UserRoleAdmin,
	}

	t.Run("should show no questions", func(t *testing.T) {
		mockInteractor.Expect("Выберите тематику квиза.")
		mockInteractor.SlipMessages(user, "unknown topic")
		mockInteractor.Expect("Список вопросов для данной тематике пуст.")

		mockQuizService.EXPECT().
			ViewQuestionsByTopic("unknown topic").
			Return(nil, nil).
			Times(1)

		controller.ViewQuestions()
	})

	t.Run("should show single page of questions", func(t *testing.T) {
		questionList := []string{
			"1. question 1",
			"2. question 2",
			"3. question 3",
		}
		questions := strings.Join(questionList, "\n")

		mockInteractor.Expect("Выберите тематику квиза.")
		mockInteractor.SlipMessages(user, "valid topic")
		mockInteractor.Expect(fmt.Sprintf("Список вопросов:\n%s", questions))

		mockQuizService.EXPECT().
			ViewQuestionsByTopic("valid topic").
			Return(questionList, nil).
			Times(1)

		controller.ViewQuestions()
	})

	t.Run("should show multiple pages of questions", func(t *testing.T) {
		questionList := []string{
			"1. question 1",
			"2. question 2",
			"3. question 3",
			"4. question 4",
			"5. question 5",
			"6. question 6",
			"7. question 7",
			"8. question 8",
			"9. question 9",
			"10. question 10",
			"11. question 11",
		}
		questionsPage1 := strings.Join(questionList[:10], "\n")
		questionsPage2 := strings.Join(questionList[10:], "\n")

		mockInteractor.Expect("Выберите тематику квиза.")
		mockInteractor.SlipMessages(user, "valid topic")
		mockInteractor.Expect(fmt.Sprintf("Список вопросов:\n%s", questionsPage1))
		mockInteractor.SlipButtonAction(user, actionNextPage)
		mockInteractor.Expect(fmt.Sprintf("Список вопросов:\n%s", questionsPage2))

		mockQuizService.EXPECT().
			ViewQuestionsByTopic("valid topic").
			Return(questionList, nil).
			Times(1)

		controller.ViewQuestions()
	})
}

func TestEditQuiz(t *testing.T) {

}
