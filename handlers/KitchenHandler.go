package handlers

import (
	"bytes"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"gopkg.in/tucnak/telebot.v2"
	tb "gopkg.in/tucnak/telebot.v2"
)

//KitchenHandle Присылает меню кухни из базы
func KitchenHandle(b *tb.Bot, DB *mongo.Database) func(*tb.Message) {

	return func(msg *tb.Message) {
		b.Notify(msg.Chat, telebot.UploadingPhoto)

		bucket, err := gridfs.NewBucket(DB)
		if err != nil {
			log.Println(err)
		}

		var buffer1 bytes.Buffer
		_, err = bucket.DownloadToStreamByName("osnovnoe_menyu.jpg", &buffer1)
		if err != nil {
			log.Println(err)
		}
		_, err = b.Send(msg.Sender, &tb.Photo{File: tb.FromReader(&buffer1)})
		if err != nil {
			log.Println(err)
		}
		// mainmenu := &tb.Photo{File: tb.FromReader(&buffer1)}
		// b.SendAlbum(msg.Sender, tb.Album{mainmenu})

	}

}
