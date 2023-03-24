package service

import model "github.com/RullDeef/telegram-quiz-bot/model"

type UserService struct {
	UserRepo *model.UserRepository
}

func (s *UserService) CreateUser(username string, telegramId string) bool {
	temp := UserRepo.findByTelegramId(u.TelegramId)
	if (temp != nil) {
		return false;
	}
	var user User;
	user.Nickname := username;
	user.TelegramId := telegramId;
	user.Role := "user";
	UserRepo.Create(user);
	return true;
}

func (s *UserService) SetUserRole(role string, userId string) bool {
	temp := UserRepo.findByTelegramId(userId);
	if (temp != nil) {
		return false;
	}
	temp.Role := role;
	UserRepo.Update(temp);
	return true;
}

func (s *UserService) GetUserByTelegramId(id string) User {
	temp := UserRepo.findByTelegramId(id);
	if (temp != nil) {
		return temp;
	}
	return nil;
}
