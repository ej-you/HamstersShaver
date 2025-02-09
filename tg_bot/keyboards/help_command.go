package keyboards

import (
	telebot "gopkg.in/telebot.v3"
)


var InlineKeyboardHelp = &telebot.ReplyMarkup{}
var BtnHideHelp = InlineKeyboardHelp.Data("скрыть справку", "hide_help")
