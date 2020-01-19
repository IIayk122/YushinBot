package main

import (
	"NewYushinBot/handlers"
	"NewYushinBot/keyboard"
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/proxy"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	tokenBot         = "******:*****_*****************************"
	connectionString = "mongodb://user:pass@localhost:27017"
	nameDB           = "YushinDB"
)

func main() {
	//---------------------------------------------------------------------------------------------------
	dialer, err := proxy.SOCKS5("tcp", "80.211.29.222:8975", &proxy.Auth{User: "user", Password: "pass"}, proxy.Direct) // 54.37.18.209:58184
	if err != nil {
		log.Fatal(err)
	}

	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}
	httpTransport.Dial = dialer.Dial
	poller := &tb.LongPoller{Timeout: 15 * time.Second}
	spamProtected := tb.NewMiddlewarePoller(poller, func(upd *tb.Update) bool {
		if upd.Message == nil {
			return true
		}
		if strings.Contains(upd.Message.Text, "spam") {
			return false
		}
		return true
	})
	b, err := tb.NewBot(tb.Settings{
		Token:  tokenBot,
		Poller: spamProtected,
		Client: httpClient,
	})
	if err != nil {
		log.Fatal(err)

	}

	//---------------------------------------------------------------------------------------------------
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())
	DB := client.Database(nameDB)

	log.Println("Connected to MongoDB!")

	//---------------------------------------------------------------------------------------------------КНОПКА ГЛАВНОГО МЕНЮ
	b.Handle(&keyboard.MainMenuBtn, MainMenuHandle(b))
	//---------------------------------------------------------------------------------------------------ГЛАВНОЕ МЕНЮ
	b.Handle("/start", handlers.StartHandle(b, DB))
	b.Handle(&keyboard.Record, handlers.RecordHandle(b))
	b.Handle(&keyboard.EventsDay, handlers.EventDayHandle(b, DB))
	b.Handle(&keyboard.YushinMenuBtn, handlers.YushinMenuBtnHandle(b, DB))
	b.Handle(&keyboard.FirstVisit, handlers.FirstVisitHandle(b))
	b.Handle(&keyboard.SecondVisit, handlers.SecondVisitHandle(b))
	//---------------------------------------------------------------------------------------------------Yushin МЕНЮ ЗАВЕДЕНИЯ
	b.Handle(&keyboard.HairCuts, handlers.HairCutsHandle(b, DB))
	b.Handle(&keyboard.Wear, handlers.WearHandle(b, DB)) //Мерч - не сделано
	b.Handle(&keyboard.Kitchen, handlers.KitchenHandle(b, DB))
	b.Handle(&keyboard.Smoke, handlers.SmokeHandle(b, DB))
	b.Handle(&keyboard.Bar, handlers.BarHandle(b, DB)) //Карта бара -- не сделана
	//---------------------------------------------------------------------------------------------------Я СТАРОЖИЛ
	b.Handle(&keyboard.SubscribeEvent, handlers.SubscribeEventHandle(b, DB))
	b.Handle(&keyboard.Comment, handlers.CommentHandle(b, DB))
	b.Handle(&keyboard.WantSong, handlers.WantSongHandle(b))
	b.Handle(&keyboard.WantLearn, handlers.WantLearnHandle(b))
	b.Handle(&keyboard.Photos, handlers.PhotosHandle(b))
	b.Handle(&keyboard.News, handlers.NewsHandle(b, DB))
	b.Handle(&keyboard.Lost, handlers.LostHandle(b, DB))
	//---------------------------------------------------------------------------------------------------Я НОВИЧОК
	b.Handle(&keyboard.Out, handlers.OutHandle(b))
	b.Handle(&keyboard.In, handlers.InHandle(b))               //я уже тут
	b.Handle(&keyboard.WhatDoing, handlers.WhatDoingHandle(b)) // Что у вас делать?
	//---------------------------------------------------------------------------------------------------Что меня у вас делать?
	b.Handle(&keyboard.MapYushin, handlers.MapYushinHandle(b, DB))
	b.Handle(&keyboard.EventsWeek, handlers.EventsWeekHandle(b, DB))
	b.Handle(&keyboard.Geo, handlers.GeoHandle(b, DB))
	//---------------------------------------------------------------------------------------------------Я погнал дальше
	b.Handle(&keyboard.WantHome, handlers.WantHomeHandle(b))
	b.Handle(&keyboard.WantClub, handlers.WantClubHandle(b))
	//---------------------------------------------------------------------------------------------------Хочу Тусить
	// b.Handle(&keyboard.MediumRare, handlers.MediumRareHandle(b, DB))
	// b.Handle(&keyboard.MediumWell, handlers.MediumWellHandle(b, DB))
	// b.Handle(&keyboard.WellDone, handlers.WellDoneHandle(b, DB))
	// b.Handle(&keyboard.Craft, handlers.CraftHandle(b, DB))
	//---------------------------------------------------------------------------------------------------Админка
	b.Handle("Админка", handlers.AdminMenuHandle(b, DB))
	b.Handle(&keyboard.AddEvent, handlers.AddEventHandle(b, DB))
	b.Handle(&keyboard.AddNews, handlers.AddNewsHandle(b, DB))
	b.Handle(&keyboard.AddLost, handlers.AddLostHandle(b, DB))
	b.Handle(&keyboard.ViewSubs, handlers.ViewSubsHandle(b, DB))
	b.Handle(&keyboard.ViewComment, handlers.ViewCommentHandle(b, DB))
	b.Handle(&keyboard.OnGive, handlers.OnGiveHandle(b))
	//---------------------------------------------------------------------------------------------------TEST
	// b.Handle(&keyboard.Test, func(m *tb.Message) {
	// 	b.Send(m.Sender, "PRIVET")
	// })
	b.Handle(&keyboard.GiveButton, handlers.QREventHandle(b, DB)) // https://vkqr.ru/text

	//---------------------------------------------------------------------------------------------------TEST

	b.Start()
}

func MainMenuHandle(b *tb.Bot) func(*tb.Message) {
	return func(msg *tb.Message) {
		if keyboard.Give == false {
			b.Send(msg.Sender, "Возвращаемся :3", &tb.ReplyMarkup{
				ResizeReplyKeyboard: true,
				ReplyKeyboard:       keyboard.MainMenu})
		} else {
			b.Send(msg.Sender, "Возвращаемся :3", &tb.ReplyMarkup{
				ResizeReplyKeyboard: true,
				ReplyKeyboard:       keyboard.MainMenuGive})
		}
	}
}
