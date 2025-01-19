package main

import (
	"time"

	telebot "gopkg.in/telebot.v3"

	handlersTradeCell "github.com/ej-you/HamstersShaver/tg_bot/handlers/trade/cell"
	handlersTradeBuy "github.com/ej-you/HamstersShaver/tg_bot/handlers/trade/buy"
	handlersTrade "github.com/ej-you/HamstersShaver/tg_bot/handlers/trade"
	handlersHelpers "github.com/ej-you/HamstersShaver/tg_bot/handlers/helpers"
	"github.com/ej-you/HamstersShaver/tg_bot/handlers"

	"github.com/ej-you/HamstersShaver/tg_bot/mongo"
	"github.com/ej-you/HamstersShaver/tg_bot/redis"
	
	customErrors "github.com/ej-you/HamstersShaver/tg_bot/errors"
	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"
	"github.com/ej-you/HamstersShaver/tg_bot/middlewares"
	"github.com/ej-you/HamstersShaver/tg_bot/settings"
)


func main() {
	// проверка переменных окружения
	settings.CheckEnv()
	// получаем клиенты для redis и mongo для проверки, что соединение есть
	_ = redis.GetRedisClient()
	_ = mongo.NewMongoDB()

	// настройки бота
	pref := telebot.Settings{
		Token:  settings.BotToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		// Verbose: true,
		OnError: customErrors.MainErrorHandler,
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
	callbackHandlers.Handle(&keyboards.BtnToHome, handlers.HomeHandler)

	commandsHandlers.Handle("/cancel", handlers.CancelHandler)
	callbackHandlers.Handle(&keyboards.BtnCancel, handlers.CancelHandler)

	commandsHandlers.Handle("/trade", handlersTrade.TradeHandler)
	callbackHandlers.Handle(&keyboards.BtnToTrade, handlersTrade.TradeHandler)
	
	commandsHandlers.Handle("/buy", handlersTradeBuy.BuyHandlerCommand)
	callbackHandlers.Handle(&keyboards.BtnToBuy, handlersTradeBuy.BuyHandlerCallback)

	commandsHandlers.Handle("/cell", handlersTradeCell.CellHandlerCommand)
	callbackHandlers.Handle(&keyboards.BtnToCell, handlersTradeCell.CellHandlerCallback)

	// в разработке (Dedust.io DEX-биржа)
	callbackHandlers.Handle(&keyboards.BtnDedust, handlersHelpers.InDevelopmentHandler)

	// в разработке (функция авто)
	commandsHandlers.Handle("/auto", handlersHelpers.InDevelopmentHandler)
	callbackHandlers.Handle(&keyboards.BtnToAuto, handlersHelpers.InDevelopmentHandler)
	
	// в разработке (сохранённые токены для покупки)
	commandsHandlers.Handle("/tokens", handlersHelpers.InDevelopmentHandler)
	callbackHandlers.Handle(&keyboards.BtnToTokens, handlersHelpers.InDevelopmentHandler)

	bot.Handle(telebot.OnText, handlersHelpers.HandlersDistributor, middlewares.GeneralCommandsStatusFilter)
	bot.Handle(telebot.OnCallback, handlersHelpers.HandlersDistributor, middlewares.GeneralCallbackStatusFilter)

	// запуск бота
	settings.InfoLog.Printf("Start bot %s...", bot.Me.Username)
	bot.Start()
}
