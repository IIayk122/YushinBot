package handlers

import (
	"NewYushinBot/keyboard"

	tb "gopkg.in/tucnak/telebot.v2"
)

func WantClubHandle(b *tb.Bot) func(*tb.Message) {
	return func(msg *tb.Message) {
		b.Send(msg.Sender,
			"Сколько у тебя есть шекелей?",
			&tb.ReplyMarkup{
				ResizeReplyKeyboard: true,
				ReplyKeyboard:       keyboard.ClubMenu})
	}
}
