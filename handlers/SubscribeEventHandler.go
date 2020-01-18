package handlers

import (
	"go.mongodb.org/mongo-driver/mongo"
	tb "gopkg.in/tucnak/telebot.v2"
)

//SubscribeEventHandle записывает в базу тех кому присылать увдомления
func SubscribeEventHandle(b *tb.Bot, DB *mongo.Database) func(*tb.Message) {

	return func(msg *tb.Message) {
		b.Send(msg.Sender, "Subscirbe")
	}
}
