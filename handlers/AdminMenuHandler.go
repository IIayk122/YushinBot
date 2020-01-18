package handlers

import (
	"NewYushinBot/keyboard"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	tb "gopkg.in/tucnak/telebot.v2"
)

//AdminMenuHandle проверяет есть ли пользователь в базе админов и в случае успеха, открывает АдминМеню
func AdminMenuHandle(b *tb.Bot, DB *mongo.Database) func(*tb.Message) {

	return func(msg *tb.Message) {
		if CheckAdmin(msg, DB) {
			b.Send(msg.Sender, "Привет Админ))))", &tb.ReplyMarkup{ResizeReplyKeyboard: true, ReplyKeyboard: keyboard.AdminMenu})
		} else {
			b.Send(msg.Sender, "Это тебе не доступно :3")
		}
	}
}

//CheckAdmin проверяет есть ли отправитель в коллекции админов в базе
func CheckAdmin(msg *tb.Message, DB *mongo.Database) bool {
	var User tb.User
	collection := DB.Collection("Admins")
	filter := bson.D{primitive.E{Key: "id", Value: msg.Sender.ID}}
	err := collection.FindOne(context.TODO(), filter).Decode(&User)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
