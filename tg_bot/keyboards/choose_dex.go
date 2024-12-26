package keyboards

import (
	telebot "gopkg.in/telebot.v3"
)


var InlineKeyboardChooseDEX = &telebot.ReplyMarkup{}
var BtnStonfi = InlineKeyboardChooseDEX.Data("Ston.fi", "stonfi")
var BtnDedust = InlineKeyboardChooseDEX.Data("Dedust.io", "dedust")
