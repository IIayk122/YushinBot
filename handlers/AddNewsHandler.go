package handlers

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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
						log.Println(err)
					}
					WriteFlag = false
					b.Send(msg.Sender, "Добавлено")

					collection = DB.Collection("Users")

					res, err := collection.Distinct(context.TODO(), "id", bson.M{})
					if err != nil {
						log.Println(err)
					}
					log.Println(res)
					idList := []int32{}

					for _, item := range res {
						idList = append(idList, item.(int32))
					}

					for _, id := range idList {
						_, err = b.Send(
							&tb.User{
								ID: int(id)},
							"Привет!🤪\n"+
								"У нас для тебя новость!\n\n\n"+
								m.Text)
						if err != nil {
							if err.Error() == "api error: Forbidden: bot was blocked by the user" {
								_, _ = collection.DeleteOne(context.TODO(), bson.M{"id": id})
							} else {
								log.Println(err)
							}
						}
					}
				}
			})
		} else {
			b.Send(msg.Sender, "Это тебе не доступно :3")
		}
	}
}
