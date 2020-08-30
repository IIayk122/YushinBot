package handlers

import (

	//"github.com/tuotoo/qrcode"

	"NewYushinBot/keyboard"
	"context"
	"image"
	_ "image/jpeg"
	"log"
	"time"

	"github.com/liyue201/goqr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Gift struct {
	ID   interface{} `bson:"_id"`
	Gift string      `bson:"gift"`
}

type Winners struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	Username  string    `json:"username"`
	Time      time.Time `json:"date"`
	Gift      string    `json:"gift"`
}

//QREventHandle присылает список потеряшек из базы
func QREventHandle(b *tb.Bot, DB *mongo.Database) func(*tb.Message) {
	return func(msg *tb.Message) {
		if keyboard.Give {

			WriteFlag := true
			b.Send(msg.Sender, "Пришлите фото")
			b.Handle(tb.OnPhoto, func(m *tb.Message) {
				if WriteFlag {
					fl, err := b.GetFile(&m.Photo.File)
					if err != nil {
						b.Send(m.Sender, "Файл до меня не дошел 😔")
					}
					img, _, err := image.Decode(fl)
					if err != nil {
						b.Send(m.Sender, "Не распознано 😔")
						log.Println(err)
					}
					qrCodes, err := goqr.Recognize(img)
					if err != nil || len(qrCodes) == 0 { //на случай плохого QR кода, библеотека инногда возвращает пустой массив без ошибки
						b.Send(m.Sender, "Не распознано 😔")
					}
					for _, qrCode := range qrCodes {

						var gift Gift
						collection := DB.Collection("Gifts")

						objID, _ := primitive.ObjectIDFromHex(string(qrCode.Payload))
						filter := bson.M{"_id": objID}
						err := collection.FindOne(context.TODO(), filter).Decode(&gift) //поиск по ID приза
						if err != nil {
							log.Println(err)
							b.Send(m.Sender, "Этот приз уже забрали")
						} else {

							winner := Winners{
								ID:        m.Sender.ID,
								FirstName: m.Sender.FirstName,
								Username:  m.Sender.Username,
								Time:      m.Time().Add(time.Hour * 7),
								Gift:      gift.Gift}
							userCollection := DB.Collection("Winners")
							_, err := userCollection.InsertOne(context.TODO(), winner) // добавление победителя в базу
							if err != nil {
								log.Println(err)
							}

							_, err = b.Send(m.Sender,
								"Тысяча поздравлений!!!!\n"+
									"🥳🥳🥳🥳🥳🥳🥳\n"+
									"Твои выигрыш: \n"+
									gift.Gift+
									"\n\nПиши @borozdaa, что бы получить его!!!") //отправка приза победителю
							if err != nil {
								log.Println(err)
							}
							_, err = b.Send(&tb.User{ID: 175785539},
								"Розыгрыш: \n"+
									m.Sender.FirstName+" выиграл:\n"+
									gift.Gift)
							if err != nil {
								log.Println(err)
							}
							_, err = collection.DeleteOne(context.TODO(), filter) //удаление приза после обнаружения QR кода
							if err != nil {
								log.Println(err)
							}
						}
					}
				}
				WriteFlag = false
			})

		} else {
			b.Send(msg.Sender, "🤗")
		}
	}
}
