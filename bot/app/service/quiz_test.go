package service

import (
	"fmt"
	"testing"

	"github.com/RullDeef/telegram-quiz-bot/model"
	mem_repo "github.com/RullDeef/telegram-quiz-bot/repository/mem"
)

func TestQuizService(t *testing.T) {
	question_repo := mem_repo.NewQuestionsRepository()
	quiz_service := NewQuizService(question_repo)

	var question_id int64

	// заполнение репозитория тестовыми вопросами
	for _, topic := range []string{"lisp", "prolog", "python", "Go"} {
		for question_num := 1; question_num < 30; question_num++ {
			var answers []model.Answer
			for _, answer_num := range []int{1, 2, 3, 4} {
				answers = append(answers, model.Answer{
					Text:      fmt.Sprintf("тестовый ответ %d на вопрос %d", answer_num, question_num),
					IsСorrect: false,
				})
			}

			correct_index := question_num % len(answers)
			answers[correct_index].IsСorrect = true

			_, err := question_repo.Create(model.Question{
				Topic:   topic,
				Text:    fmt.Sprintf("тестовый вопрос %d", question_num),
				Answers: answers,
			})

			if err != nil {
				t.Errorf(`failed to initialize repository: %s`, err)
				t.FailNow()
			}
		}
	}

	t.Run("AddQuestionToTopic: ok", func(t *testing.T) {
		question := "Кто проживает на дне океана?"
		topic := "lisp"

		q, err := quiz_service.AddQuestionToTopic(question, topic)
		question_id = q

		if err != nil {
			t.Errorf("Add question to topic: err = %s; want nil", err)
		}
	})

	t.Run("AddQuestionToTopic: empty topic", func(t *testing.T) {
		topic := ""
		question := "WHO?"
		if _, err := quiz_service.AddQuestionToTopic(topic, question); err == nil {
			t.Error("Add question to topic: expected err, got nil")
		}
	})

	t.Run("AddQuestionToTopic: empty question", func(t *testing.T) {
		topic := "prolog"
		question := ""
		if _, err := quiz_service.AddQuestionToTopic(topic, question); err == nil {
			t.Error("Add question to topic: expected err, got nil")
		}
	})

	t.Run("AddAnswer", func(t *testing.T) {
		text := "Губка Боб"
		correct := true
		if err := quiz_service.AddAnswer(question_id, text, correct); err != nil {
			t.Errorf("Add answer: err = %s; want nil", err)
		}
	})

	t.Run("AddAnswer: another correct answer", func(t *testing.T) {
		text := "Губка Боб"
		correct := true
		if err := quiz_service.AddAnswer(question_id, text, correct); err == nil {
			t.Errorf("Add answer: err is nil; want not nil")
		}
	})

	t.Run("AddAnswer: empty answer", func(t *testing.T) {
		text := ""
		correct := false
		if err := quiz_service.AddAnswer(question_id, text, correct); err == nil {
			t.Errorf("Add answer: err is nil; want not nil")
		}
	})

	t.Run("AddAnswer: wrong answer - ok", func(t *testing.T) {
		text := "Мистер Пиклз"
		correct := false
		if err := quiz_service.AddAnswer(question_id, text, correct); err != nil {
			t.Errorf("Add answer: err is %s; want nil", err)
		}
	})

	t.Run("Create Quiz: ok", func(t *testing.T) {
		topic := "lisp"
		_, err := quiz_service.CreateQuiz(topic)

		if err != nil {
			t.Errorf("Create quiz err = %s; want nil", err)
		}
	})
}
