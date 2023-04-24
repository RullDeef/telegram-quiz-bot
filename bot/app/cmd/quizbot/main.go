package main

import (
	"os"

	"github.com/RullDeef/telegram-quiz-bot/manager"
	"github.com/RullDeef/telegram-quiz-bot/model"
	"github.com/RullDeef/telegram-quiz-bot/repository/orm"
	"github.com/RullDeef/telegram-quiz-bot/tginteractor"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logger := log.New()

	dsn := "host=db user=postgres password=root port=5432 dbname=quizdb"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		logger.Fatal(err)
	}

	userRepo := orm.NewUserRepo(db)
	statRepo := orm.NewStatisticsRepo(db)

	publisher := tginteractor.NewTGBotPublisher(os.Getenv("TELEGRAM_API_TOKEN"))

	botMngr := manager.NewBotManager(func(bm *manager.BotManager, i int64, c chan model.Message) model.Interactor {
		return tginteractor.NewInteractor(publisher, i, c)
	}, userRepo, statRepo, logger)

	publisher.Run(botMngr)
}
