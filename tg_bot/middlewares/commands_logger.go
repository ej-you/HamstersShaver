package middlewares

import (
	telebot "gopkg.in/telebot.v3"
	
	"github.com/ej-you/HamstersShaver/tg_bot/services"
	"github.com/ej-you/HamstersShaver/tg_bot/settings"
)


// перед каждой введённой юзером командой записывает её в лог
func CommandsLogger(nextHandler telebot.HandlerFunc) telebot.HandlerFunc {
	return func(context telebot.Context) error {
		// ТГ ID юзера
		userId := services.GetUserID(context.Chat())

		// callback := context.Callback()
		// // если был переход по кнопке
		// if callback != nil {
		// 	var analogCommand string

		// 	// определяем аналогичную команду по уникальному описанию кнопки (даётся при создании)
		// 	switch callback.Unique {
		// 		case "back_to_home":
		// 			analogCommand = "/home"
		// 		case "course_back_to_home":
		// 			analogCommand = "/home"

		// 		case "get_currencies":
		// 			analogCommand = "/currencies"

		// 		case "get_cur_course":
		// 			analogCommand = "/course"
		// 		case "get_course_again":
		// 			analogCommand = "/course"
		// 		default:
		// 			return nextHandler(context)
		// 	}
		// 	settings.InfoLog.Printf("User %s use button like %q", userId, analogCommand)
		// // если был переход по команде
		// } else {
		settings.InfoLog.Printf("User %s use command %q", userId, context.Message().Text)
		// }

		return nextHandler(context)
	}
}
