package services

import (
	"strconv"

	telebot "gopkg.in/telebot.v3"
)


// Возвращает строковый id юзера ТГ
func GetUserID(chat *telebot.Chat) string {
	return strconv.FormatInt(chat.ID, 10)
}
