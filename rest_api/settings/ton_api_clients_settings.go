package settings

import (
	"time"
	"errors"

	tongo "github.com/tonkeeper/tongo/liteapi"
	tonapi "github.com/tonkeeper/tonapi-go"
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

	// тестовый конфиг
	if conType == "testnet" {
		client, err := tonapi.NewClient(tonapi.TestnetTonApiURL)
		if err != nil {
			return tonClientTonapi{Client: client, Error: err}
		}
		// InfoLog.Println("(tonapi) Connected to testnet TON node")
		return tonClientTonapi{Client: client, Error: nil}

	// основной конфиг
	} else if conType == "mainnet" {
		client, err := tonapi.New()
		if err != nil {
			return tonClientTonapi{Client: client, Error: err}
		}
		// InfoLog.Println("(tonapi) Connected to mainnet TON node")
		return tonClientTonapi{Client: client, Error: nil}

	// неправильный параметр конфига
	} else {
		conTypeError := errors.New("(tonapi) Invalid conType parameter was given")
		ErrorLog.Println("Failed to get tonapi client:", conTypeError.Error())
		return tonClientTonapi{Client: client, Error: conTypeError}
	}
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
			return result.Client, result.Error
		// если прошло время timeout, а данные не получены, то возвращаем ошибку таймаута
		case <- time.After(timeout):
			timeoutError := errors.New("Failed to get tonapi-go client: timeout error")
			ErrorLog.Println(timeoutError.Error())
			var emptyClient *tonapi.Client
			return emptyClient, timeoutError
	}
}


// создание клиента TON для tongo
func getTonClientTongo(conType string) tonClientTongo {
	var client *tongo.Client

	// тестовый конфиг
	if conType == "testnet" {
		client, err := tongo.NewClientWithDefaultTestnet()
		if err != nil {
			return tonClientTongo{Client: client, Error: err}
		}
		// InfoLog.Println("(tongo) Connected to testnet TON node")
		return tonClientTongo{Client: client, Error: nil}

	// основной конфиг
	} else if conType == "mainnet" {
		client, err := tongo.NewClientWithDefaultMainnet()
		if err != nil {
			return tonClientTongo{Client: client, Error: err}
		}
		// InfoLog.Println("(tongo) Connected to mainnet TON node")
		return tonClientTongo{Client: client, Error: nil}

	// неправильный параметр конфига
	} else {
		conTypeError := errors.New("(tongo) Invalid conType parameter was given")
		ErrorLog.Println("Failed to get tongo client:", conTypeError.Error())
		return tonClientTongo{Client: client, Error: conTypeError}
	}
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
			return result.Client, result.Error
		// если прошло время timeout, а данные не получены, то возвращаем ошибку таймаута
		case <- time.After(timeout):
			timeoutError := errors.New("Failed to get tongo client: timeout error")
			ErrorLog.Println(timeoutError.Error())
			var emptyClient *tongo.Client
			return emptyClient, timeoutError
	}
}
