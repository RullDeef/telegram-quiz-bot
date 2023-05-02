package mem_repo

import (
	"testing"

	"github.com/RullDeef/telegram-quiz-bot/model"
)

func TestStatisticsInterface(t *testing.T) {
	var _ model.StatisticsRepository = &StatisticsRepository{}
}
