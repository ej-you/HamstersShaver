package keyboards

import (
	telebot "gopkg.in/telebot.v3"
)


var InlineKeyboardJettonsAmountChoices = &telebot.ReplyMarkup{}
var BtnJettonsAmountChoice25 = InlineKeyboardJettonsAmountChoices.Data("25%", "jettons_amount_choice", "25")
var BtnJettonsAmountChoice50 = InlineKeyboardJettonsAmountChoices.Data("50%", "jettons_amount_choice", "50")
var BtnJettonsAmountChoice75 = InlineKeyboardJettonsAmountChoices.Data("75%", "jettons_amount_choice", "75")
var BtnJettonsAmountChoice100 = InlineKeyboardJettonsAmountChoices.Data("100%", "jettons_amount_choice", "100")
