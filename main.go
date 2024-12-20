package main

import (
	"time"

	telebot "gopkg.in/telebot.v3"

	"github.com/ej-you/HamstersShaver/tg_bot/handlers"
	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"

	"github.com/ej-you/HamstersShaver/tg_bot/middlewares"
	"github.com/ej-you/HamstersShaver/tg_bot/settings"
)


func main() {
	// настройки бота
	pref := telebot.Settings{
		Token:  settings.BotToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		// Verbose: true,
		OnError: handlers.UnknownErrorHandler,
	}

	// инициализация бота
	bot, err := telebot.NewBot(pref)
	settings.DieIf(err)
	// фильт доступа юзеров к боту
	bot.Use(middlewares.AllowedUsersFilter)

	// инициализация клавиатур
	keyboards.InitKeyboards()

	// группа хендлеров для обработки всех команд
	commandsHandlers := bot.Group()
	commandsHandlers.Use(middlewares.GeneralCommandsLogger, middlewares.GeneralCommandsStatusFilter)
	// группа хендлеров для обработки всех основных (статичных, hard-code) инлайн-кнопок
	callbackHandlers := bot.Group()
	callbackHandlers.Use(middlewares.GeneralCallbackLogger, middlewares.GeneralCallbackStatusFilter)

	// инициализация хендлеров
	commandsHandlers.Handle("/start", handlers.StartHandler)
	
	commandsHandlers.Handle("/help", handlers.HelpHandler)
	callbackHandlers.Handle(&keyboards.BtnHideHelp, handlers.HelpHandler)

	commandsHandlers.Handle("/home", handlers.HomeHandler)
	commandsHandlers.Handle("/cancel", handlers.HomeHandler)
	callbackHandlers.Handle(&keyboards.BtnToHome, handlers.HomeHandler)

	commandsHandlers.Handle("/trade", handlers.TradeHandler)
	callbackHandlers.Handle(&keyboards.BtnToTrade, handlers.TradeHandler)
	
	// ВРЕМЕННО
	commandsHandlers.Handle("/buy", handlers.InDevelopmentHandler)
	callbackHandlers.Handle(&keyboards.BtnToBuy, handlers.InDevelopmentHandler)

	commandsHandlers.Handle("/cell", handlers.CellHandlerCommand)
	callbackHandlers.Handle(&keyboards.BtnToCell, handlers.CellHandlerCallback)

	// в разработке
	commandsHandlers.Handle("/auto", handlers.InDevelopmentHandler)
	callbackHandlers.Handle(&keyboards.BtnToAuto, handlers.InDevelopmentHandler)
	
	// в разработке
	commandsHandlers.Handle("/tokens", handlers.InDevelopmentHandler)
	callbackHandlers.Handle(&keyboards.BtnToTokens, handlers.InDevelopmentHandler)

	// запуск бота
	settings.InfoLog.Printf("Start bot %s...", bot.Me.Username)
	bot.Start()
}
