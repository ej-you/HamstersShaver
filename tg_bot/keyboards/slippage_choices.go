package keyboards

import (
	telebot "gopkg.in/telebot.v3"
)


var InlineKeyboardSlippageChoices = &telebot.ReplyMarkup{}
var BtnSlippageChoice15 = InlineKeyboardSlippageChoices.Data("15%", "slippage_choice", "15")
var BtnSlippageChoice30 = InlineKeyboardSlippageChoices.Data("30%", "slippage_choice", "30")
var BtnSlippageChoice60 = InlineKeyboardSlippageChoices.Data("60%", "slippage_choice", "60")
var BtnSlippageChoice100 = InlineKeyboardSlippageChoices.Data("100%", "slippage_choice", "100")
