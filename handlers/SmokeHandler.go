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

//SmokeHandle Присылает меню кальянов из базы
func SmokeHandle(b *tb.Bot, DB *mongo.Database) func(*tb.Message) {
	return func(msg *tb.Message) {
		b.Notify(msg.Chat, telebot.UploadingPhoto)
		var buffer1 bytes.Buffer
		var buffer2 bytes.Buffer
		var buffer3 bytes.Buffer

		Smoke1w := bufio.NewWriter(&buffer1)
		Smoke1r := bufio.NewReader(&buffer1)
		Smoke2w := bufio.NewWriter(&buffer2)
		Smoke2r := bufio.NewReader(&buffer2)
		Smoke3w := bufio.NewWriter(&buffer3)
		Smoke3r := bufio.NewReader(&buffer3)

		bucket, err := gridfs.NewBucket(DB)
		if err != nil {
			log.Println(err)
		}
		_, err = bucket.DownloadToStreamByName("1Smoke.jpeg", Smoke1w)
		if err != nil {
			log.Println(err)
		}
		_, err = bucket.DownloadToStreamByName("2Smoke.jpeg", Smoke2w)
		if err != nil {
			log.Println(err)
		}
		_, err = bucket.DownloadToStreamByName("3Smoke.jpeg", Smoke3w)
		if err != nil {
			log.Println(err)
		}

		smoke1 := &tb.Photo{File: tb.FromReader(Smoke1r)}
		smoke2 := &tb.Photo{File: tb.FromReader(Smoke2r)}
		smoke3 := &tb.Photo{File: tb.FromReader(Smoke3r)}
		b.SendAlbum(msg.Sender, tb.Album{smoke1, smoke2, smoke3})
	}
}
