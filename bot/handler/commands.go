package handler

import (
	"encoding/base64"
	"fmt"
	botAPI "github.com/go-telegram-bot-api/telegram-bot-api"
	root "github.com/mrNobody95/adminyar/bot"
	"github.com/mrNobody95/adminyar/bot/model"
	"strconv"
)

func handleCommand(message *botAPI.Message) {
	switch message.Command() {
	case "start":
		args := message.CommandArguments()
		if args == "" {
			SendMessage(message.Chat.ID, model.Message{Text: root.WelcomeMessage}, 0, nil)
		} else {
			uuid, err := base64.StdEncoding.DecodeString(args)
			if err != nil {
				SendMessage(message.Chat.ID, model.Message{Text: root.WrongStartLink}, 0, nil)
				return
			}
			num, err := strconv.ParseUint(string(uuid), 10, 32)
			if err != nil {
				SendMessage(message.Chat.ID, model.Message{Text: root.TryLaterNotice}, 0, nil)
				return
			}
			channel, err := model.StartSession(message.From, uint32(num))
			if err != nil {
				SendMessage(message.Chat.ID, model.Message{Text: err.Error()}, 0, nil)
				return
			}
			SendMessage(message.Chat.ID, model.Message{Text: fmt.Sprintf(root.StartConversation, channel.Name)}, 0, nil)
		}

	default:

	}
}
