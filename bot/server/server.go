package server

import (
	"NewYushinBot/bot/fsm"
	"NewYushinBot/bot/userdata"

	"gopkg.in/tucnak/telebot.v2"
)

var Bot *BotServer

type BotServer struct {
	Bot           *telebot.Bot
	UsersData     *userdata.UsersData
	InlineButtons map[string]*telebot.InlineButton
	ReplyButtons  map[string]*telebot.ReplyButton
}

func NewBotServer(bot *telebot.Bot) *BotServer {
	return &BotServer{
		Bot:           bot,
		UsersData:     userdata.NewUsersData(fsm.None),
		InlineButtons: make(map[string]*telebot.InlineButton),
		ReplyButtons:  make(map[string]*telebot.ReplyButton),
	}
}

func (b *BotServer) Start() {
	b.Bot.Start()
}

func (b *BotServer) Stop() {
	b.Bot.Stop()
}

func (b *BotServer) Delete(msg telebot.Editable) error {
	return b.Bot.Delete(msg)
}

// Handle text endpoint with Message handler.
func (b *BotServer) HandleMessage(endpoint string, handler func(*telebot.Message)) {
	b.Bot.Handle(endpoint, func(m *telebot.Message) {
		handler(m)
	})
}

// Handle inline button endpoint with callback handler.
func (b *BotServer) HandleInline(endpoint *telebot.InlineButton, handler func(*telebot.Callback)) {
	b.Bot.Handle(endpoint, func(c *telebot.Callback) {
		handler(c)
	})
}

// Handle reply button endpoint with callback handler.
func (b *BotServer) HandleReply(endpoint *telebot.ReplyButton, handler func(*telebot.Message)) {
	b.Bot.Handle(endpoint, func(m *telebot.Message) {
		handler(m)
	})
}

// Reply content to a specific message.
func (b *BotServer) Reply(m *telebot.Message, what interface{}, options ...interface{}) error {
	_, err := b.Bot.Send(m.Chat, what, options...)
	if err != nil {
		_, err = b.Bot.Send(m.Chat, "Something went wrong, Please retry.")
	}
	return err
}

// Send content to a specific chat (ByID).
func (b *BotServer) Send(c *telebot.Chat, what interface{}, options ...interface{}) error {
	_, err := b.Bot.Send(c, what, options...)
	return err
}
