package handler

import (
	"fmt"
	botAPI "github.com/go-telegram-bot-api/telegram-bot-api"
	root "github.com/mrNobody95/adminyar/bot"
	"github.com/mrNobody95/adminyar/bot/model"
)

func handleMessage(message *botAPI.Message) {
	switch message.Text {
	case AddNewChannel:
		err := model.ChangeUserStatus(message.From.ID, model.AddChannelId)
		if err != nil {
			panic(err)
		}
		SendMessage(message.Chat.ID, model.Message{
			Text: "آیدی کانال خود را وارد کنید:",
			Type: model.TextMessage,
		}, 0, nil)
		return
	case ChannelList:
		channels, err := model.GetChannelList(message.From.ID)
		if err != nil {
			panic(err)
		}
		rows := make([][]botAPI.InlineKeyboardButton, len(channels)/2+len(channels)%2)
		if len(channels) > 1 {
			for i := 0; i < len(channels); i++ {
				rows[i/2] = botAPI.NewInlineKeyboardRow(
					botAPI.NewInlineKeyboardButtonData(channels[i].Name, fmt.Sprintf("%d", channels[i].UUID)),
					botAPI.NewInlineKeyboardButtonData(channels[i+1].Name, fmt.Sprintf("%d", channels[i+1].UUID)),
				)
			}
		}
		if len(channels)%2 == 1 {
			rows[len(rows)-1] = botAPI.NewInlineKeyboardRow(
				botAPI.NewInlineKeyboardButtonData(channels[len(channels)-1].Name, fmt.Sprintf("%d", channels[len(channels)-1].UUID)),
			)
		}

		channelsKeyboard := botAPI.InlineKeyboardMarkup{InlineKeyboard: rows}
		SendMessage(message.Chat.ID, model.Message{
			Type: model.TextMessage,
			Text: root.ChannelListFormat}, 0, &KeyboardSettings{
			UseInline: true,
			Inline:    &channelsKeyboard,
		})
		return
	case Help:
		return
	case About:
		return
	}
	user, err := model.LoadUser(message.From.ID)
	if err != nil {
		panic(err)
	}
	switch user.Status {
	case model.Free:
		return
	case model.AddChannelId:
		if root.ValidateChannelId(message.Text) {
			err = model.CreateNewChannel(message.Text)
			if err != nil {
				panic(err)
			}
			err = model.ChangeUserStatus(message.From.ID, model.AddChannelName)
			if err != nil {
				panic(err)
			}
			SendMessage(message.Chat.ID, model.Message{
				Type: model.TextMessage,
				Text: "نام کانال خود را وارد کنید:",
			}, 0, nil)
		} else {
			SendMessage(message.Chat.ID, model.Message{
				Type: model.TextMessage,
				Text: root.WrongChannelId,
			}, 0, nil)
		}
		return
	case model.AddChannelName:
		if root.ValidateChannelName(message.Text) {
			uuid, err := model.UpdateChannelName(message.Text)
			if err != nil {
				panic(err)
			}
			err = model.ChangeUserStatus(message.From.ID, model.Free)
			if err != nil {
				panic(err)
			}
			SendMessage(message.Chat.ID, model.Message{
				Type: model.TextMessage,
				Text: fmt.Sprintf(root.ChannelAdded, message.Text, model.GenerateChannelLink(uuid)),
			}, 0, &KeyboardSettings{
				UseKeyboard: true,
				Keyboard:    &WelcomeKeyboard,
			})
		} else {
			SendMessage(message.Chat.ID, model.Message{
				Type: model.TextMessage,
				Text: root.WrongChannelId,
			}, 0, nil)
		}
		return
	}
}
