package service

import model "github.com/RullDeef/telegram-quiz-bot/model"

type UserService struct {
	UserRepo model.UserRepository
}

func NewUserService(UserRepository model.UserRepository) UserService {
	return UserService{UserRepository};
}

func (s *UserService) CreateUser(username string, telegramId string) bool {
	temp := s.UserRepo.FindByTelegramID(telegramId)
	if (temp != nil) {
		return false;
	}
	var user model.User;
	user.Nickname = username;
	user.TelegramID = telegramId;
	user.Role = "user";
	s.UserRepo.Create(user);
	return true;
}

func (s *UserService) SetUserRole(role string, userId string) bool {
	temp := s.UserRepo.FindByTelegramID(userId);
	if (temp != nil) {
		return false;
	}
	temp.Role = role;
	s.UserRepo.Update(temp);
	return true;
}

func (s *UserService) GetUserByTelegramId(id string) model.User {
	temp := s.UserRepo.FindByTelegramID(id);
	if (temp != nil) {
		return temp;
	}
	return nil;
}
