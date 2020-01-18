package handlers

import (
	"context"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	tb "gopkg.in/tucnak/telebot.v2"
)

//ViewSubsHandle выводить колличество записей в таблице Users
func ViewSubsHandle(b *tb.Bot, DB *mongo.Database) func(*tb.Message) {

	return func(msg *tb.Message) {
		if CheckAdmin(msg, DB) {
			coll := DB.Collection("Users")
			opts := options.Count().SetMaxTime(2 * time.Second)
			count, err := coll.CountDocuments(context.TODO(), bson.D{{}}, opts)
			if err != nil {
				log.Println(err)
			}
			b.Send(msg.Sender, "Колличество подписчиков: "+strconv.Itoa(int(count)))
		} else {
			b.Send(msg.Sender, "Это тебе не доступно :3")
		}
	}
}
