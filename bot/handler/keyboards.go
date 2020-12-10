package handler

import botAPI "github.com/go-telegram-bot-api/telegram-bot-api"

type KeyboardSettings struct {
	UseInline   bool
	UseKeyboard bool
	Inline      *botAPI.InlineKeyboardMarkup
	Keyboard    *botAPI.ReplyKeyboardMarkup
}
