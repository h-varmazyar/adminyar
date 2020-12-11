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
	fmt.Println("command:", message.Command())
	switch message.Command() {
	case "start":
		args := message.CommandArguments()
		if args == "" {
			fmt.Println("start with no args")
			SendMessage(message.Chat.ID, model.Message{Text: root.WelcomeMessage, Type: model.TextMessage}, 0, &KeyboardSettings{
				UseInline:   false,
				UseKeyboard: true,
				Inline:      nil,
				Keyboard:    &WelcomeKeyboard,
			})
		} else {
			uuid, err := base64.StdEncoding.DecodeString(args)
			if err != nil {
				SendMessage(message.Chat.ID, model.Message{Text: root.WrongStartLink, Type: model.TextMessage}, 0, nil)
				return
			}
			num, err := strconv.ParseUint(string(uuid), 10, 32)
			if err != nil {
				SendMessage(message.Chat.ID, model.Message{Text: root.TryLaterNotice, Type: model.TextMessage}, 0, nil)
				return
			}
			channel, err := model.StartSession(message.From, uint32(num))
			if err != nil {
				SendMessage(message.Chat.ID, model.Message{Text: err.Error(), Type: model.TextMessage}, 0, nil)
				return
			}
			SendMessage(message.Chat.ID, model.Message{Text: fmt.Sprintf(root.StartConversation, channel.Name), Type: model.TextMessage}, 0, nil)
		}
	default:

	}
}
