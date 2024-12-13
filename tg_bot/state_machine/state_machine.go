package state_machine

import (
	"errors"
	"fmt"

	"github.com/ej-you/HamstersShaver/tg_bot/redis"
)


type UserStateMachine struct {
	userTelegramID string
}

// установка тг ID юзера для экземпляра машины состояний
func (this *UserStateMachine) SetUserTelegramID(userTelegramID string) {
	if this.userTelegramID == "" {
		this.userTelegramID = userTelegramID
	}
}

// возвращает ошибку, если тг ID юзера для экземпляра машины состояний не установлен
func (this UserStateMachine) errEmptyUserTelegramID() error {
	if this.userTelegramID == "" {
		return errors.New("userTelegramID is not specified")
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

// проверка статуса на совпадение с переданным статусом
func (this UserStateMachine) StatusEquals(otherStatus string) (bool, error) {
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

	return otherStatus == redisStatus, nil
}


// установка действия (buy/cell)
func (this UserStateMachine) SetAction(action string) error {
	return this.setCacheValue("action", action)
}

// установка DEX-биржи (Stonfi/Dedust)
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
