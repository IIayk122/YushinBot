package handlers

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

//RecordHandle записывает на стрижку
func RecordHandle(b *tb.Bot) func(*tb.Message) {
	return func(msg *tb.Message) {
		b.Send(msg.Sender, "https://vk.cc/9Tbfwv")
	}
}
