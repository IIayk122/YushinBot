package controller

import (
	"NewYushinBot/bot/repository"
	"NewYushinBot/bot/server"
	"NewYushinBot/keyboard"

	"github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
)

//StartHandle приветствует нового пользователя и записывает его в базу если он уникальный.
func StartHandle(m *telebot.Message) {

	err := repository.SetUser(*m.Sender)
	if err != nil {
		logrus.Errorf("%s", err)
	}
	if !keyboard.Give {
		err = server.Bot.Send(m.Chat,
			"Привет "+m.Sender.FirstName+"!\n"+"Добро пожаловать в YushinBot",
			&telebot.ReplyMarkup{
				ResizeReplyKeyboard: true,
				ReplyKeyboard:       keyboard.MainMenu},
		)
		if err != nil {
			logrus.Errorf("%s", err)
		}
	} else {
		err = server.Bot.Send(m.Chat,
			"Привет "+m.Sender.FirstName+"!\n"+"Добро пожаловать в YushinBot",
			&telebot.ReplyMarkup{
				ResizeReplyKeyboard: true,
				ReplyKeyboard:       keyboard.MainMenuGive},
		)
		if err != nil {
			logrus.Errorf("%s", err)
		}
	}

}

// func StartCtrl(m *telebot.Message) {
// 	_ = server.Bot.Reply(m, "ololo")
// }
