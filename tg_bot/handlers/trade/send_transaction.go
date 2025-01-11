package trade

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	telebot "gopkg.in/telebot.v3"

	"github.com/ej-you/HamstersShaver/tg_bot/mongo"
	mongoSchemas "github.com/ej-you/HamstersShaver/tg_bot/mongo/schemas"
	
	backgroundTrading "github.com/ej-you/HamstersShaver/tg_bot/background/trading"
	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
	
	customErrors "github.com/ej-you/HamstersShaver/tg_bot/errors"
	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
)


// отправка подготовленной транзакции
func SendTransaction(context telebot.Context) error {
	var err error

	mongoDB := mongo.NewMongoDB()
	// получение информации о последней транзакции (из БД)
	var lastTransInfo mongoSchemas.Transaction
	err = mongoDB.GetLastTransaction(&lastTransInfo)
	if err != nil {
		dbNotFoundErr := new(customErrors.DBNotFoundError)
		// если ошибка не потому, что не было найдено ни одной транзакции
		if !errors.As(err, dbNotFoundErr) {
			return fmt.Errorf("SendTransaction: %w", err)
		}
	}
	// если последняя транзакция ещё не завершена (и нет ошибки ненахождения ни одной транзакции)
	if err == nil && lastTransInfo.Finished == false { // NOT err
		return customErrors.LastTransNotFinishedError("last transaction not finished")
	}

	userId := services.GetUserID(context.Chat())
	// получение машины состояний текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(userId)
	// установка нового состояния
	if err = userStateMachine.SetStatus("send_transaction"); err != nil {
		return fmt.Errorf("SendTransaction: %w", err)
	}

	// получение данных для новой транзакции
	newTrans, err := userStateMachine.GetNewTransactionPreparation()
	if err != nil {
		return fmt.Errorf("SendTransaction: %w", err)
	}

	// сбор информации о новой транзакции для добавления в БД
	newDBTransaction := mongoSchemas.Transaction{
		UserID: userId,
		ID: uuid.New(),
		Action: newTrans.Action,
		DEX: newTrans.DEX,
		JettonCA: newTrans.JettonCA,
		Finished: false,
	}
	if newTrans.Action == "cell" {
		newDBTransaction.UsedJettons = newTrans.Amount
	} else {
		newDBTransaction.UsedTON = newTrans.Amount
	}
	// добавление информации о новой транзакции в БД
	if err := mongoDB.InsertOne(newDBTransaction); err != nil {
		return fmt.Errorf("SendTransaction: %w", err)
	}

	// сообщение о начале обработки транзакции
	transMsg := fmt.Sprintf("▶️ Начата обработка транзакции выше... 👆\n(ID: %s)", newDBTransaction.ID.String())
	sentTransMsg, err := context.Bot().Send(context.Recipient(), transMsg, keyboards.InlineKeyboardToHome)
	if err != nil {
		return fmt.Errorf("SendTransaction: %w", err)
	}

	// запуск обработки в фоне
	go backgroundTrading.ProcessTransaction(&context, sentTransMsg, newTrans, newDBTransaction.ID)
	return nil
}
