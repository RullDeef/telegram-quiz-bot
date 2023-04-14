package orm

import (
	"errors"
	"fmt"

	"github.com/RullDeef/telegram-quiz-bot/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//New version
type UserRepositoryNewStruct struct {
	Db     *gorm.DB
	LastId int
}

func create_connection(host, user, password, dbname, port string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s", host, user, password, port, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})

	return db, err
}

func (ur *UserRepositoryNewStruct) Create(user model.User) error {
	return ur.Db.Create(&user).Error
}

func (ur *UserRepositoryNewStruct) FindByID(id int64) (model.User, error) {
	var user model.User
	result := ur.Db.Find(&user, id)

	err := result.Error

	if result.RowsAffected == 0 {
		err = errors.New("null")
	}

	return user, err
}

func (ur *UserRepositoryNewStruct) FindByTelegramID(id string) (model.User, error) {
	var user model.User
	result := ur.Db.Where("telegram_id = ?", id).Find(&user)

	err := result.Error

	if result.RowsAffected == 0 {
		err = errors.New("null")
	}

	return user, err
}

func (ur *UserRepositoryNewStruct) Update(user model.User) error {
	result := ur.Db.Table("users").Where("id = ?", user.ID).Updates(&user)

	err := result.Error

	if result.RowsAffected == 0 {
		err = errors.New("null")
	}

	return err
}

func (ur *UserRepositoryNewStruct) Delete(id int64) error {
	result := ur.Db.Table("users").Delete(&model.User{}, id)

	err := result.Error

	if result.RowsAffected == 0 {
		err = errors.New("null")
	}

	return err
}
