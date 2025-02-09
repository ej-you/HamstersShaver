package keyboards

import (
	telebot "gopkg.in/telebot.v3"
)


var InlineKeyboardToHome = &telebot.ReplyMarkup{}
var BtnToHome = InlineKeyboardToHome.Data("главное меню", "to_home")
