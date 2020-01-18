package handlers

import (
	"NewYushinBot/vkapi"
	"encoding/json"
	"log"
	"strconv"

	tb "gopkg.in/tucnak/telebot.v2"
)

var api = &vkapi.Api{}

//PhotosHandle ..
func PhotosHandle(b *tb.Bot) func(*tb.Message) {
	return func(msg *tb.Message) {
		api.AccessToken = "23fd9e9323fd9e9323fd9e93932395e535223fd23fd9e937fa4acc326ca587052f84589"
		b.Send(msg.Sender, "Последние фотоотчеты:")
		a := GetAlbums()
		for _, row := range a.Response.Items {
			d := GetPhoto(row.ID, row.ThumbID)
			b.Send(msg.Sender,
				&tb.Photo{
					File:    tb.FromURL(d.Response.Items[0].Sizes[4].URL),
					Width:   1280,
					Height:  853,
					Caption: row.Title + "\nhttps://vk.com/album-149577615_" + strconv.Itoa(row.ID)})
		}

	}
}

//GetAlbums ..
func GetAlbums() vkapi.Album {
	getalbums := api.Request("photos.getAlbums", map[string]string{
		"owner_id":  "-149577615",
		"album_ids": "",
		"v":         "5.30",
		"count":     "5"})
	b := []byte(getalbums)
	var app = vkapi.Album{}
	err1 := json.Unmarshal(b, &app)
	if err1 != nil {
		log.Println(err1)
	}
	return app
}

//GetPinPost ...
func GetPinPost() vkapi.Wall {
	getWall := api.Request("wall.get", map[string]string{
		"owner_id": "-149577615",
		"domain":   "",
		"offset":   "",
		"count":    "1",
		"filter":   "owner",
		"v":        "5.60"})
	b := []byte(getWall)
	var wall = vkapi.Wall{}
	err1 := json.Unmarshal(b, &wall)
	if err1 != nil {
		log.Println(err1)
	}
	return wall

}

//GetPhoto ...
func GetPhoto(albumID int, thumbID int) vkapi.Photo {
	GetPhoto := api.Request("photos.get", map[string]string{
		"owner_id":  "-149577615",
		"album_id":  strconv.Itoa(albumID),
		"photo_ids": strconv.Itoa(thumbID),
		"count":     "1",
		"rev":       "1",
		"v":         "5.102"})
	b := []byte(GetPhoto)
	var photo = vkapi.Photo{}
	err1 := json.Unmarshal(b, &photo)
	if err1 != nil {
		log.Println(err1)
	}
	return photo

}
