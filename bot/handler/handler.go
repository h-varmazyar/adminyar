package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	botAPI "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gobuffalo/envy"
	"github.com/mrNobody95/adminyar/bot/model"
	"log"
)

var bot *botAPI.BotAPI

func StartBotAPI() error {
	token := envy.Get("BOT_TOKEN", "")
	if token == "" {
		return errors.New("Bot token not found!")
	}
	var err error
	bot, err = botAPI.NewBotAPI(token)
	if err != nil {
		return err
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := botAPI.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	defer func() {
		for update := range updates {
			go func() {
				defer func() {
					if r := recover(); r != nil {
						log.Println(r)
					}
				}()
				if update.ChannelPost != nil || update.EditedChannelPost != nil {
					return
				}
				fmt.Println("in new go")
				if update.CallbackQuery != nil {
					if update.CallbackQuery.From.IsBot || model.CheckUser(update.CallbackQuery.From) != nil {
						//todo send error message
						return
					}
					fmt.Println("new 1")
					handleCallback(update.CallbackQuery)
				} else if update.Message.IsCommand() {
					if update.Message.From.IsBot || model.CheckUser(update.Message.From) != nil {
						//todo send error message
						return
					}
					fmt.Println("new 2")
					handleCommand(update.Message)
				} else if update.Message != nil {
					if update.Message.From.IsBot || model.CheckUser(update.Message.From) != nil {
						//todo send error message
						return
					}
					fmt.Println("new 3")
					handleMessage(update.Message)
				}
			}()
		}
	}()
	return nil
}

func SendMessage(chatId int64, message model.Message, replyId int, inline *botAPI.InlineKeyboardMarkup) {
	//var msg botAPI.Chattable
	var err error
	switch message.Type {
	case model.TextMessage:
		msg := botAPI.NewMessage(chatId, message.Text)
		if replyId > 0 {
			msg.ReplyToMessageID = replyId
		}
		if inline != nil {
			msg.ReplyMarkup = inline
		}
		_, err = bot.Send(msg)
	case model.VideoMessage:
		msg := botAPI.NewVideoShare(chatId, message.FileId)
		if replyId > 0 {
			msg.ReplyToMessageID = replyId
		}
		if inline != nil {
			msg.ReplyMarkup = inline
		}
		_, err = bot.Send(msg)
	case model.AudioMessage:
		msg := botAPI.NewAudioShare(chatId, message.FileId)
		if replyId > 0 {
			msg.ReplyToMessageID = replyId
		}
		if inline != nil {
			msg.ReplyMarkup = inline
		}
		_, err = bot.Send(msg)
	case model.VoiceMessage:
		msg := botAPI.NewVoiceShare(chatId, message.FileId)
		if replyId > 0 {
			msg.ReplyToMessageID = replyId
		}
		if inline != nil {
			msg.ReplyMarkup = inline
		}
		_, err = bot.Send(msg)
	case model.PhotoMessage:
		msg := botAPI.NewPhotoShare(chatId, message.FileId)
		if replyId > 0 {
			msg.ReplyToMessageID = replyId
		}
		if inline != nil {
			msg.ReplyMarkup = inline
		}
		_, err = bot.Send(msg)
	case model.ContactMessage:
		var c struct {
			FirstName   string
			PhoneNumber string
		}
		json.Unmarshal([]byte(message.MetaData), &c)
		msg := botAPI.NewContact(chatId, c.PhoneNumber, c.FirstName)
		if replyId > 0 {
			msg.ReplyToMessageID = replyId
		}
		if inline != nil {
			msg.ReplyMarkup = inline
		}
		_, err = bot.Send(msg)
	case model.StickerMessage:
		msg := botAPI.NewStickerShare(chatId, message.FileId)
		if replyId > 0 {
			msg.ReplyToMessageID = replyId
		}
		if inline != nil {
			msg.ReplyMarkup = inline
		}
		_, err = bot.Send(msg)
	case model.UnknownMessage:
		msg := botAPI.NewMessage(chatId, "Unknown message")
		if replyId > 0 {
			msg.ReplyToMessageID = replyId
		}
		if inline != nil {
			msg.ReplyMarkup = inline
		}
		_, err = bot.Send(msg)
	case model.LocationMessage:
		var c struct {
			Latitude  float64
			Longitude float64
		}
		json.Unmarshal([]byte(message.MetaData), &c)
		msg := botAPI.NewLocation(chatId, c.Latitude, c.Longitude)
		if replyId > 0 {
			msg.ReplyToMessageID = replyId
		}
		if inline != nil {
			msg.ReplyMarkup = inline
		}
		_, err = bot.Send(msg)
	case model.DocumentMessage:
		msg := botAPI.NewDocumentShare(chatId, message.FileId)
		msg.Caption = message.Text
		if replyId > 0 {
			msg.ReplyToMessageID = replyId
		}
		if inline != nil {
			msg.ReplyMarkup = inline
		}
		_, err = bot.Send(msg)
	case model.VideoNoteMessage:
		var length int
		json.Unmarshal([]byte(message.MetaData), &length)
		msg := botAPI.NewVideoNoteShare(chatId, length, message.FileId)
		if replyId > 0 {
			msg.ReplyToMessageID = replyId
		}
		if inline != nil {
			msg.ReplyMarkup = inline
		}
		_, err = bot.Send(msg)
	}
	if err != nil {
		panic(err.Error())
	}
}
