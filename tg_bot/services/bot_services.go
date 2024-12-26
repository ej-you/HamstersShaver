package services

import (
	"strconv"
	"strings"

	telebot "gopkg.in/telebot.v3"
)


// Возвращает строковый id юзера ТГ
func GetUserID(chat *telebot.Chat) string {
	return strconv.FormatInt(chat.ID, 10)
}


// возвращает данные из callback
func GetCallbackData(callback *telebot.Callback) string {
	callbackInfo := callback.Unique
	if callbackInfo == "" {
		// \f добавляется библиотекой, поэтому убираем его
		callbackInfo = strings.TrimPrefix(callback.Data, "\f")
		callback.Data = callbackInfo
	}
	return callbackInfo
}
