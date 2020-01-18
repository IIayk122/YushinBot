package handlers

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	tb "gopkg.in/tucnak/telebot.v2"
)

//AddLostHandle добавляет потеряшку в базу
func AddLostHandle(b *tb.Bot, DB *mongo.Database) func(*tb.Message) {

	return func(msg *tb.Message) {
		if CheckAdmin(msg, DB) {
			WriteFlag := true
			b.Send(msg.Sender, "Пришли фото")
			b.Handle(tb.OnPhoto, func(m *tb.Message) {
				if WriteFlag {
					option := options.UploadOptions{}
					option.SetMetadata("потеряшка")
					file, err := b.GetFile(&m.Photo.File)
					if err != nil {
						fmt.Println(err)
					}
					bucket, err := gridfs.NewBucket(DB)
					if err != nil {
						log.Println(err)
					}
					err = bucket.UploadFromStreamWithID(primitive.NewObjectID(), m.Photo.FileID, file, options.GridFSUpload().SetMetadata(option))
					if err != nil {
						log.Println(err)
					}
					b.Send(m.Sender, "Добавил :3")

					WriteFlag = false
				}
			})

		} else {
			b.Send(msg.Sender, "Это тебе не доступно :3")
		}
	}
}
