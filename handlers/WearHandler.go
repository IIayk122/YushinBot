package handlers

import (
	"go.mongodb.org/mongo-driver/mongo"
	tb "gopkg.in/tucnak/telebot.v2"
)

//WearHandle Присылает мерч из базы
func WearHandle(b *tb.Bot, DB *mongo.Database) func(*tb.Message) {

	return func(msg *tb.Message) {
		b.Send(msg.Sender, "МЕРЧ......")
	}
}
