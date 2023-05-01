package mem_repo

import (
	"errors"
	"fmt"

	"github.com/RullDeef/telegram-quiz-bot/model"
)

type QuestionsRepository struct {
	lastId    int
	questions []model.Question
}

func NewQuestionsRepository() *QuestionsRepository {
	return &QuestionsRepository{
		lastId:    1,
		questions: nil,
	}
}

func (ur *QuestionsRepository) Create(q model.Question) (model.Question, error) {
	q.ID = int64(len(ur.questions))
	ur.questions = append(ur.questions, q)
	return q, nil
}

func (ur *QuestionsRepository) FindByID(id int64) (model.Question, error) {
	for _, q := range ur.questions {
		if q.ID == id {
			return q, nil
		}
	}
	return model.Question{}, errors.New("not found")
}

func (ur *QuestionsRepository) FindByTopic(topic string) ([]model.Question, error) {
	var res []model.Question
	for _, q := range ur.questions {
		if q.Topic == topic {
			res = append(res, q)
		}
	}
	return res, nil
}

func (ur *QuestionsRepository) GetAllTopics() ([]string, error) {
	topics := make(map[string]bool)
	for _, q := range ur.questions {
		topics[q.Topic] = true
	}

	var res []string
	for topic := range topics {
		res = append(res, topic)
	}
	return res, nil
}

func (ur *QuestionsRepository) Update(q model.Question) error {
	for i, u := range ur.questions {
		if u.ID == q.ID {
			ur.questions[i] = q
			return nil
		}
	}
	return errors.New("not found")
}

func (ur *QuestionsRepository) Delete(id int64) error {
	for i, u := range ur.questions {
		if u.ID == id {
			ur.questions = append(ur.questions[:i], ur.questions[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("question with id=%d not found", id)
}
