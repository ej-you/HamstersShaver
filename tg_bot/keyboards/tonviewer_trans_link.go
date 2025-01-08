package keyboards

import (
	"fmt"

	telebot "gopkg.in/telebot.v3"
)


var InlineKeyboardTonviewerTransLink = &telebot.ReplyMarkup{}

// настройка клавиатуры с кнопкой с ссылкой на полную информацию по транзакции
func SetTonviewerTransLink(transactionHash string) error {
	btnWithLink := InlineKeyboardTonviewerTransLink.Data("смотреть больше", "")
	btnWithLink.URL = fmt.Sprintf("https://tonviewer.com/transaction/%s?section=valueFlow", transactionHash)

	InlineKeyboardTonviewerTransLink.Inline(
		InlineKeyboardTonviewerTransLink.Row(btnWithLink),
	)
	return nil
}
