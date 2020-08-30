package bot

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"NewYushinBot/bot/controller"
	"NewYushinBot/bot/server"
	"NewYushinBot/config"
	"NewYushinBot/keyboard"

	"gopkg.in/tucnak/telebot.v2"
)

func Setup() error {
	b, err := telebot.NewBot(telebot.Settings{
		Token:   config.Configs.Bot.Token,
		Verbose: false,
		Poller: &telebot.LongPoller{
			Timeout: time.Second * time.Duration(config.Configs.Bot.PollerTimeout),
		},
	})
	if err != nil {
		return err
	}

	fmt.Println()
	log.Println("[Telebot] Success to connect telegram bot:", b.Me.Username)
	fmt.Println()

	server.Bot = server.NewBotServer(b)
	initHandler(server.Bot)

	return nil
}
func initHandler(b *server.BotServer) {
	//	b.HandleReply(&keyboard.MainMenuBtn, MainMenuHandle(b.Bot))
	// b.InlineButtons["btn_unbind"] = &telebot.InlineButton{Unique: "btn_unbind", Text: "Unbind"}
	// b.InlineButtons["btn_cancel_unbind"] = &telebot.InlineButton{Unique: "btn_cancel_unbind", Text: "Cancel"}
	// b.InlineButtons["btn_delete_notifier"] = &telebot.InlineButton{Unique: "btn_delete_notifier", Text: "Delete"}
	// b.InlineButtons["btn_cancel_delete_notifier"] = &telebot.InlineButton{Unique: "btn_cancel_delete_notifier", Text: "Cancel"}

	b.HandleMessage("/hello", controller.StartHandle)
	// b.HandleMessage("/help", controller.HelpCtrl)
	// b.HandleMessage("/cancel", controller.CancelCtrl)
	// b.HandleMessage("/action", controller.ActionCtrl)

	// b.HandleMessage("/bind", controller.BindCtrl)
	// b.HandleMessage("/unbind", controller.UnbindCtrl)
	// b.HandleInline(b.InlineButtons["btn_unbind"], controller.InlBtnUnbindCtrl)
	// b.HandleInline(b.InlineButtons["btn_cancel_unbind"], controller.InlBtnCancelUnbindCtrl)

	// b.HandleMessage("/shownotifiers", controller.ShowNotifiersCtrl)
	// b.HandleMessage("/shownotifier", controller.ShowNotifierCtrl)
	// b.HandleMessage("/addnotifier", controller.AddNotifierCtrl)
	// b.HandleMessage("/updatenotifiertitle", controller.UpdateNotifierTitleCtrl)
	// b.HandleMessage("/updatenotifier", controller.UpdateNotifierCtrl)
	// b.HandleMessage("/deletenotifier", controller.DeleteNotifierCtrl)
	// b.HandleInline(b.InlineButtons["btn_delete_notifier"], controller.InlBtnDeleteNotifierCtrl)
	// b.HandleInline(b.InlineButtons["btn_cancel_delete_notifier"], controller.InlBtnCancelDeleteNotifierCtrl)

	// b.HandleMessage(telebot.OnText, controller.OnTextCtrl)
}

// func StartCtrl(m *telebot.Message) {
// 	_ = server.Bot.Reply(m, "Hello Wrold )))")
// }

func SendToChat(chatId int64, what interface{}, options ...interface{}) error {
	chat, err := server.Bot.Bot.ChatByID(strconv.FormatInt(chatId, 10))
	if err != nil {
		return err
	}

	return server.Bot.Send(chat, what, options...)
}
func MainMenuHandle(b *telebot.Bot) func(*telebot.Message) {
	return func(msg *telebot.Message) {
		if !keyboard.Give {
			b.Send(msg.Sender, "Возвращаемся :3", &telebot.ReplyMarkup{
				ResizeReplyKeyboard: false,
				ReplyKeyboard:       keyboard.MainMenu})
		} else {
			b.Send(msg.Sender, "Возвращаемся :3", &telebot.ReplyMarkup{
				ResizeReplyKeyboard: false,
				ReplyKeyboard:       keyboard.MainMenuGive})
		}
	}
}
