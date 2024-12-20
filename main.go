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

	// инициализация клавиатур
	keyboards.InitKeyboards()

	// создание группы хэндлеров и добавление к ней middleware
	commandsHandlers := bot.Group()
	commandsHandlers.Use(middlewares.AllowedUsersFilter)
	commandsHandlers.Use(middlewares.CommandsLogger)

	// инициализация хендлеров
	commandsHandlers.Handle("/start", handlers.StartHandler)
	
	commandsHandlers.Handle("/help", handlers.HelpHandler)
	commandsHandlers.Handle(&keyboards.BtnHideHelp, handlers.HelpHandler)

	commandsHandlers.Handle("/home", handlers.HomeHandler)
	commandsHandlers.Handle("/cancel", handlers.HomeHandler)
	commandsHandlers.Handle(&keyboards.BtnToHome, handlers.HomeHandler)

	commandsHandlers.Handle("/trade", handlers.TradeHandler)
	commandsHandlers.Handle(&keyboards.BtnToTrade, handlers.TradeHandler)
	
	// ВРЕМЕННО
	commandsHandlers.Handle("/buy", handlers.InDevelopmentHandler)
	commandsHandlers.Handle(&keyboards.BtnToBuy, handlers.InDevelopmentHandler)

	commandsHandlers.Handle("/cell", handlers.CellHandlerCommand)
	commandsHandlers.Handle(&keyboards.BtnToCell, handlers.CellHandlerCallback)

	// в разработке
	commandsHandlers.Handle("/auto", handlers.InDevelopmentHandler)
	commandsHandlers.Handle(&keyboards.BtnToAuto, handlers.InDevelopmentHandler)
	
	// в разработке
	commandsHandlers.Handle("/tokens", handlers.InDevelopmentHandler)
	commandsHandlers.Handle(&keyboards.BtnToTokens, handlers.InDevelopmentHandler)

	// запуск бота
	settings.InfoLog.Printf("Start bot %s...", bot.Me.Username)
	bot.Start()
}
