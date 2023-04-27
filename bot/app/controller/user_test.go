package controller

import (
	"testing"

	"github.com/RullDeef/telegram-quiz-bot/mockinteractor"
	"github.com/RullDeef/telegram-quiz-bot/model"
	mem_repo "github.com/RullDeef/telegram-quiz-bot/repository/mem"
	log "github.com/sirupsen/logrus"
)

func TestUserController_Register(t *testing.T) {
	_, _, interactor, controller := initUserController()
	defer interactor.Dispose()
	defer interactor.AssertErrors(t)

	userVasya := model.User{
		Nickname:   "Vasya",
		TelegramID: "vasyatg",
		Role:       model.UserRoleUser,
	}

	interactor.Expect("Вы успешно зарегистрированы под ником Vasya.")

	controller.Register(userVasya)
}

func TestUserController_RegisterTwice(t *testing.T) {
	userRepo, statRepo, interactor, controller := initUserController()
	defer interactor.Dispose()
	defer interactor.AssertErrors(t)

	userVasya := model.User{
		Nickname:   "Vasya",
		TelegramID: "vasyatg",
		Role:       model.UserRoleUser,
	}

	userVasya, err := userRepo.Create(userVasya)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	stat := model.Statistics{UserID: userVasya.ID}
	if err = statRepo.Create(stat); err != nil {
		t.Error(err)
		t.FailNow()
	}

	interactor.Expect("Вы уже зарегистрированы.")

	controller.Register(userVasya)
}

func TestUserController_ChangeNickname(t *testing.T) {
	userRepo, statRepo, interactor, controller := initUserController()
	defer interactor.Dispose()
	defer interactor.AssertErrors(t)

	userVasya := model.User{
		Nickname:   "Vasya",
		TelegramID: "vasyatg",
		Role:       model.UserRoleUser,
	}

	userVasya, err := userRepo.Create(userVasya)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	stat := model.Statistics{UserID: userVasya.ID}
	if err = statRepo.Create(stat); err != nil {
		t.Error(err)
		t.FailNow()
	}

	interactor.Expect("Напишите, как мне Вас теперь называть?")
	interactor.SlipMessages(&userVasya, "Вася")
	interactor.Expect("Ваш новый никнейм сохранен. Рад иметь с вами дело, Вася.")

	controller.ChangeNickname()
}

func initUserController() (*mem_repo.UserRepository, *mem_repo.StatisticsRepository, *mockinteractor.MockInteractor, *UserController) {
	userRepo := mem_repo.NewUserRepository()
	statRepo := mem_repo.NewStatisticsRepository()
	logger := log.New()
	interactor := mockinteractor.New()

	ctrl := NewUserController(
		userRepo,
		statRepo,
		interactor,
		logger,
	)

	return userRepo, statRepo, interactor, ctrl
}
