package orm

import (
	"testing"
	"time"

	"github.com/RullDeef/telegram-quiz-bot/model"
)

func TestQuizInterface(t *testing.T) {
	var quiz_repo QuizRepositoryStruct
	var err error

	quiz_repo.Db, err = create_connection("localhost", "postgres", "root", "quizdb_test", "5432")

	if err != nil {
		t.Errorf("Get connection to db = %s; want nil", err)
	}

	t.Run("Create", func(t *testing.T) {
		quiz := model.QuizNew{ID: 1, Topic: "lala", Creator_id: 1, Created_at: time.Now()}
		err = quiz_repo.Create(quiz)

		if err != nil {
			t.Errorf("Create Quiz err = %s; want nil", err)
		}
	})

	t.Run("FindAll", func(t *testing.T) {
		var all_quizes []model.QuizNew
		all_quizes, err = quiz_repo.FindAll()

		n_quizzes := len(all_quizes)
		if n_quizzes == 0 {
			t.Errorf("FindAll no one quiz; want > 0")
		}
	})

	t.Run("FindByTopicIsFound", func(t *testing.T) {
		_, err = quiz_repo.FindByTopic("lala")

		if err != nil {
			t.Errorf("FindByTopicIsFound quiz count = %d; want > 0", 0)
		}
	})

	t.Run("FindByTopicNotFound", func(t *testing.T) {
		_, err = quiz_repo.FindByTopic("lolo")

		if err == nil {
			t.Errorf("FindByTopicNotFound found lolo topic; want 0")
		}
	})

	t.Run("FindByIDFound", func(t *testing.T) {
		var quiz_id int64
		quiz_id = 1

		_, err = quiz_repo.FindByID(quiz_id)

		if err != nil {
			t.Errorf("FindByIDFound found 0 quizes; want > 0")
		}
	})

	t.Run("Update", func(t *testing.T) {
		quiz := model.QuizNew{ID: 1, Topic: "lele", Creator_id: 1, Created_at: time.Now()}

		err = quiz_repo.Update(quiz)

		if err != nil {
			t.Errorf("Update no one row updated; want 1")
		}
	})

	t.Run("Delete", func(t *testing.T) {
		var quiz_id int64
		quiz_id = 1

		err = quiz_repo.Delete(quiz_id)

		if err != nil {
			t.Errorf("Delete no one row updated; want 1")
		}
	})
}
