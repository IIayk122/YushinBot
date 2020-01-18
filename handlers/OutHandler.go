package handlers

import (
	"NewYushinBot/keyboard"

	tb "gopkg.in/tucnak/telebot.v2"
)

//OutHandle присылает кнопки "Я погнал дальше"
func OutHandle(b *tb.Bot) func(*tb.Message) {

	return func(msg *tb.Message) {
		b.Send(msg.Sender, "Я погнал дальше", &tb.ReplyMarkup{ResizeReplyKeyboard: true, ReplyKeyboard: keyboard.OutMenu})
	}
}
