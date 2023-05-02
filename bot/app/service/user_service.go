// Данный модуль реализует бизнес-логику, связанную с пользователем:
//   - создание нового пользователя по telegram id;
//   - смена роли существующего пользователя;
//   - получение пользователя из базы данных по его telegram id.
package service

import (
	"errors"
	"log"

	model "github.com/RullDeef/telegram-quiz-bot/model"
)

type UserService struct {
	UserRepo model.UserRepository
}

func NewUserService(UserRepository model.UserRepository) *UserService {
	return &UserService{UserRepository}
}

// Функция создает нового пользователя в базе данных.
//
// Формальные параметры: имя пользователя для использования в приложении (по умолчанию используется никнейм из telegram), telegram id.
//
// При успешном выполнении возвращает пользователя с созданным id базы данных.
//
//  1. В случае существования пользователя с полученным telegramId функция возвращает пустую структуру пользователя и ошибку дубликата пользователя.
//  2. В случае получения ошибки из репозитория при сохранении пользователя возвращает пустую структуру пользователя и ошибку базы данных.
func (s *UserService) CreateUser(username string, telegramId string) (model.User, error) {
	_, err := s.UserRepo.FindByTelegramID(telegramId)
	if err == nil {
		return model.User{}, errors.New("Duplicate user")
	}
	var user model.User
	user.Nickname = username
	user.TelegramID = telegramId
	user.Role = model.UserRoleUser
	temp, err := s.UserRepo.Create(user)
	if err != nil {
		return model.User{}, errors.New("Database error")
	}
	return temp, nil
}

// Функция заменяет роль пользователя на другую.
//
// Формальные параметры: role - роль, которая заменит текущую роль пользователя, telegram id для получения пользователя из базы.
//
// При успешном выполнении возвращает true.
//
//  1. Если пользователь с таким telegram id не найден в базе данных, возвращает false.
//  2. Если функция обновления пользователя репозитория вернула ошибку, возвращает false.
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

// Функция возвращает структуру пользователя из базы данных по telegram id.
//
// Формальные параметры: telegram id.
//
// При успешном выполнении возвращает пользователя.
//
// В случае отсутствия пользователя с таким id возвращает пустую структуру пользователя и ошибку отсутствия пользователя в базе данных.
func (s *UserService) GetUserByTelegramId(id string) (model.User, error) {
	temp, err := s.UserRepo.FindByTelegramID(id)
	if err != nil {
		return model.User{}, errors.New("No user found")
	}
	return temp, nil
}

// Функция заменяет имя пользователя на другое.
//
// Формальные параметры: username - имя пользователя, которое заменит текущее имя пользователя, telegram id для получения пользователя из базы.
//
// При успешном выполнении возвращает true.
//
//  1. Если пользователь с таким telegram id не найден в базе данных, возвращает false.
//  2. Если функция обновления пользователя репозитория вернула ошибку, возвращает false.
func (s *UserService) ChangeUsername(username string, telegramId string) bool {
	temp, err := s.UserRepo.FindByTelegramID(telegramId)
	if err != nil {
		return false
	}
	temp.Nickname = username
	err = s.UserRepo.Update(temp)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return false
	}
	return true
}
