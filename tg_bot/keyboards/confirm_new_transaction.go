package keyboards

import (
	telebot "gopkg.in/telebot.v3"
)


var InlineKeyboardConfirmNewTransaction = &telebot.ReplyMarkup{}
var BtnConfirm = InlineKeyboardConfirmNewTransaction.Data("подтвердить", "confirm")
var BtnCancel = InlineKeyboardConfirmNewTransaction.Data("отмена", "cancel")
