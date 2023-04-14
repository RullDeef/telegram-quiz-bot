package service

import (
	"testing"

	"github.com/RullDeef/telegram-quiz-bot/model"
	mem_repo "github.com/RullDeef/telegram-quiz-bot/mem_repo"
)

// test interface realization
func TestUserServiceCreation(t *testing.T) {
	var repo mem_repo.UserRepository = &UserRepository{}
	service := NewUserService(repo)
}

func TestUserService(t *testing.T) {

}
