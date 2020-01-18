package handlers

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

//WantSongHandle присылает контакты отвественного за выступления
func WantSongHandle(b *tb.Bot) func(*tb.Message) {

	return func(msg *tb.Message) {
		b.Send(msg.Sender,
			"Пиши ему, он всё раскажет 🤩\n\n"+
				"Саша Овчаренко\n"+
				"+7 (913) 188-65-87\n"+
				"@aleksandr_ovch")
	}
}
