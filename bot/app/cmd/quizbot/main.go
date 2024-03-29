package main

import (
	"fmt"
	"os"

	"github.com/RullDeef/telegram-quiz-bot/manager"
	"github.com/RullDeef/telegram-quiz-bot/model"
	"github.com/RullDeef/telegram-quiz-bot/repository/orm"
	"github.com/RullDeef/telegram-quiz-bot/service"
	"github.com/RullDeef/telegram-quiz-bot/tginteractor"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logger := log.New()
	logger.Level = log.DebugLevel

	db, err := buildDBConnection()
	if err != nil {
		logger.Fatal(err)
	}

	userRepo := orm.NewUserRepo(db)
	statRepo := orm.NewStatisticsRepo(db)
	questionRepo := orm.NewQuestionsRepository(db)

	userService := service.NewUserService(userRepo)
	statService := service.NewStatisticsService(userRepo, statRepo, logger)
	quizService := service.NewQuizService(questionRepo)

	publisher := tginteractor.NewTGBotPublisher(os.Getenv("TELEGRAM_API_TOKEN"), userRepo, logger)

	botMngr := manager.NewBotManager(func(bm *manager.BotManager, i int64, c chan model.Message) model.Interactor {
		return tginteractor.NewInteractor(publisher, i, c)
	}, userService, statService, quizService, logger)

	publisher.Run(botMngr)
}

func buildDBConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DBNAME"),
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
}
