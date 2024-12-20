package keyboards

import (
	telebot "gopkg.in/telebot.v3"
)


var InlineKeyboardTrade = &telebot.ReplyMarkup{}
var BtnToBuy = InlineKeyboardTrade.Data("купить", "to_buy")
var BtnToCell = InlineKeyboardTrade.Data("продать", "to_cell")
