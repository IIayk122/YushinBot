package handlers

import (
	"bufio"
	"bytes"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	tb "gopkg.in/tucnak/telebot.v2"
)

//SmokeHandle Присылает меню кальянов из базы
func SmokeHandle(b *tb.Bot, DB *mongo.Database) func(*tb.Message) {

	return func(msg *tb.Message) {
		var buffer bytes.Buffer
		SmokeW := bufio.NewWriter(&buffer)
		bucket, err := gridfs.NewBucket(DB)
		if err != nil {
			log.Println(err)
		}
		_, err = bucket.DownloadToStreamByName("Smoke.jpg", SmokeW)
		if err != nil {
			log.Println(err)
		}
		SmokeR := bufio.NewReader(&buffer)

		b.Send(msg.Sender, &tb.Photo{File: tb.FromReader(SmokeR)})
	}
}
