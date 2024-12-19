package keyboards

import (
	telebot "gopkg.in/telebot.v3"
)


var InlineKeyboardMainMenu = &telebot.ReplyMarkup{}
var BtnToTrade = InlineKeyboardMainMenu.Data("трейдинг", "to_trade")
var BtnToAuto = InlineKeyboardMainMenu.Data("авто", "to_auto")
var BtnToTokens = InlineKeyboardMainMenu.Data("сохранённые CA токенов", "to_tokens")
