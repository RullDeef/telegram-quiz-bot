package mem_repo

import (
	"testing"

	"github.com/RullDeef/telegram-quiz-bot/model"
)

func TestUserInterface(t *testing.T) {
	var _ model.UserRepository = &UserRepository{}
}
