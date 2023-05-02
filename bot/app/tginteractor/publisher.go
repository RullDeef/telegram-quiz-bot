package tginteractor

import (
	"fmt"
	"strconv"
	"time"

	"github.com/RullDeef/telegram-quiz-bot/manager"
	"github.com/RullDeef/telegram-quiz-bot/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

type TGBotPublisher struct {
	userRepo    model.UserRepository
	bot         *tgbotapi.BotAPI
	logger      *log.Logger
	chatMembers map[int64][]int64
}

func NewTGBotPublisher(
	token string,
	userRepo model.UserRepository,
	logger *log.Logger,
) *TGBotPublisher {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic("TELEGRAM_API_TOKEN env variable is not set")
	}

	// bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &TGBotPublisher{
		userRepo:    userRepo,
		bot:         bot,
		logger:      logger,
		chatMembers: make(map[int64][]int64),
	}
}

func (bp *TGBotPublisher) Run(mngr *manager.BotManager) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bp.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			mngr.DispatchMessage(bp.tgMessageToModel(update.Message))
		} else if update.CallbackQuery != nil { // If someone waits an inline keyboard action
			mngr.DispatchMessage(bp.tgCallbackToModel(update.CallbackQuery))
		}
	}
}

func (bp *TGBotPublisher) tgCallbackToModel(query *tgbotapi.CallbackQuery) model.Message {
	buttonActionID, err := strconv.ParseInt(query.Data, 10, 64)
	if err != nil {
		log.Printf("ERROR: %v", err)
	}
	return model.Message{
		ChatID:         query.Message.Chat.ID,
		Sender:         bp.tgUserToModel(query.From),
		Text:           "",
		ReceiveTime:    time.Now(),
		IsButtonAction: true,
		IsPrivate:      query.Message.Chat.IsPrivate(),
		ButtonActionID: buttonActionID,
	}
}

func (bp *TGBotPublisher) tgMessageToModel(message *tgbotapi.Message) model.Message {
	return model.Message{
		ChatID:         message.Chat.ID,
		Sender:         bp.tgUserToModel(message.From),
		Text:           message.Text,
		ReceiveTime:    time.Now(),
		IsPrivate:      message.Chat.IsPrivate(),
		IsButtonAction: false,
		ButtonActionID: model.INVALID_BUTTON_ACTION_ID,
	}
}

func (bp *TGBotPublisher) tgUserToModel(user *tgbotapi.User) *model.User {
	modelUser, err := bp.userRepo.FindByTelegramID(user.UserName)
	if err == nil {
		// Пользователь существует в базе
		return &modelUser
	}

	modelUser = model.User{
		Nickname:   fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		TelegramID: user.UserName,
		Role:       modelUser.Role,
	}

	// Иначе необходимо зарегистрировать пользователя
	modelUser, err = bp.userRepo.Create(modelUser)
	if err != nil {
		bp.logger.WithField("modelUser", modelUser).Error(err)
	}

	return &modelUser
}
