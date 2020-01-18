package handlers

import (
	"NewYushinBot/keyboard"

	tb "gopkg.in/tucnak/telebot.v2"
)

//SecondVisitHandle присылает меню для старожил
func SecondVisitHandle(b *tb.Bot) func(*tb.Message) {
	return func(msg *tb.Message) {
		b.Send(msg.Sender, "Повторный визит", &tb.ReplyMarkup{ResizeReplyKeyboard: true, ReplyKeyboard: keyboard.SecondVisitMenu})
	}
}
