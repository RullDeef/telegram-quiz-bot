package mem_repo

import (
	"fmt"

	"github.com/RullDeef/telegram-quiz-bot/model"
)

type QuizRepository struct {
	lastId  int64
	quizzes []model.Quiz
}

func NewQuizRepository() *QuizRepository {
	return &QuizRepository{
		lastId:  1,
		quizzes: nil,
	}
}

func (qr *QuizRepository) Create(q model.Quiz) (model.Quiz, error) {
	q.ID = qr.lastId
	qr.quizzes = append(qr.quizzes, q)
	return q, nil
}

func (qr *QuizRepository) FindAll() ([]model.Quiz, error) {
	return qr.quizzes, nil
}

func (qr *QuizRepository) FindByID(id int64) (model.Quiz, error) {
	for _, q := range qr.quizzes {
		if q.ID == id {
			return q, nil
		}
	}
	return model.Quiz{}, fmt.Errorf(`quiz with id="%d" not found`, id)
}

func (qr *QuizRepository) FindByTopic(topic string) (model.Quiz, error) {
	for _, q := range qr.quizzes {
		if q.Topic == topic {
			return q, nil
		}
	}
	return model.Quiz{}, fmt.Errorf(`quiz with topic="%s" not found`, topic)
}

func (qr *QuizRepository) Update(quiz model.Quiz) error {
	for i, q := range qr.quizzes {
		if q.ID == quiz.ID {
			qr.quizzes[i] = quiz
			return nil
		}
	}
	return fmt.Errorf(`quiz with id="%d" not found`, quiz.ID)
}

func (qr *QuizRepository) Delete(quiz model.Quiz) error {
	for i, q := range qr.quizzes {
		if q.ID == quiz.ID {
			qr.quizzes = append(qr.quizzes[:i], qr.quizzes[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf(`quiz with id="%d" not found`, quiz.ID)
}
