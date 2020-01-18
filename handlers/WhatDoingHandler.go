package handlers

import (
	"NewYushinBot/keyboard"

	tb "gopkg.in/tucnak/telebot.v2"
)

//WhatDoingHandle присылает кнопки "Что у вас"
func WhatDoingHandle(b *tb.Bot) func(*tb.Message) {

	return func(msg *tb.Message) {
		b.Send(msg.Sender, "Что у вас делать?", &tb.ReplyMarkup{ResizeReplyKeyboard: true, ReplyKeyboard: keyboard.WhatDoingMenu})
	}
}
