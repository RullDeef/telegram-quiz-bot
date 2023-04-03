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

func (s *UserService) CreateUser(username string, telegramId string) bool {
	temp, _ := s.UserRepo.FindByTelegramID(telegramId)
	if temp != (model.User{}) {
		return false
	}
	var user model.User
	user.Nickname = username
	user.TelegramID = telegramId
	user.Role = "user"
	s.UserRepo.Create(user)
	return true
}

func (s *UserService) SetUserRole(role string, userId string) bool {
	temp, _ := s.UserRepo.FindByTelegramID(userId)
	if temp != (model.User{}) {
		return false
	}
	temp.Role = role
	err := s.UserRepo.Update(temp)
	if err != nil {
		log.Printf("ERROR: %v", err)
	}
	return true
}

func (s *UserService) GetUserByTelegramId(id string) (model.User, error) {
	temp, _ := s.UserRepo.FindByTelegramID(id)
	if temp != (model.User{}) {
		return temp, nil
	}
	return temp, errors.New("No user found")
}
