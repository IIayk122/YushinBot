package handlers

import (
	"NewYushinBot/keyboard"

	"go.mongodb.org/mongo-driver/mongo"
	tb "gopkg.in/tucnak/telebot.v2"
)

//YushinMenuBtnHandle присылает кнопки с меню заведения
func YushinMenuBtnHandle(b *tb.Bot, DB *mongo.Database) func(*tb.Message) {
	return func(msg *tb.Message) {
		b.Send(msg.Sender, "Меню", &tb.ReplyMarkup{ResizeReplyKeyboard: true, ReplyKeyboard: keyboard.YushinMenu})
	}
}
