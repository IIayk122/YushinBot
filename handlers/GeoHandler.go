package handlers

import (
	"bufio"
	"bytes"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	tb "gopkg.in/tucnak/telebot.v2"
)

//GeoHandle Присылает инструкцию как добраться
func GeoHandle(b *tb.Bot, DB *mongo.Database) func(*tb.Message) {

	return func(msg *tb.Message) {
		b.Send(msg.Sender, "Как добраться")

		var buffer bytes.Buffer
		HowFindW := bufio.NewWriter(&buffer)
		bucket, err := gridfs.NewBucket(DB)
		if err != nil {
			log.Println(err)
		}
		_, err = bucket.DownloadToStreamByName("HowFind.jpg", HowFindW)
		if err != nil {
			log.Println(err)
		}
		HowFindR := bufio.NewReader(&buffer)

		b.Send(msg.Sender, &tb.Photo{
			File: tb.FromReader(HowFindR),
			Caption: "YUSHIN BROTHERS, Маркса 102а,\n" +
				"вывеска и вход со стороны Маркса, 4 этаж\n\n" +
				"Whatsapp, Telegram: +79233554321\n" +
				"Телефон: 21-543-21"})
	}
}
