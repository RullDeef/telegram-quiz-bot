package tginteractor

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/RullDeef/telegram-quiz-bot/manager"
	"github.com/RullDeef/telegram-quiz-bot/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TGBotPublisher struct {
	bot         *tgbotapi.BotAPI
	chatMembers map[int64][]int64
}

func NewTGBotPublisher(token string) *TGBotPublisher {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic("TELEGRAM_API_TOKEN env variable is not set")
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &TGBotPublisher{
		bot:         bot,
		chatMembers: make(map[int64][]int64),
	}
}

func (bp *TGBotPublisher) Run(mngr *manager.BotManager) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bp.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			mngr.DispatchMessage(tgMessageToModel(update.Message))
		} else if update.CallbackQuery != nil { // If someone waits an inline keyboard action
			mngr.DispatchMessage(tgCallbackToModel(update.CallbackQuery))
		}
	}
}

func tgCallbackToModel(query *tgbotapi.CallbackQuery) model.Message {
	buttonActionID, err := strconv.ParseInt(query.Data, 10, 64)
	if err != nil {
		log.Printf("ERROR: %v", err)
	}
	return model.Message{
		ChatID:         query.Message.Chat.ID,
		Sender:         tgUserToModel(query.From),
		Text:           "",
		ReceiveTime:    time.Now(),
		IsButtonAction: true,
		IsPrivate:      query.Message.Chat.IsPrivate(),
		ButtonActionID: buttonActionID,
	}
}

func tgMessageToModel(message *tgbotapi.Message) model.Message {
	return model.Message{
		ChatID:         message.Chat.ID,
		Sender:         tgUserToModel(message.From),
		Text:           message.Text,
		ReceiveTime:    time.Now(),
		IsPrivate:      message.Chat.IsPrivate(),
		IsButtonAction: false,
		ButtonActionID: model.INVALID_BUTTON_ACTION_ID,
	}
}

func tgUserToModel(user *tgbotapi.User) *model.User {
	// TODO: load by telegram id from repository
	return model.NewUser(
		user.ID,
		fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		user.UserName,
		"USER",
	)
}
