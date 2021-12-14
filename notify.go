package logging

import (
	"context"

	"github.com/bot-api/telegram"
)

func SendMessage(MessageText string, ChatID int64) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//Chat ID is the chat to send the message to
	msg := telegram.NewMessage(ChatID, MessageText)
	//Send Message
	Logger.TelegramAPI.Send(ctx, msg)

}