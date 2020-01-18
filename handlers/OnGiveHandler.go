package handlers

import (
	"NewYushinBot/keyboard"
	tb "gopkg.in/tucnak/telebot.v2"
)

func OnGiveHandle(b *tb.Bot) func(*tb.Message) {
	return func(msg *tb.Message) {
		if keyboard.Give == false {
			b.Send(msg.Sender, "Розыгрыш включен")
			keyboard.Give = true
		} else {
			b.Send(msg.Sender, "Розыгрыш вЫключен")
			keyboard.Give = false
		}
	}
}
