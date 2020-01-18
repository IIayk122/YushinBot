package handlers

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/now"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	tb "gopkg.in/tucnak/telebot.v2"
)

//Event - события лежащие в базе
type Event struct {
	Name        string
	Discription string
	Date        time.Time
	Link        string
}

//EventDayHandle Присылает из базы список мероприятий на день
func EventDayHandle(b *tb.Bot, DB *mongo.Database) func(*tb.Message) {
	return func(msg *tb.Message) {
		b.Send(msg.Sender, "Сейчас посмотрим :3")
		ot := now.BeginningOfDay().UTC().Add(time.Hour * 18)           //11:00 (+7)
		do := now.EndOfDay().UTC().Add(time.Hour*12 + time.Nanosecond) //05:00 (+7)
		collection := DB.Collection("Events")
		filter := bson.M{"date": bson.M{"$gte": ot, "$lte": do}}
		var results []*Event
		cur, err := collection.Find(context.TODO(), filter)
		if err != nil {
			log.Println(err)
		}

		for cur.Next(context.TODO()) {
			var elem Event
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
		//______________________________________________________________________________________________________

		inlineKeys := [][]tb.InlineButton{[]tb.InlineButton{}}

		if results != nil {
			for _, event := range results {
				count := 0

				inlintbtn := createBtn(b, strconv.Itoa(count), "Напомнить за 2 часа")
				inlineKeys = [][]tb.InlineButton{{inlintbtn}}

				b.Handle(&inlineKeys[0][0], func(c *tb.Callback) {
					date := strings.SplitAfter(c.Message.Caption, "\n")[len(strings.SplitAfter(c.Message.Caption, "\n"))-1]
					err := notification(date, b, c)
					b.Edit(c.Message, "")
					if err != nil {
						log.Printf("error: %v\n", err)
					}
					b.Respond(c, &tb.CallbackResponse{Text: "Организуем  😉"})
				})

				b.Send(msg.Sender, &tb.Photo{
					File:    tb.FromURL(event.Link),
					Caption: event.Name + "\n\n" + event.Discription + "\n" + event.Date.Add(time.Millisecond).Format("2006-01-02 15:04")},
					&tb.ReplyMarkup{InlineKeyboard: inlineKeys})
				count++

			}
			fmt.Println(inlineKeys)

		} else {
			b.Send(msg.Sender, "На сегодня нет мероприятий")
		}

	}

}

func createBtn(b *tb.Bot, uniqueName, text string) tb.InlineButton {
	btn := tb.InlineButton{
		Unique: uniqueName,
		Text:   text,
	}
	return btn
}

func notification(eventTime string, b *tb.Bot, c *tb.Callback) error {
	EventName := strings.SplitAfter(c.Message.Caption, "\n")[0]
	// Разбираем время запуска.
	eTime, err := time.Parse("2006-01-02 15:04", eventTime)
	if err != nil {
		return err
	}
	// Вычисляем временной промежуток до запуска.
	duration := eTime.Sub(time.Now().UTC().Add(time.Hour * 9))
	if duration > 0 {
		b.EditReplyMarkup(c.Message, &tb.ReplyMarkup{})
		b.Send(c.Sender, "Не проблема 😉")
		log.Println(duration)

		go func(b *tb.Bot, name string, c *tb.User) {
			time.Sleep(duration)
			b.Send(c, "Привет :3 \nНе забудь, через 2 часа начнется:\n"+name)
		}(b, EventName, c.Sender)
	} else {
		b.Send(c.Sender, `Об этом уже поздно напоминать`+"\n"+`¯\_(ツ)_/¯`)
	}

	return nil
}
