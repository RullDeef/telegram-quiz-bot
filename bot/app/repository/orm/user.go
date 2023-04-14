package orm

import (
	"fmt"

	"github.com/RullDeef/telegram-quiz-bot/model"
	"gorm.io/gorm"
)

type userEntity struct {
	ID         uint `gorm:"primaryKey"`
	Nickname   string
	TelegramId string `gorm:"column:telegram_id"`
	Role       string
}

type ORMUserRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *ORMUserRepository {
	return &ORMUserRepository{
		db: db,
	}
}

func (ur *ORMUserRepository) Create(user model.User) (model.User, error) {
	entity := userModelToEntity(user)
	err := ur.db.Create(&entity).Error
	if err != nil {
		return model.User{}, err
	}
	return userEntityToModel(entity), err
}

func (ur *ORMUserRepository) FindByID(id int64) (model.User, error) {
	entity := userEntity{
		ID: uint(id),
	}
	err := ur.db.First(&entity).Error
	if err != nil {
		return model.User{}, fmt.Errorf(`user with id="%d" not found: %w`, id, err)
	}
	return userEntityToModel(entity), err
}

func (ur *ORMUserRepository) FindByTelegramID(id string) (model.User, error) {
	var entity userEntity
	err := ur.db.First(&entity, "telegram_id = ?", id).Error
	if err != nil {
		return model.User{}, fmt.Errorf(`user with telegram_id="%s" not found: %w`, id, err)
	}
	return userEntityToModel(entity), err
}

func (ur *ORMUserRepository) Update(user model.User) error {
	entity := userModelToEntity(user)
	err := ur.db.Updates(&entity).Error
	if err != nil {
		err = fmt.Errorf(`failed to update user with id="%d": %w`, user.ID, err)
	}
	return err
}

func (ur *ORMUserRepository) Delete(user model.User) error {
	err := ur.db.Delete(&userEntity{}, user.ID).Error
	if err != nil {
		err = fmt.Errorf(`failed to delete user with id="%d": %w`, user.ID, err)
	}
	return err
}

func (userEntity) TableName() string {
	return "users"
}

func userEntityToModel(user userEntity) model.User {
	return model.User{
		ID:         int64(user.ID),
		Nickname:   user.Nickname,
		TelegramID: user.TelegramId,
		Role:       user.Role,
	}
}

func userModelToEntity(user model.User) userEntity {
	return userEntity{
		ID:         uint(user.ID),
		Nickname:   user.Nickname,
		TelegramId: user.TelegramID,
		Role:       user.Role,
	}
}
