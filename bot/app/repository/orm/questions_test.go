package orm

import (
	"testing"

	"github.com/RullDeef/telegram-quiz-bot/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestQuestionInterface(t *testing.T) {
	dsn := "host=testdb user=postgres password=root port=5432 dbname=quizdb"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		t.Errorf("Get connection to db = %s; want nil", err)
		t.FailNow()
	}

	question_repo := NewQuestionsRepository(db)

	question_add := model.Question{Text: "Где диплом?", Topic: "Диплом"}
	question_update := model.Question{Text: "Делается", Topic: "Диплом"}

	t.Run("Create", func(t *testing.T) {
		question_add, err = question_repo.Create(question_add)
		if err != nil {
			t.Errorf("Create question err = %s; want nil", err)
		}
		question_update.ID = question_add.ID
	})

	t.Run("FindByID", func(t *testing.T) {
		_, err := question_repo.FindById(question_add.ID)
		if err != nil {
			t.Errorf("FindByID no one question; want 1")
		}
	})

	t.Run("FindByIDFail", func(t *testing.T) {
		_, err := question_repo.FindById(question_add.ID + 200)
		if err == nil {
			t.Errorf("FindByID got question; want 0")
		}
	})

	t.Run("FindByTopic", func(t *testing.T) {
		_, err := question_repo.FindByTopic(question_add.Topic)
		if err != nil {
			t.Errorf("FindByTopic no one question; want 1")
		}
	})

	t.Run("FindByTopicFail", func(t *testing.T) {
		questions, err := question_repo.FindByTopic("Random Topic")
		if err != nil {
			t.Errorf("Find question err = %s; want nil", err)
		}
		if len(questions) > 0 {
			t.Errorf("FindByTopic got questions; want 0")
		}
	})

	t.Run("Update", func(t *testing.T) {
		err = question_repo.Update(question_update)
		if err != nil {
			t.Errorf("Update no one question; want 1")
		}
	})

	t.Run("Delete", func(t *testing.T) {
		err = question_repo.Delete(question_add.ID)
		if err != nil {
			t.Errorf("Delete no one question; want 1")
		}
	})
}
