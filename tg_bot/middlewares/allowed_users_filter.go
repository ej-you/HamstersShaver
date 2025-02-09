package middlewares

import (
	"fmt"
	"slices"

	telebot "gopkg.in/telebot.v3"
	
	customErrors "github.com/ej-you/HamstersShaver/tg_bot/errors"
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
			return fmt.Errorf("user %s is not allowed to use this bot: %w", userId, customErrors.AccessError("access denied"))
		}
		return nextHandler(context)
	}
}
