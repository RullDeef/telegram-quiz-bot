package service

import (
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/RullDeef/telegram-quiz-bot/repository/orm"
)

func TestQuizService(t *testing.T) {

	dsn := "host=testdb user=postgres password=root port=5432 dbname=quizdb"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		t.Errorf("Get connection to db = %s; want nil", err)
		t.FailNow()
	}
	question_repo := orm.NewQuestionsRepository(db)

	quiz_service := NewQuizService(question_repo)

	t.Run("Create Quiz: ok", func(t *testing.T) {
		topic := "lisp"
		_, err := quiz_service.CreateQuiz(topic)

		if err != nil {
			t.Errorf("Create quiz err = %s; want nil", err)
		}
	})

	t.Run("AddQuestionToTopic: ok", func(t *testing.T) {
		//вот здесь траблы: надо надо где-то определять last id
		// вай? Тут не надо айдишник указывать (при создании) да да да. Ну ето орм....
		//он сам определяется? искал, не нашел в question repo
		//орм сам добавить новый айди что ли  -- уже забыл; фига мощный чел он. Ааааааа, понял
		// да. ID primaryKey поле оно автоматически при создании записи заполнит
		// вот и отличная документация получилась к тесту)))))
		// оставим в качестве пасхалки Россинскому
		// тип будет ли он читать код?)
		// да)
		question_add := "Кто проживает на дне океана?"
		topic := "lisp"
		_, err := quiz_service.AddQuestionToTopic(question_add, topic)

		if err != nil {
			t.Errorf("Add question to topic: err = %s; want nil", err)
		}
	})

	t.Run("AddQuestionToTopic: exact same question", func(t *testing.T) {})
	t.Run("AddQuestionToTopic: empty question", func(t *testing.T) {
		question_add := ""
		topic := "prolog"
		_, err := quiz_service.AddQuestionToTopic(question_add, topic)

		if err == nil {
			t.Error("Add question to topic: expected err, got nil")
		}

		//а как мы это всё синхронизировать будем? А то меня начинает рубить спать
		// давай закоммитишь все что есть кроме тестов вот в этом файле и подем спать (меня тож рубит)
		// мб эти тесты оставить, только закомментировать их?
	})

	t.Run("AddAnswer", func(t *testing.T) {
		text := "Губка Боб"
		is_correct := true
		var question_id int64 = 1
		err := quiz_service.AddAnswer(question_id, text, is_correct)

		if err != nil {
			t.Errorf("Add answer: err = %s; want nil", err)
		}
	})

}
