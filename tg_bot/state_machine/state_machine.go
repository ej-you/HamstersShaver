package state_machine

import (
	"fmt"
	"slices"

	customErrors "github.com/ej-you/HamstersShaver/tg_bot/errors"
	"github.com/ej-you/HamstersShaver/tg_bot/redis"
)


// машина состояний юзера
type UserStateMachine struct {
	userTelegramID string
}

// установка тг ID юзера для экземпляра машины состояний
func (this *UserStateMachine) setUserTelegramID(userTelegramID string) {
	if this.userTelegramID == "" {
		this.userTelegramID = userTelegramID
	}
}

// возвращает ошибку, если тг ID юзера для экземпляра машины состояний не установлен
func (this UserStateMachine) errEmptyUserTelegramID() error {
	if this.userTelegramID == "" {
		return fmt.Errorf("use state machine: userTelegramID is not specified: %w", customErrors.InternalError("cannot use state machine"))
	}
	return nil
}

// установка значения в кэш
func (this UserStateMachine) setCacheValue(subkey, value string) error {
	var err error

	if err = this.errEmptyUserTelegramID(); err != nil {
		return err
	}
	key := fmt.Sprintf("user:%s:%s", this.userTelegramID, subkey)

	// вызов redis функции
	if err = redis.SetString(key, value); err != nil {
		return err
	}
	return nil
}


// установка статуса
func (this UserStateMachine) SetStatus(newStatus string) error {
	return this.setCacheValue("status", newStatus)
}

// проверка статуса на совпадение с одним из переданных статусов
func (this UserStateMachine) StatusEquals(otherStatuses ...string) (bool, error) {
	var err error

	if err = this.errEmptyUserTelegramID(); err != nil {
		return false, err
	}
	key := fmt.Sprintf("user:%s:status", this.userTelegramID)

	// получение статуса из redis
	redisStatus, err := redis.GetString(key)
	if err != nil {
		return false, err
	}

	return slices.Contains(otherStatuses, redisStatus), nil
}


// установка действия (buy/cell)
func (this UserStateMachine) SetAction(action string) error {
	return this.setCacheValue("action", action)
}

// установка DEX-биржи (stonfi/dedust)
func (this UserStateMachine) SetDEX(dex string) error {
	return this.setCacheValue("dex", dex)
}

// установка кол-ва монет для транзакции
func (this UserStateMachine) SetJettonsAmount(jettonsAmount string) error {
	return this.setCacheValue("jettonsAmount", jettonsAmount)
}

// установка процента проскальзывания
func (this UserStateMachine) SetSlippage(slippage string) error {
	return this.setCacheValue("slippage", slippage)
}

// установка CA монеты для покупки/продажи
func (this UserStateMachine) SetJettonCA(jettonCA string) error {
	return this.setCacheValue("jettonCA", jettonCA)
}

// получение CA монеты для покупки/продажи
func (this UserStateMachine) GetJettonCA() (string, error) {
	var err error
	var jettonCA string

	if err = this.errEmptyUserTelegramID(); err != nil {
		return "", fmt.Errorf("get jettonCA: %w", err)
	}
	
	jettonCA, err = redis.GetString(fmt.Sprintf("user:%s:jettonCA", this.userTelegramID))
	if err != nil {
		return "", fmt.Errorf("get new transaction preparation: %w", err)
	}
	return jettonCA, nil
}