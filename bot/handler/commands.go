package handler

import (
	"encoding/base64"
	botAPI "github.com/go-telegram-bot-api/telegram-bot-api"
	root "github.com/mrNobody95/adminyar/bot"
	"github.com/mrNobody95/adminyar/bot/model"
	"strconv"
	"strings"
)

func handleCommand(message *botAPI.Message) {
	switch message.Command() {
	case "start":
		args := message.CommandArguments()
		if args == "" {
		} else {
			uuid, err := base64.StdEncoding.DecodeString(args)
			if err != nil {
				SendMessage(message.Chat.ID, model.Message{Text: root.WrongStartLink}, 0, nil)
				return
			}
			strconv.ParseUint(string(uuid), 10, 32)
			model.StartSession(message.From, uuid)

		}

	default:

	}
}
