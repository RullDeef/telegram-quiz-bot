package mockinteractor

import (
	"testing"

	"github.com/RullDeef/telegram-quiz-bot/model"
)

func TestInteractorInterface(t *testing.T) {
	var _ model.Interactor = &MockInteractor{}
}
