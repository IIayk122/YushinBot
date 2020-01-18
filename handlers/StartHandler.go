package handlers

import (
	"NewYushinBot/keyboard"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	tb "gopkg.in/tucnak/telebot.v2"
)

//StartHandle приветствует нового пользователя и записывает его в базу если он уникальный.
func StartHandle(b *tb.Bot, DB *mongo.Database) func(*tb.Message) {
	return func(msg *tb.Message) {
		collection := DB.Collection("Users")
		_, err := collection.InsertOne(context.TODO(), msg.Sender)
		if err != nil {
			log.Print("Дубль:\n", err)
		}
		if keyboard.Give == false {
			b.Send(msg.Sender, "Привет "+msg.Sender.FirstName+"!\n"+
				"Добро пожаловать в YushinBot", &tb.ReplyMarkup{
				ResizeReplyKeyboard: true,
				ReplyKeyboard:       keyboard.MainMenu})
		} else {
			b.Send(msg.Sender, "Привет "+msg.Sender.FirstName+"!\n"+
				"Добро пожаловать в YushinBot", &tb.ReplyMarkup{
				ResizeReplyKeyboard: true,
				ReplyKeyboard:       keyboard.MainMenuGive})
		}
	}
}
