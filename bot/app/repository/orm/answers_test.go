package orm

// func TestAnswersInterface(t *testing.T) {
// 	var err error
// 	answer_repo := &AnswerRepositoryStruct{}
// 	dsn := "host=testdb user=postgres password=root port=5432 dbname=quizdb"
// 	answer_repo.Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
// 	if err != nil {
// 		t.Errorf("Get connection to db = %s; want nil", err)
// 		t.FailNow()
// 	}

// 	answer_add := model.Answer{ID: 1, Text: "Существует", IsСorrect: true}
// 	answer_upd := model.Answer{ID: 1, Text: "Не Существует", IsСorrect: false}

// 	t.Run("Create", func(t *testing.T) {
// 		err = answer_repo.Create(answer_add)

// 		if err != nil {
// 			t.Errorf("Create Answer err = %s; want nil", err)
// 		}
// 	})

// 	t.Run("FindByAnswerId", func(t *testing.T) {
// 		var answer_id int64 = 1

// 		_, err := answer_repo.FindByAnswerId(answer_id)

// 		if err != nil {
// 			t.Errorf("FindByAnswerId no one answer found; want 1")
// 		}
// 	})

// 	t.Run("FindByQuestionId", func(t *testing.T) {
// 		var answer_qid int64 = 1

// 		_, err := answer_repo.FindByQuestionId(answer_qid)

// 		if err != nil {
// 			t.Errorf("FindByQuestionId found no one; want > 0")
// 		}
// 	})

// 	t.Run("Update", func(t *testing.T) {
// 		err = answer_repo.Update(answer_upd)
// 		if err != nil {
// 			t.Errorf("Update no one answer; want 1")
// 		}
// 	})

// 	t.Run("Delete", func(t *testing.T) {
// 		var answer_qid int64 = 1

// 		err = answer_repo.Delete(answer_qid)

// 		if err != nil {
// 			t.Errorf("Delete no one answer; want 1")
// 		}
// 	})
// }
