package main

import (
	"time"

	telebot "gopkg.in/telebot.v3"

	"github.com/ej-you/HamstersShaver/tg_bot/handlers"
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

	// создание группы хэндлеров и добавление к ней middleware
	commandsHandlers := bot.Group()
	commandsHandlers.Use(middlewares.AllowedUsersFilter)
	commandsHandlers.Use(middlewares.CommandsLogger)

	// инициализация хендлеров
	commandsHandlers.Handle("/start", handlers.StartHandler)

	// запуск бота
	settings.InfoLog.Printf("Start bot %s...", bot.Me.Username)
	bot.Start()
}
