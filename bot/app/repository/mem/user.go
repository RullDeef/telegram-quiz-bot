package mem_repo

import (
	"errors"

	"github.com/RullDeef/telegram-quiz-bot/model"
)

type UserRepository struct {
	lastId int
	users  []model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		lastId: 1,
		users:  nil,
	}
}

func (ur *UserRepository) Create(user model.User) model.User {
	user.ID = int64(len(ur.users))
	ur.users = append(ur.users, user)
	return user
}

func (ur *UserRepository) FindByID(id int64) (model.User, error) {
	for _, u := range ur.users {
		if u.ID == id {
			return u, nil
		}
	}
	return model.User{}, errors.New("not found")
}

func (ur *UserRepository) FindByTelegramID(id string) (model.User, error) {
	for _, u := range ur.users {
		if u.TelegramID == id {
			return u, nil
		}
	}
	return model.User{}, errors.New("not found")
}

func (ur *UserRepository) Update(user model.User) error {
	for i, u := range ur.users {
		if u.ID == user.ID {
			ur.users[i] = user
			return nil
		}
	}
	return errors.New("not found")
}

func (ur *UserRepository) Delete(user model.User) {
	for i, u := range ur.users {
		if u.ID == user.ID {
			ur.users = append(ur.users[:i], ur.users[i+1:]...)
			break
		}
	}
}
