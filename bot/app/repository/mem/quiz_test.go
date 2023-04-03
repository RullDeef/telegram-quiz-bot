package mem_repo

import (
	"testing"

	"github.com/RullDeef/telegram-quiz-bot/model"
)

func TestQuizInterface(t *testing.T) {
	var _ model.QuizRepository = &QuizRepository{}
}
