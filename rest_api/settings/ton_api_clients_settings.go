package settings

import (
	"fmt"
	"time"

	tongo "github.com/tonkeeper/tongo/liteapi"
	tonapi "github.com/tonkeeper/tonapi-go"

	coreErrors "github.com/ej-you/HamstersShaver/rest_api/core/errors"
)


// структура для получения результата от горутины в функции с таймаутом
type tonClientTonapi struct {
	Client *tonapi.Client
	Error  error
}
// структура для получения результата от горутины в функции с таймаутом
type tonClientTongo struct {
	Client *tongo.Client
	Error  error
}


// создание клиента TON для tonapi-go
func getTonClientTonapi(conType string) tonClientTonapi {
	var client *tonapi.Client
	var err error

	// тестовый конфиг
	if conType == "testnet" {
		client, err = tonapi.NewClient(tonapi.TestnetTonApiURL)
	// основной конфиг
	} else if conType == "mainnet" {
		client, err = tonapi.New()
	// неправильный параметр конфига
	} else {
		panic("(tonapi) Invalid conType parameter was given")
	}

	if err != nil {
		return tonClientTonapi{Client: client, Error: err}
	}
	return tonClientTonapi{Client: client, Error: nil}
}


// создание клиента TON для tonapi-go с таймаутом
func GetTonClientTonapiWithTimeout(conType string, timeout time.Duration) (*tonapi.Client, error) {
	// если таймаут равен 0
	if timeout == 0 {
		result := getTonClientTonapi(conType)
		return result.Client, result.Error
	}

	// создание небуферизированного канала
	ch := make(chan tonClientTonapi)
	// вызов горутины
	go func() {
		ch <- getTonClientTonapi(conType)
	}()

	select {
		// если данные получены, то возвращаем их
		case result := <- ch:
			if result.Error != nil {
				return result.Client, fmt.Errorf("failed to get tonapi-go client: %v: %w", result.Error, coreErrors.TonApiError)
			}
			return result.Client, nil
		// если прошло время timeout, а данные не получены, то возвращаем ошибку таймаута
		case <- time.After(timeout):
			return nil, fmt.Errorf("failed to get tonapi-go client: %w", coreErrors.TimeoutError)
	}
}


// создание клиента TON для tongo
func getTonClientTongo(conType string) tonClientTongo {
	var client *tongo.Client
	var err error

	// тестовый конфиг
	if conType == "testnet" {
		client, err = tongo.NewClientWithDefaultTestnet()
	// основной конфиг
	} else if conType == "mainnet" {
		client, err = tongo.NewClientWithDefaultMainnet()
	// неправильный параметр конфига
	} else {
		panic("(tongo) Invalid conType parameter was given")
	}

	if err != nil {
		return tonClientTongo{Client: client, Error: err}
	}
	return tonClientTongo{Client: client, Error: nil}	
}

// создание клиента TON для tongo с таймаутом
func GetTonClientTongoWithTimeout(conType string, timeout time.Duration) (*tongo.Client, error) {
	// если таймаут равен 0
	if timeout == 0 {
		result := getTonClientTongo(conType)
		return result.Client, result.Error
	}

	// создание небуферизированного канала
	ch := make(chan tonClientTongo)
	// вызов горутины
	go func() {
		ch <- getTonClientTongo(conType)
	}()

	select {
		// если данные получены, то возвращаем их
		case result := <- ch:
			if result.Error != nil {
				return result.Client, fmt.Errorf("failed to get tongo client: %v: %w", result.Error, coreErrors.TonApiError)
			}
			return result.Client, nil
		// если прошло время timeout, а данные не получены, то возвращаем ошибку таймаута
		case <- time.After(timeout):
			return nil, fmt.Errorf("failed to get tongo client: %w", coreErrors.TimeoutError)
	}
}
