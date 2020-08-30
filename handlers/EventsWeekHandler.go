package handlers

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/tucnak/telebot.v2"
	tb "gopkg.in/tucnak/telebot.v2"
)

//EventsWeekHandle Присылает фотки из закрепленного поста в группе в вк
func EventsWeekHandle(b *tb.Bot, DB *mongo.Database) func(*tb.Message) {
	return func(msg *tb.Message) {
		api.AccessToken = ""
		b.Notify(msg.Chat, telebot.UploadingPhoto)
		b.Send(msg.Sender, "Мероприятия на неделю")
		album := tb.Album{}
		post := GetPinPost()
		for _, asd := range post.Response.Items[0].Attachments {
			tbphoto := &tb.Photo{File: tb.FromURL(asd.Photo.Photo2560)}
			album = append(album, tbphoto)
		}
		if album != nil {
			_, err := b.SendAlbum(msg.Sender, album)

			if err != nil {
				log.Println(err)
			}
		} else {
			_, err := b.Send(msg.Sender, "Пока незивестно 😔")

			if err != nil {
				log.Println(err)
			}
		}
	}
}
