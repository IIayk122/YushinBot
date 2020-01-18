package handlers

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/tealeg/xlsx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	tb "gopkg.in/tucnak/telebot.v2"
)

type event struct {
	Name        string    `xlsx:"0"`
	Discription string    `xlsx:"1"`
	Date        time.Time `xlsx:"2"`
	Link        string    `xlsx:"3"`
}

//AddEventHandle Принимает файлы с событиями от пользователя парсит и добавляет в базу
func AddEventHandle(b *tb.Bot, DB *mongo.Database) func(*tb.Message) {
	return func(msg *tb.Message) {
		var i int
		if CheckAdmin(msg, DB) {
			b.Send(msg.Sender, "Отправте файл")
			b.Handle(tb.OnDocument, func(m *tb.Message) {
				err := b.Download(&m.Document.File, m.Document.FileName)
				if err != nil {
					log.Println(err)
				} else {
					b.Send(msg.Sender, "принял :3")

					excelFileName := m.Document.FileName
					xlFile, err := xlsx.OpenFile(excelFileName)
					if err != nil {
						log.Println(err)
					}
					collection := DB.Collection("Events")
					_, err = collection.DeleteMany(context.TODO(), bson.D{{}})
					if err != nil {
						log.Println(err)
					}

					ololo := &event{}
					i = 0
					for _, sheet := range xlFile.Sheets {
						for _, row := range sheet.Rows {

							err := row.ReadStruct(ololo)
							if err != nil {
								log.Println(err)
								b.Send(msg.Sender, "Что то пошло не так при чтении файла")
								break
							} else {
								_, err = collection.InsertOne(context.TODO(), ololo)
								if err != nil {
									b.Send(msg.Sender, "Что то пошло не так при записи в базу")
									log.Println(err)
									break
								} else {
									i++
								}
							}
						}
					}
					b.Send(msg.Sender, "Добавленно записей: "+strconv.Itoa(i))
					err = os.Remove(excelFileName)
					if err != nil {
						log.Println(err)
					}
				}
			})
		} else {
			b.Send(msg.Sender, "Это тебе не доступно :3")
		}
	}
}
