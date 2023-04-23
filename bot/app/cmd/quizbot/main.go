package main

import (
	"os"

	"github.com/RullDeef/telegram-quiz-bot/manager"
	"github.com/RullDeef/telegram-quiz-bot/model"
	mem_repo "github.com/RullDeef/telegram-quiz-bot/repository/mem"
	"github.com/RullDeef/telegram-quiz-bot/tginteractor"
)

func main() {
	publisher := tginteractor.NewTGBotPublisher(os.Getenv("TELEGRAM_API_TOKEN"))

	userRepo := mem_repo.NewUserRepository()

	botMngr := manager.NewBotManager(func(bm *manager.BotManager, i int64, c chan model.Message) model.Interactor {
		return tginteractor.NewInteractor(publisher, i, c)
	}, userRepo)

	publisher.Run(botMngr)
}
