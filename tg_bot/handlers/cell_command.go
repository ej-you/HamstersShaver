package handlers

import (
	telebot "gopkg.in/telebot.v3"

	"github.com/ej-you/HamstersShaver/tg_bot/services"
	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
)


// команды: /cell
func CellHandlerCommand(context telebot.Context) error {
	var err error
	userId := services.GetUserID(context.Chat())

	// получение машины состоянию текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(userId)

	// игнорирование сообщения при несоответствии статуса
	accepted, err := cellStatusIsAccepted(userStateMachine)
	if err != nil {
		return err
	}
	if !accepted {
		return nil
	}

	// установка нового состояния
	if err = userStateMachine.SetStatus("cell"); err != nil {
		return err
	}

	msgText := `📉 Активирован диалог продажи монет. Для отмены всех действий и выхода в главное меню используйте /cancel

Выберите из имеющихся у вас на аккаунте монет ту, которую хотите продать 👇`

	return context.Send(msgText)
}


// возвращает true, если при текущем статусе можно использовать данную функцию
func cellStatusIsAccepted(userStateMachine stateMachine.UserStateMachine) (bool, error) {
	equals, err := userStateMachine.StatusEquals("trade")
	if err != nil {
		return false, err
	}
	if equals {
		return true, nil
	}

	equals, err = userStateMachine.StatusEquals("home")
	if err != nil {
		return false, err
	}
	if equals {
		return true, nil
	}

	equals, err = userStateMachine.StatusEquals("start")
	if err != nil {
		return false, err
	}
	if equals {
		return true, nil
	}

	return false, nil
}


// кнопки: to_cell
func CellHandlerCallback(context telebot.Context) error {
	var err error
	userId := services.GetUserID(context.Chat())

	// получение машины состоянию текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(userId)

	// игнорирование сообщения при несоответствии статуса
	equals, err := userStateMachine.StatusEquals("trade")
	if err != nil {
		return err
	}
	if !equals {
		return nil
	}

	// установка нового состояния
	if err = userStateMachine.SetStatus("cell"); err != nil {
		return err
	}
	// установка значения действия
	if err = userStateMachine.SetAction("cell"); err != nil {
		return err
	}

	msgText := `Отлично! Выбрано действие продажи монет 📉

Выберите из имеющихся у вас на аккаунте монет ту, которую хотите продать 👇`

	return context.Send(msgText)
}