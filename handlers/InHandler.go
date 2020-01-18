package handlers

import (
	"NewYushinBot/keyboard"

	tb "gopkg.in/tucnak/telebot.v2"
)

//InHandle присылает кнопки "Я уже тут"
func InHandle(b *tb.Bot) func(*tb.Message) {

	return func(msg *tb.Message) {
		b.Send(msg.Sender, "Я уже тут", &tb.ReplyMarkup{ResizeReplyKeyboard: true, ReplyKeyboard: keyboard.InMenu})
	}
}
