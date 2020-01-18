package handlers

import (
	"fmt"

	tb "gopkg.in/tucnak/telebot.v2"
)

//RecordHandle записывает на стрижку
func RecordHandle(b *tb.Bot) func(*tb.Message) {
	return func(msg *tb.Message) {
		b.Send(msg.Sender, "Запись на стрижку")
		b.Handle(tb.OnContact, func(m *tb.Message) {
			fmt.Println(m.Contact)
		})


	}
}
