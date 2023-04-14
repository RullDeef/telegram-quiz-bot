package orm

// func TestQuestionInterface(t *testing.T) {
// 	var err error
// 	question_repo := &QuestionsRepositoryStruct{}
// 	dsn := "host=testdb user=postgres password=root port=5432 dbname=quizdb"
// 	question_repo.Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
// 	if err != nil {
// 		t.Errorf("Get connection to db = %s; want nil", err)
// 		t.FailNow()
// 	}

// 	question_add := model.Question{ID: 1, Text: "Где диплом?"}
// 	question_update := model.Question{ID: 1, Text: "Делается"}

// 	t.Run("Create", func(t *testing.T) {
// 		err = question_repo.Create(question_add)

// 		if err != nil {
// 			t.Errorf("Create question err = %s; want nil", err)
// 		}
// 	})

// 	t.Run("FindByID", func(t *testing.T) {
// 		var question_id int64 = 1

// 		_, err := question_repo.FindById(question_id)

// 		if err != nil {
// 			t.Errorf("FindByID no one question; want 1")
// 		}
// 	})

// 	t.Run("Update", func(t *testing.T) {
// 		err = question_repo.Update(question_update)

// 		if err != nil {
// 			t.Errorf("Update no one question; want 1")
// 		}
// 	})

// 	t.Run("Delete", func(t *testing.T) {
// 		var question_id int64 = 1

// 		err = question_repo.Delete(question_id)

// 		if err != nil {
// 			t.Errorf("Delete no one question; want 1")
// 		}
// 	})
// }
