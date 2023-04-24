// Данный модуль представляет репозиторий, отвечающий за Пользователь, Вопросы и Статистика
package orm

import (
	"fmt"

	"github.com/RullDeef/telegram-quiz-bot/model"
	"gorm.io/gorm"
)

// БД Сущность пользователя
type userEntity struct {
	// ID пользователя
	ID uint `gorm:"primaryKey"`

	// Имя пользователя
	Nickname string

	// Telegram ID пользователя
	TelegramId string `gorm:"column:telegram_id"`

	// роль пользователя (ADMIN, USER)
	Role string
}

type ORMUserRepository struct {
	db *gorm.DB
}

// Создание экземпляра репозитория User
func NewUserRepo(db *gorm.DB) *ORMUserRepository {
	return &ORMUserRepository{
		db: db,
	}
}

// Создание пользователей
//
// 	- user - модель пользователя
//
// Возвращается созданный пользователь с установленным ID, а также ошибка выполнения.
// В случае успешного выполнения возвращается nil
func (ur *ORMUserRepository) Create(user model.User) (model.User, error) {
	entity := userModelToEntity(user)
	err := ur.db.Create(&entity).Error
	if err != nil {
		return model.User{}, err
	}
	return userEntityToModel(entity), err
}

// Нахождение пользователя по его ID
//
// 	- id - ID пользователя
//
// Возвращается модель пользователя по найденному ID.
// В случае успешного выполнения возвращается nil. Возможные ошибки:
//  - Пользователь с данным идентификатором не найден
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

// Нахождение пользователя по его Telegram ID
//
// id - Telegram ID пользователя
//
// Возвращается модель пользователя по найденному Telegram ID.
//
// В случае успешного выполнения возвращается nil. Возможные ошибки:
//  - Пользователь с данным Telegram ID не найден
func (ur *ORMUserRepository) FindByTelegramID(id string) (model.User, error) {
	var entity userEntity
	err := ur.db.First(&entity, "telegram_id = ?", id).Error
	if err != nil {
		return model.User{}, fmt.Errorf(`user with telegram_id="%s" not found: %w`, id, err)
	}
	return userEntityToModel(entity), err
}

// Обновление пользователя по его ID
//
// id - ID пользователя
//
// В случае успешного выполнения возвращается nil. Возможные ошибки:
//  - Пользователь с данным идентификатором не найден
func (ur *ORMUserRepository) Update(user model.User) error {
	entity := userModelToEntity(user)
	err := ur.db.Updates(&entity).Error
	if err != nil {
		err = fmt.Errorf(`failed to update user with id="%d": %w`, user.ID, err)
	}
	return err
}

// Удаление пользователя из БД по его идентификатору
//
// user - модель пользователя
//
// В случае успешного выполнения возвращается nil. Возможные ошибки:
//  - Внутренние ошибки базы данных
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

// Перевод БД сущности Пользователь в модельную
func userEntityToModel(user userEntity) model.User {
	return model.User{
		ID:         int64(user.ID),
		Nickname:   user.Nickname,
		TelegramID: user.TelegramId,
		Role:       user.Role,
	}
}

// Перевод сущности Пользователь из модельную в сущность БД
func userModelToEntity(user model.User) userEntity {
	return userEntity{
		ID:         uint(user.ID),
		Nickname:   user.Nickname,
		TelegramId: user.TelegramID,
		Role:       user.Role,
	}
}
