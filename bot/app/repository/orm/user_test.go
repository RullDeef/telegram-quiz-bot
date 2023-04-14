package orm

import (
	"testing"

	"github.com/RullDeef/telegram-quiz-bot/model"
)

func TestUserInterface(t *testing.T) {
	var user_repo UserRepositoryNewStruct
	var err error

	user_repo.Db, err = create_connection("localhost", "postgres", "root", "quizdb_test", "5432")

	if err != nil {
		t.Errorf("Get connection to db = %s; want nil", err)
	}

	//need one user for quiz tests
	user_add := model.User{ID: 1, Nickname: "Jacob", TelegramID: "some_id1", Role: "USER"}
	user_update := model.User{ID: 1, Nickname: "James", TelegramID: "some_id1", Role: "USER"}

	t.Run("Create", func(t *testing.T) {
		err = user_repo.Create(user_add)

		if err != nil {
			t.Errorf("Create User err = %s; want nil", err)
		}
	})

	t.Run("FindByID", func(t *testing.T) {
		var user_id int64
		user_id = 1

		_, err := user_repo.FindByID(user_id)

		if err != nil {
			t.Errorf("FindByID no one user; want 1")
		}
	})

	t.Run("FindByTelegramID", func(t *testing.T) {
		telegram_id := "some_id1"
		_, err = user_repo.FindByTelegramID(telegram_id)

		if err != nil {
			t.Errorf("FindByTelegramID no one user; want 1")
		}
	})

	t.Run("Update", func(t *testing.T) {
		//bug: is_correct not update
		err = user_repo.Update(user_update)

		if err != nil {
			t.Errorf("Update no one user; want 1")
		}
	})

	t.Run("Delete", func(t *testing.T) {
		var user_id int64
		user_id = 1

		err = user_repo.Delete(user_id)

		if err != nil {
			t.Errorf("Delete no one user; want 1")
		}
	})
}
