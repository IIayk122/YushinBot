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

//BarHandle Присылает меню бара из базы
func BarHandle(b *tb.Bot, DB *mongo.Database) func(*tb.Message) {

	return func(msg *tb.Message) {
		b.Notify(msg.Chat, telebot.UploadingPhoto)
		var buffer1 bytes.Buffer
		var buffer2 bytes.Buffer
		var buffer3 bytes.Buffer
		var buffer4 bytes.Buffer
		var buffer5 bytes.Buffer
		var buffer6 bytes.Buffer

		Bar1w := bufio.NewWriter(&buffer1)
		Bar1r := bufio.NewReader(&buffer1)
		Bar2w := bufio.NewWriter(&buffer2)
		Bar2r := bufio.NewReader(&buffer2)
		Bar3w := bufio.NewWriter(&buffer3)
		Bar3r := bufio.NewReader(&buffer3)
		Bar4w := bufio.NewWriter(&buffer4)
		Bar4r := bufio.NewReader(&buffer4)
		Bar5w := bufio.NewWriter(&buffer5)
		Bar5r := bufio.NewReader(&buffer5)
		Bar6w := bufio.NewWriter(&buffer6)
		Bar6r := bufio.NewReader(&buffer6)

		bucket, err := gridfs.NewBucket(DB)
		if err != nil {
			log.Println(err)
		}
		_, err = bucket.DownloadToStreamByName("1Bar.png", Bar1w)
		if err != nil {
			log.Println(err)
		}
		_, err = bucket.DownloadToStreamByName("2Bar.png", Bar2w)
		if err != nil {
			log.Println(err)
		}
		_, err = bucket.DownloadToStreamByName("3Bar.png", Bar3w)
		if err != nil {
			log.Println(err)
		}
		_, err = bucket.DownloadToStreamByName("4Bar.png", Bar4w)
		if err != nil {
			log.Println(err)
		}
		_, err = bucket.DownloadToStreamByName("5Bar.png", Bar5w)
		if err != nil {
			log.Println(err)
		}
		_, err = bucket.DownloadToStreamByName("6Bar.png", Bar6w)
		if err != nil {
			log.Println(err)
		}

		Bar1 := &tb.Photo{File: tb.FromReader(Bar1r)}
		Bar2 := &tb.Photo{File: tb.FromReader(Bar2r)}
		Bar3 := &tb.Photo{File: tb.FromReader(Bar3r)}
		Bar4 := &tb.Photo{File: tb.FromReader(Bar4r)}
		Bar5 := &tb.Photo{File: tb.FromReader(Bar5r)}
		Bar6 := &tb.Photo{File: tb.FromReader(Bar6r)}
		b.SendAlbum(msg.Sender, tb.Album{Bar1, Bar2, Bar3, Bar4, Bar5, Bar6})
	}
}
