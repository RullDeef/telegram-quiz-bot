package orm

import (
	"testing"

	"github.com/RullDeef/telegram-quiz-bot/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// test interface realization
func TestQuizRepoInterface(t *testing.T) {
	var _ model.QuizRepository = &ORMQuizRepository{}
}

func TestQuizInterface(t *testing.T) {
	dsn := "host=testdb user=postgres password=root port=5432 dbname=quizdb"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		t.Errorf("Get connection to db = %s; want nil", err)
		t.FailNow()
	}

	user_repo := NewUserRepo(db)
	quiz_repo := NewQuizRepo(db)

	var quizID int64

	creator := model.User{
		Nickname:   "mega user",
		TelegramID: "mega_user",
		Role:       "USER",
	}

	t.Run("create creator", func(t *testing.T) {
		creator, err = user_repo.Create(creator)
		if err != nil {
			t.Errorf("failed to insert creator user")
			t.FailNow()
		}
	})

	t.Run("Create", func(t *testing.T) {
		quiz := model.Quiz{
			Topic:     "lala",
			Creator:   creator,
			Questions: nil,
		}
		quiz, err = quiz_repo.Create(quiz)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("quiz created: %v", quiz)
			quizID = quiz.ID
		}
	})

	t.Run("FindAll", func(t *testing.T) {
		all_quizes, err := quiz_repo.FindAll()
		if err != nil {
			t.Error(err)
		}
		if len(all_quizes) == 0 {
			t.Errorf("FindAll no one quiz; want > 0")
		}
	})

	t.Run("FindByTopicIsFound", func(t *testing.T) {
		_, err = quiz_repo.FindByTopic("lala")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("FindByTopicNotFound", func(t *testing.T) {
		_, err = quiz_repo.FindByTopic("lolo")
		if err == nil {
			t.Errorf("FindByTopicNotFound found lolo topic; want 0")
		}
	})

	t.Run("FindByIDFound", func(t *testing.T) {
		_, err = quiz_repo.FindByID(quizID)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Update", func(t *testing.T) {
		quiz := model.Quiz{
			ID:        quizID,
			Topic:     "lele",
			Creator:   creator,
			Questions: nil,
		}
		err = quiz_repo.Update(quiz)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		err = quiz_repo.Delete(model.Quiz{
			ID:        quizID,
			Topic:     "lele",
			Creator:   creator,
			Questions: nil,
		})
		if err != nil {
			t.Error(err)
		}
	})
}
