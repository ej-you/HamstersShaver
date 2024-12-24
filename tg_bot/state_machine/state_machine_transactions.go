package state_machine

import (
	"fmt"
	"slices"

	"github.com/pkg/errors"

	"github.com/ej-you/HamstersShaver/tg_bot/redis"
)


// собранные данные из диалога для создания новой транзакции
type NewTransactionPreparation struct {
	Action string
	DEX string
	Amount string
	Slippage string
	JettonCA string
}

// получение всей информации для новой транзакции
func (this UserStateMachine) GetNewTransactionPreparation() (NewTransactionPreparation, error) {
	var err error
	var transInfo NewTransactionPreparation

	if err = this.errEmptyUserTelegramID(); err != nil {
		return transInfo, fmt.Errorf("get new transaction preparation: %w", err)
	}
	keyPrefix := fmt.Sprintf("user:%s:", this.userTelegramID)
	
	transInfo.Action, err = redis.GetString(keyPrefix + "action")
	if err != nil {
		return transInfo, fmt.Errorf("get new transaction preparation: %w", err)
	}
	transInfo.DEX, err = redis.GetString(keyPrefix + "dex")
	if err != nil {
		return transInfo, fmt.Errorf("get new transaction preparation: %w", err)
	}
	transInfo.Amount, err = redis.GetString(keyPrefix + "jettonsAmount")
	if err != nil {
		return transInfo, fmt.Errorf("get new transaction preparation: %w", err)
	}
	transInfo.Slippage, err = redis.GetString(keyPrefix + "slippage")
	if err != nil {
		return transInfo, fmt.Errorf("get new transaction preparation: %w", err)
	}
	transInfo.JettonCA, err = redis.GetString(keyPrefix + "jettonCA")
	if err != nil {
		return transInfo, fmt.Errorf("get new transaction preparation: %w", err)
	}

	return transInfo, nil
}

// удаление всех ключей со значениями для новой транзакции
func (this UserStateMachine) ClearNewTransactionPreparation() error {
	var err error

	if err = this.errEmptyUserTelegramID(); err != nil {
		return fmt.Errorf("clear new transaction preparation: %w", err)
	}
	keyPrefix := fmt.Sprintf("user:%s:", this.userTelegramID)

	err = redis.DeleteKey(keyPrefix + "action")
	if err != nil {
		return fmt.Errorf("clear new transaction preparation: %w", err)
	}
	err = redis.DeleteKey(keyPrefix + "dex")
	if err != nil {
		return fmt.Errorf("clear new transaction preparation: %w", err)
	}
	err = redis.DeleteKey(keyPrefix + "jettonsAmount")
	if err != nil {
		return fmt.Errorf("clear new transaction preparation: %w", err)
	}
	err = redis.DeleteKey(keyPrefix + "slippage")
	if err != nil {
		return fmt.Errorf("clear new transaction preparation: %w", err)
	}
	err = redis.DeleteKey(keyPrefix + "jettonCA")
	if err != nil {
		return fmt.Errorf("clear new transaction preparation: %w", err)
	}
	return nil
}


// получение транзакций в списке ожидания
func (this UserStateMachine) GetPendingTransactions() ([]string, error) {
	var err error
	var pendingTransactions []string

	if err = this.errEmptyUserTelegramID(); err != nil {
		return pendingTransactions, fmt.Errorf("get pending transactions: %w", err)
	}
	key := fmt.Sprintf("user:%s:transactions", this.userTelegramID)

	// получение существующего среза транзакций из redis
	pendingTransactions, err = redis.GetStringSlice(key)
	
	// если такого ключа в кэше ещё нет
	nilKeyError := errors.New("value: redis: nil: failed to get string slice value")
	if err != nil && errors.As(errors.Unwrap(err), &nilKeyError) {
		return pendingTransactions, nil
	}
	// неизвестная ошибка
	if err != nil {
		return pendingTransactions, fmt.Errorf("get pending transactions: %w", err)
	}
	return pendingTransactions, nil
}

// добавление транзакции в список ожидания
func (this UserStateMachine) AddPendingTransaction(transactionUUID string) error {
	var err error

	if err = this.errEmptyUserTelegramID(); err != nil {
		return fmt.Errorf("add pending transactions: %w", err)
	}
	key := fmt.Sprintf("user:%s:transactions", this.userTelegramID)

	// получение существующего среза транзакций из redis
	redisTransactions, err := redis.GetStringSlice(key)

	// если такого ключа в кэше ещё нет
	nilKeyError := errors.New("value: redis: nil: failed to get string slice value")
	if err != nil && errors.As(errors.Unwrap(err), &nilKeyError) {
		err = redis.SetStringSlice(key, []string{transactionUUID})
	}
	// неизвестная ошибка
	if err != nil {
		return fmt.Errorf("add pending transactions: %w", err)
	}

	// добавление UUID новой транзакции в срез транзакций из redis и отправка нового среза на установку в redis
	if err = redis.SetStringSlice(key, append(redisTransactions, transactionUUID)); err != nil {
		return fmt.Errorf("add pending transactions: %w", err)
	}
	return nil
}

// удаление транзакции из списка ожидания
func (this UserStateMachine) DeletePendingTransaction(transactionUUID string) error {
	var err error

	if err = this.errEmptyUserTelegramID(); err != nil {
		return fmt.Errorf("delete pending transactions: %w", err)
	}
	key := fmt.Sprintf("user:%s:transactions", this.userTelegramID)

	// получение существующего среза транзакций из redis
	redisTransactions, err := redis.GetStringSlice(key)
	// если такого ключа в кэше нет

	nilKeyError := errors.New("value: redis: nil: failed to get string slice value")
	if err != nil && errors.As(errors.Unwrap(err), &nilKeyError) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("delete pending transactions: %w", err)
	}

	transactionToDeleteIdx := slices.Index(redisTransactions, transactionUUID)
	// если такой UUID не найден в срезе транзакций из redis
	if transactionToDeleteIdx == -1 {
		return nil
	}

	// если найден, то удаляем его оттуда и обновляем значение ключа в redis
	updatedTransactionsSlice := slices.Delete(redisTransactions, transactionToDeleteIdx, transactionToDeleteIdx+1)
	if len(updatedTransactionsSlice) == 0 {
		if err = redis.DeleteKey(key); err != nil {
			return fmt.Errorf("delete pending transactions: %w", err)
		}
		return nil
	}

	if err = redis.SetStringSlice(key, updatedTransactionsSlice); err != nil {
		return fmt.Errorf("delete pending transactions: %w", err)
	}

	return nil
}
