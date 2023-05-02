package controller

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"

	"github.com/RullDeef/telegram-quiz-bot/mock_model"
	"github.com/RullDeef/telegram-quiz-bot/mockinteractor"
	"github.com/RullDeef/telegram-quiz-bot/model"
)

func TestPlayersGathering(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserService := mock_model.NewMockUserService(mockCtrl)
	mockStatisticsService := mock_model.NewMockStatisticsService(mockCtrl)
	mockQuizService := mock_model.NewMockQuizService(mockCtrl)

	mockInteractor := mockinteractor.New()
	defer mockInteractor.Dispose()
	defer mockInteractor.AssertErrors(t)

	controller := NewSessionController(
		mockUserService,
		mockStatisticsService,
		mockQuizService,
		mockInteractor,
		&logrus.Logger{},
	)

	controller.SetGatherPlayerTimeout(time.Second)

	t.Run("should inform noone plays quiz", func(t *testing.T) {
		mockInteractor.Expect("Сбор участников квиза.")
		mockInteractor.Expect("Никто не захотел участвовать в квизе.")

		controller.Run()
	})

	t.Run("should inform single player not enough", func(t *testing.T) {
		testUser := &model.User{Nickname: "pepe", TelegramID: "pepega"}

		mockInteractor.Expect("Сбор участников квиза.")
		mockInteractor.SlipButtonAction(testUser, confirmActionID)

		mockInteractor.Expect("Одного человека недостаточно для игры в квиз.")

		controller.Run()
	})
}

func TestInterrupts(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserService := mock_model.NewMockUserService(mockCtrl)
	mockStatisticsService := mock_model.NewMockStatisticsService(mockCtrl)
	mockQuizService := mock_model.NewMockQuizService(mockCtrl)

	mockInteractor := mockinteractor.New()
	defer mockInteractor.Dispose()
	defer mockInteractor.AssertErrors(t)

	controller := NewSessionController(
		mockUserService,
		mockStatisticsService,
		mockQuizService,
		mockInteractor,
		&logrus.Logger{},
	)

	controller.SetGatherPlayerTimeout(time.Second)
	controller.SetWaitForAnswerTimeout(time.Second)

	t.Run("normal quiz, no answers", func(t *testing.T) {
		testUser1 := &model.User{Nickname: "pepe1", TelegramID: "pepega1"}
		testUser2 := &model.User{Nickname: "pepe2", TelegramID: "pepega2"}
		testQuiz := model.Quiz{Topic: "topic", Questions: []model.Question{
			{
				Text: "question1",
				Answers: []model.Answer{
					{
						Text: "answer1",
					},
					{
						Text:      "answer2",
						IsСorrect: true,
					},
				},
			},
		}}

		mockInteractor.Expect("Сбор участников квиза.")
		mockInteractor.SlipButtonAction(testUser1, confirmActionID)
		mockInteractor.SlipButtonAction(testUser2, confirmActionID)
		mockInteractor.Expect("Начинаем квиз! Список участников:\npepe1,\npepe2.")

		mockInteractor.Expect("question1\n\n1. answer1\n2. answer2")
		mockInteractor.Expect("Никто не дал правильного ответа.")
		mockInteractor.Expect("Квиз завершен. Спасибо за участие!")

		mockQuizService.EXPECT().
			CreateRandomQuiz().
			Return(testQuiz, nil)

		mockStatisticsService.EXPECT().
			SubmitQuizComplete(gomock.Eq(*testUser1), gomock.Any()).
			Return(nil).
			Times(1)

		mockStatisticsService.EXPECT().
			SubmitQuizComplete(gomock.Eq(*testUser2), gomock.Any()).
			Return(nil).
			Times(1)

		controller.Run()
	})
}
