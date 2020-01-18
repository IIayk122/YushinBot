package handlers

import (
	"bufio"
	"bytes"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	tb "gopkg.in/tucnak/telebot.v2"
)

type File struct {
	ID interface{} `bson:"_id"`
}

//LostHandle присылает список потеряшек из базы
func LostHandle(b *tb.Bot, DB *mongo.Database) func(*tb.Message) {
	return func(msg *tb.Message) {
		collection := DB.Collection("fs.files")
		options := options.Find()
		options.SetLimit(10)
		filter := bson.M{"metadata.metadata": "потеряшка"}
		var results []*File
		cur, err := collection.Find(context.TODO(), filter, options)
		if err != nil {
			log.Println(err)
		}
		for cur.Next(context.TODO()) {
			var elem File
			err := cur.Decode(&elem)
			if err != nil {
				log.Println(err)
			}
			results = append(results, &elem)
		}
		if err := cur.Err(); err != nil {
			log.Println(err)
		}
		cur.Close(context.TODO())

		album := tb.Album{}
		if results != nil {
			b.Send(msg.Sender, "Вот что у нас осталось:")
			for _, file := range results {
				var buffer bytes.Buffer
				forbearW := bufio.NewWriter(&buffer)
				bucket, err := gridfs.NewBucket(DB)
				if err != nil {
					log.Println(err)
				}
				_, err = bucket.DownloadToStream(file.ID, forbearW)
				if err != nil {
					log.Println(err)
				}
				forbearR := bufio.NewReader(&buffer)
				LostPhoto := &tb.Photo{File: tb.FromReader(forbearR)}
				album = append(album, LostPhoto)
			}
			_, err = b.SendAlbum(msg.Sender, album)
			if err != nil {
				log.Println(err)
			}
		} else {
			b.Send(msg.Sender, "Потеряшек нет 🥳")
		}
	}
}
