package handlers

import (
	"bufio"
	"bytes"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	tb "gopkg.in/tucnak/telebot.v2"
)

//KitchenHandle Присылает меню кухни из базы
func KitchenHandle(b *tb.Bot, DB *mongo.Database) func(*tb.Message) {

	return func(msg *tb.Message) {
		var buffer bytes.Buffer
		forbearW := bufio.NewWriter(&buffer)
		bucket, err := gridfs.NewBucket(DB)
		if err != nil {
			log.Println(err)
		}
		_, err = bucket.DownloadToStreamByName("zakuski.jpg", forbearW)
		if err != nil {
			log.Println(err)
		}
		forbearR := bufio.NewReader(&buffer)
		//__________________________________________________________________________
		var buffer1 bytes.Buffer
		mainKitchenW := bufio.NewWriter(&buffer1)

		_, err = bucket.DownloadToStreamByName("osnovnoe_menyu.jpg", mainKitchenW)
		if err != nil {
			log.Println(err)
		}
		mainKitchenR := bufio.NewReader(&buffer1)
		zakuski := &tb.Photo{File: tb.FromReader(forbearR)}
		mainmenu := &tb.Photo{File: tb.FromReader(mainKitchenR)}
		b.SendAlbum(msg.Sender, tb.Album{zakuski, mainmenu})
	}

}
