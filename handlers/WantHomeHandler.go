package handlers

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

//WantHomeHandle ссылку на яндекс такси
func WantHomeHandle(b *tb.Bot) func(*tb.Message) {

	return func(msg *tb.Message) {

		inlineBtn := tb.InlineButton{
			Unique: "GoHome",
			Text:   "Нажми меня",
			URL:    "vk.cc/a9Ct6i",
		}
		inlineKeys := [][]tb.InlineButton{
			[]tb.InlineButton{inlineBtn},
		}

		b.Send(msg.Sender, "Твой билет домой путник :3\n", &tb.ReplyMarkup{
			InlineKeyboard: inlineKeys,
		})

	}
}
