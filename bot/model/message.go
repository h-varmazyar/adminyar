package model

import botAPI "github.com/go-telegram-bot-api/telegram-bot-api"

type MessageType string

const (
	TextMessage      MessageType = "Text"
	VideoMessage     MessageType = "Video"
	AudioMessage     MessageType = "Audio"
	VoiceMessage     MessageType = "Voice"
	PhotoMessage     MessageType = "photo"
	ContactMessage   MessageType = "Contact"
	StickerMessage   MessageType = "Sticker"
	UnknownMessage   MessageType = "Unknown"
	LocationMessage  MessageType = "Location"
	DocumentMessage  MessageType = "Document"
	VideoNoteMessage MessageType = "VideoNote"
)

type Message struct {
	Text     string
	Type     MessageType
	MetaData string
	FileId   string
}

func StartSession(user *botAPI.User, uuid uint32) (Channel, error) {
	return Channel{}, nil
}
