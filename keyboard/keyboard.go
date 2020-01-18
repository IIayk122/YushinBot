package keyboard

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	MainMenuBtn = tb.ReplyButton{Text: "Главное меню"}
	//__________________________________________________________________________________________________________________Главное меню

	EventsDay     = tb.ReplyButton{Text: "Мероприятия"}
	Record        = tb.ReplyButton{Text: "Запись на стрижку"}
	YushinMenuBtn = tb.ReplyButton{Text: "Меню"}
	FirstVisit    = tb.ReplyButton{Text: "Я новичок"}
	SecondVisit   = tb.ReplyButton{Text: "Я старожил"}
	GiveButton    = tb.ReplyButton{Text: "Розыгрыш"}

	MainMenu = [][]tb.ReplyButton{
		[]tb.ReplyButton{Record},
		[]tb.ReplyButton{EventsDay},
		[]tb.ReplyButton{YushinMenuBtn},
		[]tb.ReplyButton{FirstVisit, SecondVisit},
	}
	
	MainMenuGive = [][]tb.ReplyButton{
		[]tb.ReplyButton{GiveButton},
		[]tb.ReplyButton{Record},
		[]tb.ReplyButton{EventsDay},
		[]tb.ReplyButton{YushinMenuBtn},
		[]tb.ReplyButton{FirstVisit, SecondVisit},
	}

	//__________________________________________________________________________________________________________________МЕНЮ
	HairCuts = tb.ReplyButton{Text: "Стрижки"}
	Bar      = tb.ReplyButton{Text: "Карта бара"}
	Kitchen  = tb.ReplyButton{Text: "Кухня"}
	Smoke    = tb.ReplyButton{Text: "Кальяны"}
	Wear     = tb.ReplyButton{Text: "Мерч"}

	YushinMenu = [][]tb.ReplyButton{
		[]tb.ReplyButton{Bar},
		[]tb.ReplyButton{HairCuts, Kitchen},
		[]tb.ReplyButton{Smoke, Wear},
		[]tb.ReplyButton{MainMenuBtn},
	}

	//__________________________________________________________________________________________________________________Я СТАРОЖИЛ

	SubscribeEvent = tb.ReplyButton{Text: "Подписаться на мероприятия"}
	Comment        = tb.ReplyButton{Text: "Добавить отзыв"}
	WantSong       = tb.ReplyButton{Text: "Хочу спеть"}
	WantLearn      = tb.ReplyButton{Text: "Хочу научить"}
	Photos         = tb.ReplyButton{Text: "Фотоотчеты"}
	News           = tb.ReplyButton{Text: "Новости"}
	Lost           = tb.ReplyButton{Text: "Потеряшки"}

	SecondVisitMenu = [][]tb.ReplyButton{
		[]tb.ReplyButton{SubscribeEvent},
		[]tb.ReplyButton{WantSong, WantLearn},
		[]tb.ReplyButton{Photos, News},
		[]tb.ReplyButton{Comment, Lost},
		[]tb.ReplyButton{MainMenuBtn},
	}
	//__________________________________________________________________________________________________________________Я НОВИЧОК

	WhatDoing = tb.ReplyButton{Text: "Что у вас делать?"}
	In        = tb.ReplyButton{Text: "Я уже тут"}
	Out       = tb.ReplyButton{Text: "Я погнал дальше"}

	FirstVisitMenu = [][]tb.ReplyButton{
		[]tb.ReplyButton{WhatDoing},
		[]tb.ReplyButton{In},
		[]tb.ReplyButton{Out},
		[]tb.ReplyButton{MainMenuBtn},
	}
	//__________________________________________________________________________________________________________________ЧТО У ВАС ДЕЛАТЬ?

	MapYushin = tb.ReplyButton{Text: "Карта пространства"}
	Geo       = tb.ReplyButton{Text: "Как добраться"}

	WhatDoingMenu = [][]tb.ReplyButton{
		[]tb.ReplyButton{EventsWeek},
		[]tb.ReplyButton{MapYushin},
		[]tb.ReplyButton{Geo},
		[]tb.ReplyButton{MainMenuBtn},
	}
	//__________________________________________________________________________________________________________________Я УЖЕ ТУТ

	EventsWeek = tb.ReplyButton{Text: "Мероприятия на неделю"}

	InMenu = [][]tb.ReplyButton{
		[]tb.ReplyButton{EventsWeek},
		[]tb.ReplyButton{Bar},
		[]tb.ReplyButton{MainMenuBtn},
	}
	//__________________________________________________________________________________________________________________Я ПОГНАЛ ДАЛЬШЕ

	WantHome = tb.ReplyButton{Text: "Хочу домой"}
	WantClub = tb.ReplyButton{Text: "ХОЧУ ТУСИТЬ!!!"}

	OutMenu = [][]tb.ReplyButton{
		[]tb.ReplyButton{WantHome, WantClub},
		[]tb.ReplyButton{MainMenuBtn},
	}
	//__________________________________________________________________________________________________________________ХОЧУ ТУСИТЬ!!!
	MediumRare = tb.ReplyButton{Text: "Меньше 1000💸"}
	MediumWell = tb.ReplyButton{Text: "Больше 1000💵"}
	WellDone   = tb.ReplyButton{Text: "Я их не считаю 💰"}
	Craft      = tb.ReplyButton{Text: "Какие шекели? Я крафт люблю"}

	ClubMenu = [][]tb.ReplyButton{
		[]tb.ReplyButton{MediumRare},
		[]tb.ReplyButton{MediumWell},
		[]tb.ReplyButton{WellDone},
		[]tb.ReplyButton{Craft},
		[]tb.ReplyButton{MainMenuBtn},
	}
	//____________________________АДМИНКА_______________________________АДМИНКА_________________________АДМИНКА__________

	AddEvent    = tb.ReplyButton{Text: "Расписание 📅"}
	AddNews     = tb.ReplyButton{Text: "Новость 📰"}
	AddLost     = tb.ReplyButton{Text: "Потеряшка 👛"}
	ViewSubs    = tb.ReplyButton{Text: "Колличество подпичисков 🧾"}
	ViewComment = tb.ReplyButton{Text: "Отзывы 📝"}
	OnGive      = tb.ReplyButton{Text: "Розыгрыш🥳"}
	Give        = false

	AdminMenu = [][]tb.ReplyButton{
		[]tb.ReplyButton{AddEvent, OnGive},
		[]tb.ReplyButton{ViewSubs},
		[]tb.ReplyButton{AddNews, AddLost, ViewComment},
		[]tb.ReplyButton{MainMenuBtn},
	}
)
