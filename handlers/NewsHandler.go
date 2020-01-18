package handlers

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	tb "gopkg.in/tucnak/telebot.v2"
)

//NewsHandle Присылает последнии 2 новости из базы
func NewsHandle(b *tb.Bot, DB *mongo.Database) func(*tb.Message) {

	return func(msg *tb.Message) {
		collection := DB.Collection("News")
		options := options.Find()
		options.SetLimit(2)
		filter := bson.M{}
		var results []*News
		cur, err := collection.Find(context.TODO(), filter)
		if err != nil {
			log.Println(err)
		}
		for cur.Next(context.TODO()) {
			var elem News
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
		//____________________________________________________________
		if results != nil {
			for _, news := range results {
				b.Send(msg.Sender, news.New)
			}
		} else {
			b.Send(msg.Sender, "В разработке 😌")
		}
	}
}
