package handlers

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	tb "gopkg.in/tucnak/telebot.v2"
)

//ViewCommentHandle подгружает отзывы из базы
func ViewCommentHandle(b *tb.Bot, DB *mongo.Database) func(*tb.Message) {

	return func(msg *tb.Message) {
		if CheckAdmin(msg, DB) {
			b.Send(msg.Sender, "Отзывы:")
			collection := DB.Collection("Comments")
			options := options.Find()
			options.SetLimit(10)
			filter := bson.M{}
			var results []*Comments
			cur, err := collection.Find(context.TODO(), filter, options)
			if err != nil {
				log.Println(err)
			}
			for cur.Next(context.TODO()) {
				var elem Comments
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
			if results != nil {
				for _, comment := range results {
					b.Send(msg.Sender, "Кто: "+comment.FirstName+
						"\nЛогин: @"+comment.Username+
						"\nКогда: "+comment.Time.Format("02.01.2006 в 15:04")+
						"\n👇👇👇👇👇👇👇👇👇👇 "+
						"\n"+comment.Comment)
				}
			}

		} else {
			b.Send(msg.Sender, "Это тебе не доступно :3")
		}
	}
}
