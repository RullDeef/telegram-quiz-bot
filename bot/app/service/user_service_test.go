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
func TestUserServiceCreateUser(t *testing.T) {
	var repo model.UserRepository = &mem_repo.UserRepository{}
	service := NewUserService(repo)
	nickname := "Johnny"
	telegramId := "telegramID#0"
	wrongRole := "superuser"
	existing_user := model.User{Nickname: nickname, TelegramID: telegramId, Role: wrongRole}
	_, err := repo.Create(existing_user)
	if (err != nil) {
		t.Errorf("Create database error")
	}

	t.Run("Add duplicate user should return error", func(t *testing.T) {
		_, err := service.CreateUser(nickname, telegramId)
		if (err != nil) {
			t.Errorf("CreateUser duplicate user; want error")
		}
	})

	err := repo.Delete(existing_user)
	if (err != nil) {
		t.Errorf("Delete database error")
	}
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

// Change user role
// 1. Role is updating return true
// 2. User not found -> return false
func TestUserServiceSetUserRole(t *testing.T) {
	var repo model.UserRepository = &mem_repo.UserRepository{}
	service := NewUserService(repo)
	nickname := "Johnny"
	telegramId := "telegramID#1"
	wrongTelegramId := "telegramID#2"
	role := "user"
	updateRole := "admin"
	existing_user := model.User{Nickname: nickname, TelegramID: telegramId, Role: role}
	_, err := repo.Create(existing_user)
	if (err != nil) {
		t.Errorf("Create database error")
	}

	t.Run("Update role should return true", func(t *testing.T) {
		status := service.SetUserRole(updateRole, telegramId)
		if (status != true) {
			t.Errorf("User is not existing should return false")
		}
	})

	t.Run("Update role of not existing user should return false", func(t *testing.T) {
		status := service.SetUserRole(updateRole, wrongTelegramId)
		if (status != false) {
			t.Errorf("User is not existing should return false")
		}
	})
}

// Get user by telegram id
// 1. User is returned by correct id
// 2. User is not found by wrong id
func TestUserServiceGetUserByTelegramId(t *testing.T) {
	var repo model.UserRepository = &mem_repo.UserRepository{}
	service := NewUserService(repo)
	nickname := "Johnny"
	telegramId := "telegramID#1"
	wrongTelegramId := "telegramID#2"
	role := "user"
	existing_user := model.User{Nickname: nickname, TelegramID: telegramId, Role: role}
	_, err = repo.Create(existing_user)
	if (err != nil) {
		t.Errorf("Create database error")
	}

	t.Run("Get user with correct telegram id should return user struct", func(t *testing.T) {
		user, err := service.GetUserByTelegramId(telegramId)
		if (err != nil) {
			t.Errorf("Find user with telegram id should return user")
		}
		if (user.Nickname != nickname || user.TelegramID != telegramId || user.Role != role) {
			t.Errorf("Found wrong user")
		}
	})

	t.Run("Get user with wrong telegram id should return error", func(t *testing.T) {
		_, err := service.GetUserByTelegramId(wrongTelegramId)
		if (err == nil) {
			t.Errorf("Found user with not existing id")
		}
	})
}
