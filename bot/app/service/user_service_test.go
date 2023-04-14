package service

import (
	"testing"

	model "github.com/RullDeef/telegram-quiz-bot/model"
	mem_repo "github.com/RullDeef/telegram-quiz-bot/repository/mem"
)

// Create new User from username + telegram id
// 1. No duplicate user can be added
// 2. When added returns true
// 3. User role is set in service and may differ
func TestUserService(t *testing.T) {
	var repo model.UserRepository = &mem_repo.UserRepository{}
	service := NewUserService(repo)
	nickname := "Johnny"
	telegramId := "telegramID#0"
	wrongRole := "superuser"
	existing_user := model.User{Nickname: nickname, TelegramID: telegramId, Role: wrongRole}
	repo.Create(existing_user)

	t.Run("Add duplicate user should return error", func(t *testing.T) {
		_, err := service.CreateUser(nickname, telegramId)
		if (err != nil) {
			t.Errorf("CreateUser duplicate user; want error")
		}
	})

	repo.Delete(existing_user)
	t.Run("Add new user should return true", func(t *testing.T) {
		_, err := service.CreateUser(nickname, telegramId)
		if (err != nil) {
			t.Errorf("CreateUser new user; want true")
		}
	})


	t.Run("Add new user should return not null user", func(t *testing.T) {
		res, _ := repo.FindByTelegramID(telegramId)
		if (res != model.User{} && res.Nickname != nickname && res.TelegramID != telegramId && res.Role == wrongRole) {
			t.Errorf("Role is not added accroding to inner rule")
		}
	})


}
