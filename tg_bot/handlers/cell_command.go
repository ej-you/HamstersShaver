package handlers

import (
	// "fmt"

	telebot "gopkg.in/telebot.v3"

	// "github.com/ej-you/HamstersShaver/tg_bot/keyboards"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
)


// команды: /cell
func CellHandlerCommand(context telebot.Context) error {
	var err error
	userId := services.GetUserID(context.Chat())

	// получение машины состоянию текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(userId)
	// обновление кэша
	if err = cellUpdateCache(userStateMachine); err != nil {
		return err
	}

	msgText := `📉 Активирован диалог продажи монет. Для отмены всех действий и выхода в главное меню используйте /cancel

Выберите из имеющихся у вас на аккаунте монет ту, которую хотите продать 👇`

	// // создание клавиатуры в соответствии с текущим списком монет на кошельке аккаунта
	// var inlineKeyboardWalletJettons = telebot.ReplyMarkup{}
	// err = keyboards.SetWalletJettonsBtnRows(&inlineKeyboardWalletJettons)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println("inlineKeyboardWalletJettons:", inlineKeyboardWalletJettons)

	return context.Send(msgText) //, &inlineKeyboardWalletJettons)
}


// кнопки: to_cell
func CellHandlerCallback(context telebot.Context) error {
	var err error
	userId := services.GetUserID(context.Chat())

	// получение машины состоянию текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(userId)
	// обновление кэша
	if err = cellUpdateCache(userStateMachine); err != nil {
		return err
	}


	msgText := `Отлично! Выбрано действие продажи монет 📉

Выберите из имеющихся у вас на аккаунте монет ту, которую хотите продать 👇`

	// // создание клавиатуры в соответствии с текущим списком монет на кошельке аккаунта
	// var inlineKeyboardWalletJettons = telebot.ReplyMarkup{}
	// err = keyboards.SetWalletJettonsBtnRows(&inlineKeyboardWalletJettons)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println("inlineKeyboardWalletJettons:", inlineKeyboardWalletJettons)
	
	return context.Send(msgText)
}

func cellUpdateCache(userStateMachine stateMachine.UserStateMachine) error {
	var err error

	// установка нового состояния
	if err = userStateMachine.SetStatus("cell"); err != nil {
		return err
	}
	// установка значения действия
	if err = userStateMachine.SetAction("cell"); err != nil {
		return err
	}
	return nil
}
