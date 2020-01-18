package handlers

import (
	"NewYushinBot/keyboard"

	tb "gopkg.in/tucnak/telebot.v2"
)

//FirstVisitHandle Присылает меню для новичка
func FirstVisitHandle(b *tb.Bot) func(*tb.Message) {
	return func(msg *tb.Message) {
		b.Send(msg.Sender, "Первый визит", &tb.ReplyMarkup{ResizeReplyKeyboard: true, ReplyKeyboard: keyboard.FirstVisitMenu})
	}
}
