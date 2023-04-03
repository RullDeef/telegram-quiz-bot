package tginteractor

import (
	"fmt"
	"log"

	"github.com/RullDeef/telegram-quiz-bot/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TGBotInteractor struct {
	publisher *TGBotPublisher
	chatID    int64
	msgChan   chan model.Message
}

func NewInteractor(
	publisher *TGBotPublisher,
	chatID int64,
	msgChan chan model.Message,
) *TGBotInteractor {
	return &TGBotInteractor{
		publisher: publisher,
		chatID:    chatID,
		msgChan:   msgChan,
	}
}

func (tgi *TGBotInteractor) MessageChan() chan model.Message {
	return tgi.msgChan
}

func (tgi *TGBotInteractor) SendResponse(response model.Response) {
	msg := tgbotapi.NewMessage(tgi.chatID, response.Text)

	// add inline keyboard for each action in response
	if len(response.Actions) > 0 {
		var keyboardRows []tgbotapi.InlineKeyboardButton

		for _, action := range response.Actions {
			data := fmt.Sprintf("%d", action.ID)
			button := tgbotapi.NewInlineKeyboardButtonData(action.Text, data)
			keyboardRow := tgbotapi.NewInlineKeyboardRow(button)
			keyboardRows = append(keyboardRows, keyboardRow...)
		}

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(keyboardRows)
		_, err := tgi.publisher.bot.Send(msg)
		if err != nil {
			log.Printf("ERROR: failed to send message with markup: \"%s\"", msg.Text)
		}
	} else {
		_, err := tgi.publisher.bot.Send(msg)
		if err != nil {
			log.Printf("ERROR: failed to send message: \"%s\"", msg.Text)
		}
	}
}
