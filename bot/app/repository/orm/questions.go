package orm

// type QuestionsRepositoryStruct struct {
// 	Db     *gorm.DB
// 	LastId int
// }

// func (qr *QuestionsRepositoryStruct) Create(question model.Question) error {
// 	return qr.Db.Table("questions").Create(&question).Error
// }

// func (qr *QuestionsRepositoryStruct) FindById(id int64) (model.Question, error) {
// 	var question model.Question
// 	result := qr.Db.Table("questions").Find(&question, id)

// 	err := result.Error

// 	if result.RowsAffected == 0 {
// 		err = errors.New("null")
// 	}

// 	return question, err
// }

// func (qr *QuestionsRepositoryStruct) Update(question model.Question) error {
// 	result := qr.Db.Table("questions").Where("id = ?", question.ID).Updates(&question)

// 	err := result.Error

// 	if result.RowsAffected == 0 {
// 		err = errors.New("null")
// 	}

// 	return err
// }

// func (qr *QuestionsRepositoryStruct) Delete(id int64) error {
// 	result := qr.Db.Table("questions").Delete(&model.Question{}, id)

// 	err := result.Error

// 	if result.RowsAffected == 0 {
// 		err = errors.New("null")
// 	}

// 	return err
// }
