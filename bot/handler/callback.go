package handler

import (
	"fmt"
	botAPI "github.com/go-telegram-bot-api/telegram-bot-api"
	root "github.com/mrNobody95/adminyar/bot"
	"github.com/mrNobody95/adminyar/bot/model"
	"strconv"
	"strings"
)

func handleCallback(query *botAPI.CallbackQuery) {
	tmp := strings.Split(query.Data, "-")
	switch tmp[0] {
	case "back":
		channels, err := model.GetChannelList(query.From.ID)
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
		EditMessage(query.Message.Chat.ID, root.ChannelListFormat, query.Message.MessageID, false, &channelsKeyboard)
	case "editName":
		err := model.ChangeUserStatus(query.From.ID, model.EditChannelName)
		if err != nil {
			panic(err)
		}
		SendMessage(query.Message.Chat.ID, model.Message{
			Text: "نام جدید کانال خود را وارد کنید:",
			Type: model.TextMessage,
		}, 0, nil)
	case "editId":
		err := model.ChangeUserStatus(query.From.ID, model.EditChannelId)
		if err != nil {
			panic(err)
		}
		SendMessage(query.Message.Chat.ID, model.Message{
			Text: "آیدی جدید کانال خود را وارد کنید:",
			Type: model.TextMessage,
		}, 0, nil)
	case "delete":
		uuid, err := strconv.ParseUint(tmp[1], 10, 32)
		if err != nil {
			panic(err)
		}
		channel, err := model.GetChannel(uint32(uuid))
		if err != nil {
			panic(err)
		}
		deleteChannelKeyboard := botAPI.NewInlineKeyboardMarkup(
			botAPI.NewInlineKeyboardRow(
				botAPI.NewInlineKeyboardButtonData("بله پاکش کن", fmt.Sprintf("confirmDelete-%s", tmp[1])),
				botAPI.NewInlineKeyboardButtonData("نه نمیخواد", fmt.Sprintf("cancelDelete-%s", tmp[1])),
			),
		)
		SendMessage(query.Message.Chat.ID, model.Message{
			Text: fmt.Sprintf(root.DeleteChannel, channel.Name, channel.Id),
			Type: model.TextMessage,
		}, 0, &KeyboardSettings{
			UseInline: true,
			Inline:    &deleteChannelKeyboard,
		})
	case "confirmDelete":
		uuid, err := strconv.ParseUint(tmp[1], 10, 32)
		if err != nil {
			panic(err)
		}
		err = model.DeleteChannel(uint32(uuid))
		if err != nil {
			panic(err)
		}
		SendMessage(query.Message.Chat.ID, model.Message{
			Text: "کانال با موفقیت حذف شد",
			Type: model.TextMessage,
		}, 0, &KeyboardSettings{
			UseInline:   false,
			UseKeyboard: true,
			Keyboard:    &WelcomeKeyboard,
		})
	case "cancelDelete":
		bot.DeleteMessage(botAPI.NewDeleteMessage(query.Message.Chat.ID, query.Message.MessageID))
	case "link":
		uuid, err := strconv.ParseUint(tmp[1], 10, 32)
		if err != nil {
			panic(err)
		}
		channel, err := model.RevokeChannel(uint32(uuid))
		if err != nil {
			panic(err)
		}
		SendMessage(query.Message.Chat.ID, model.Message{
			Type: model.TextMessage,
			Text: fmt.Sprintf(root.ChannelLink, channel.Name, model.GenerateChannelLink(channel.UUID)),
		}, 0, &KeyboardSettings{
			UseKeyboard: true,
			Keyboard:    &WelcomeKeyboard,
		})
	case "select":
		channelKeyboard := botAPI.NewInlineKeyboardMarkup(
			botAPI.NewInlineKeyboardRow(
				botAPI.NewInlineKeyboardButtonData("تغییر نام", fmt.Sprintf("editName-%s", tmp[1])),
				botAPI.NewInlineKeyboardButtonData("تغییر آیدی", fmt.Sprintf("editId-%s", tmp[1])),
			),
			botAPI.NewInlineKeyboardRow(
				botAPI.NewInlineKeyboardButtonData("لینک", fmt.Sprintf("link-%s", tmp[1])),
				botAPI.NewInlineKeyboardButtonData("حذف", fmt.Sprintf("delete-%s", tmp[1])),
			),
			botAPI.NewInlineKeyboardRow(
				botAPI.NewInlineKeyboardButtonData("بازگشت", "back"),
			),
		)
		EditMessage(query.Message.Chat.ID, "میخواین چه کاری با کانالتون انجام بدین؟", query.Message.MessageID, false, &channelKeyboard)
	default:
		SendMessage(query.Message.Chat.ID, model.Message{
			Type: model.TextMessage,
			Text: "درخواست شما معتبر نیست",
		}, 0, &KeyboardSettings{
			UseKeyboard: true,
			Keyboard:    &WelcomeKeyboard,
		})
	}

}
