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

//MapYushinHandle Присылает карту пространства из базы
func MapYushinHandle(b *tb.Bot, DB *mongo.Database) func(*tb.Message) {

	return func(msg *tb.Message) {
		b.Notify(msg.Chat, telebot.UploadingPhoto)
		var buffer bytes.Buffer
		mapWriter := bufio.NewWriter(&buffer)
		bucket, err := gridfs.NewBucket(DB)
		if err != nil {
			log.Println(err)
		}
		_, err = bucket.DownloadToStreamByName("Map.jpg", mapWriter)
		if err != nil {
			log.Println(err)
		}
		mapReader := bufio.NewReader(&buffer)

		b.Send(msg.Sender, &tb.Photo{File: tb.FromReader(mapReader)})
	}

}
