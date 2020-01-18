package handlers

import (
	"go.mongodb.org/mongo-driver/mongo"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
)

//EventsWeekHandle Присылает фотки из закрепленного поста в группе в вк
func EventsWeekHandle(b *tb.Bot, DB *mongo.Database) func(*tb.Message) {
	return func(msg *tb.Message) {
		api.AccessToken = "23fd9e9323fd9e9323fd9e93932395e535223fd23fd9e937fa4acc326ca587052f84589"
		b.Send(msg.Sender, "Мероприятия на неделю")
		album := tb.Album{}
		post := GetPinPost()
		for _, asd := range post.Response.Items[0].Attachments {
			tbphoto := &tb.Photo{File: tb.FromURL(asd.Photo.Photo2560)}
			album = append(album, tbphoto)
		}
		if album != nil {
			b.SendAlbum(msg.Sender, album)
			log.Println(album)
		} else {
			b.Send(msg.Sender, "Пока незивестно 😔")
		}
	}
}
