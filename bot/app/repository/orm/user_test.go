package orm

import (
	"testing"

	"github.com/RullDeef/telegram-quiz-bot/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// test interface realization
func TestUserRepoInterface(t *testing.T) {
	var _ model.UserRepository = &ORMUserRepository{}
}

func TestUserInterface(t *testing.T) {
	dsn := "host=testdb user=postgres password=root port=5432 dbname=quizdb"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		t.Errorf("Get connection to db = %s; want nil", err)
	}

	user_repo := NewUserRepo(db)

	//need one user for quiz tests
	user_add := model.User{Nickname: "Jacob", TelegramID: "some_id1", Role: "USER"}
	var userID int64
	user_update := model.User{Nickname: "James", TelegramID: "some_id1", Role: "USER"}

	t.Run("Create", func(t *testing.T) {
		user_add, err := user_repo.Create(user_add)
		if err != nil {
			t.Error(err)
		} else {
			userID = user_add.ID
			t.Logf("user %v created", user_add)
		}
	})

	t.Run("FindByID", func(t *testing.T) {
		_, err := user_repo.FindByID(userID)

		if err != nil {
			t.Errorf("FindByID no one user; want 1")
			t.Error(err)
		}
	})

	t.Run("FindByTelegramID", func(t *testing.T) {
		telegram_id := "some_id1"
		_, err = user_repo.FindByTelegramID(telegram_id)

		if err != nil {
			t.Errorf("FindByTelegramID no one user; want 1")
			t.Error(err)
		}
	})

	t.Run("Update", func(t *testing.T) {
		user_update.ID = userID
		err = user_repo.Update(user_update)

		if err != nil {
			t.Errorf("Update no one user; want 1")
			t.Error(err)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		err = user_repo.Delete(user_add)

		if err != nil {
			t.Errorf("Delete no one user; want 1")
			t.Error(err)
		}
	})
}
