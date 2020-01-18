package handlers

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Comments struct {
	Comment   string    `json:"comment"`
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	Username  string    `json:"username"`
	Time      time.Time `json:"date"`
}

//CommentHandle записывает в базу отзыв/предложение
func CommentHandle(b *tb.Bot, DB *mongo.Database) func(*tb.Message) {
	return func(msg *tb.Message) {
		WriteFlag := true
		if CheckAdmin(msg, DB) {
			b.Send(msg.Sender, "С нетерпением ждем 😊")
			b.Handle(tb.OnText, func(m *tb.Message) {
				if WriteFlag {
					comment := Comments{
						Comment:   m.Text,
						ID:        m.Sender.ID,
						FirstName: m.Sender.FirstName,
						Username:  m.Sender.Username,
						Time:      m.Time().Add(time.Hour * 7)}

					collection := DB.Collection("Comments")
					_, err := collection.InsertOne(context.TODO(), comment)
					if err != nil {
						log.Println("Дубль:\n", err)
					}
					WriteFlag = false
					b.Send(msg.Sender, "Спасибо за потраченное время 😘")
				}
			})
		} else {
			b.Send(msg.Sender, "Это тебе не доступно :3")
		}
	}
}
