package service

import (
	"errors"
	"log"

	model "github.com/RullDeef/telegram-quiz-bot/model"
)

type UserService struct {
	UserRepo model.UserRepository
}

func NewUserService(UserRepository model.UserRepository) UserService {
	return UserService{UserRepository}
}

func (s *UserService) CreateUser(username string, telegramId string) (model.User, error) {
	_, err := s.UserRepo.FindByTelegramID(telegramId)
	if (err != nil) {
		return model.User{}, errors.New("Duplicate user")
	}
	var user model.User
	user.Nickname = username
	user.TelegramID = telegramId
	user.Role = "user"
	temp, err := s.UserRepo.Create(user)
	if (err != nil) {
		return model.User{}, errors.New("Database error")
	}
	return temp, nil
}

func (s *UserService) SetUserRole(role string, telegramId string) bool {
	temp, err := s.UserRepo.FindByTelegramID(telegramId)
	if err != nil {
		return false
	}
	temp.Role = role
	err = s.UserRepo.Update(temp)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return false
	}
	return true
}

func (s *UserService) GetUserByTelegramId(id string) (model.User, error) {
	temp, err := s.UserRepo.FindByTelegramID(id)
	if err != nil {
		return model.User{}, errors.New("No user found")
	}
	return temp, nil
}
