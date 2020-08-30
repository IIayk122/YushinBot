package handlers

import (
	"bufio"
	"bytes"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"gopkg.in/tucnak/telebot.v2"
	tb "gopkg.in/tucnak/telebot.v2"
)

//HairCutsHandle Присылает расценки стрижек из базы
func HairCutsHandle(b *tb.Bot, DB *mongo.Database) func(*tb.Message) {

	return func(msg *tb.Message) {
		b.Notify(msg.Chat, telebot.UploadingPhoto)

		var buffer bytes.Buffer
		barberW := bufio.NewWriter(&buffer)
		bucket, err := gridfs.NewBucket(DB)
		if err != nil {
			log.Println(err)
		}
		_, err = bucket.DownloadToStreamByName("barber2.jpg", barberW)
		if err != nil {
			log.Println(err)
		}
		barberR := bufio.NewReader(&buffer)

		b.Send(msg.Sender, &tb.Photo{File: tb.FromReader(barberR)})
	}
}
