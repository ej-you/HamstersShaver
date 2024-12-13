package middlewares

import (
	"errors"
	"fmt"
	"slices"

	telebot "gopkg.in/telebot.v3"
	
	"github.com/ej-you/HamstersShaver/tg_bot/services"
	"github.com/ej-you/HamstersShaver/tg_bot/settings"
)


// отсеивание юзеров без доступа к этому боту
func AllowedUsersFilter(nextHandler telebot.HandlerFunc) telebot.HandlerFunc {
	return func(context telebot.Context) error {
		// ТГ ID юзера
		userId := services.GetUserID(context.Chat())

		// проверка на наличие юзера в списке разрешённых юзеров с доступом к боту
		if !slices.Contains(settings.AllowedUsers, userId) {
			return errors.New(fmt.Sprintf("User %s is not allowed to use this bot", userId))
		}
		return nextHandler(context)
	}
}
