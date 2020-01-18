package handlers

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	tb "gopkg.in/tucnak/telebot.v2"
)

type News struct {
	New string `json:"new"`
}

//AddNewsHandle добавляет новинку в базу
func AddNewsHandle(b *tb.Bot, DB *mongo.Database) func(*tb.Message) {
	return func(msg *tb.Message) {
		WriteFlag := true
		if CheckAdmin(msg, DB) {
			b.Send(msg.Sender, "Введите текст новости")
			b.Handle(tb.OnText, func(m *tb.Message) {
				if WriteFlag {
					news := News{New: m.Text}
					collection := DB.Collection("News")
					_, err := collection.InsertOne(context.TODO(), news)
					if err != nil {
						log.Println("Дубль:\n", err)
					}
					WriteFlag = false
					b.Send(msg.Sender, "Добавлено")
				}
			})
		} else {
			b.Send(msg.Sender, "Это тебе не доступно :3")
		}
	}
}
