package handler

import botAPI "github.com/go-telegram-bot-api/telegram-bot-api"

type KeyboardSettings struct {
	UseInline   bool
	UseKeyboard bool
	Inline      *botAPI.InlineKeyboardMarkup
	Keyboard    *botAPI.ReplyKeyboardMarkup
}

const (
	AddNewChannel = "افزودن کانال جدید"
	ChannelList   = "فهرست کانال های من"
	Help          = "راهنما"
	About         = "درباره ما"
)

var WelcomeKeyboard = botAPI.NewReplyKeyboard(
	botAPI.NewKeyboardButtonRow(
		botAPI.NewKeyboardButton(AddNewChannel),
	),
	botAPI.NewKeyboardButtonRow(
		botAPI.NewKeyboardButton(ChannelList),
	),
	botAPI.NewKeyboardButtonRow(
		botAPI.NewKeyboardButton(Help),
		botAPI.NewKeyboardButton(About),
	),
)
