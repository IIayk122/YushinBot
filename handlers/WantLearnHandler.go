package handlers

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

//WantLearnHandle Присылает контакты отвественного за мастер классы
func WantLearnHandle(b *tb.Bot) func(*tb.Message) {

	return func(msg *tb.Message) {
		b.Send(msg.Sender,
			"Пиши ему, он всё раскажет 🤓\n\n"+
				"Паша Бороздов\n"+
				"+7 (923) 376-46-33\n"+
				"@borozdaaa")
	}
}
