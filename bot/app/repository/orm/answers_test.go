package orm

import (
	"fmt"
	"testing"

	"github.com/RullDeef/telegram-quiz-bot/model"
)

func TestAnswersInterface(t *testing.T) {
	var answer_repo AnswerRepositoryStruct
	var err error

	answer_repo.Db, err = create_connection("localhost", "postgres", "root", "quizdb_test", "5432")

	if err != nil {
		t.Errorf("Get connection to db = %s; want nil", err)
	}

	answer_add := model.Answer{ID: 1, Question_ID: 1, Text: "Существует", Is_correct: true}
	answer_upd := model.Answer{ID: 1, Question_ID: 1, Text: "Не Существует", Is_correct: false}

	t.Run("Create", func(t *testing.T) {
		err = answer_repo.Create(answer_add)

		if err != nil {
			t.Errorf("Create Answer err = %s; want nil", err)
		}
	})

	t.Run("FindByAnswerId", func(t *testing.T) {
		var answer_id int64
		answer_id = 1

		_, err := answer_repo.FindByAnswerId(answer_id)

		if err != nil {
			t.Errorf("FindByAnswerId no one answer found; want 1")
		}
	})

	t.Run("FindByQuestionId", func(t *testing.T) {
		var answer_qid int64
		answer_qid = 1

		_, err := answer_repo.FindByQuestionId(answer_qid)

		if err != nil {
			t.Errorf("FindByQuestionId found no one; want > 0")
		}
	})

	t.Run("Update", func(t *testing.T) {
		err = answer_repo.Update(answer_upd)
		fmt.Print(answer_upd.Is_correct)
		if err != nil {
			t.Errorf("Update no one answer; want 1")
		}
	})

	t.Run("Delete", func(t *testing.T) {
		var answer_qid int64
		answer_qid = 1

		err = answer_repo.Delete(answer_qid)

		if err != nil {
			t.Errorf("Delete no one answer; want 1")
		}
	})
}
