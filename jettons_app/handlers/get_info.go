package handlers

import (
	// "time"
	// "context"
	"net/http"

	echo "github.com/labstack/echo/v4"

	myStonfiJettons "github.com/Danil-114195722/HamstersShaver/ton_api/stonfi/jettons"
	JettonsErrors "github.com/Danil-114195722/HamstersShaver/jettons_app/errors"
	"github.com/Danil-114195722/HamstersShaver/jettons_app/serializers"
)


// эндпоинт получения информации о монете
func GetInfo(ctx echo.Context) error {
	var err error
	var dataIn serializers.GetInfoIn
	var dataOut myStonfiJettons.JettonParams

	// парсинг JSON-body
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = dataIn.Validate(); err != nil {
		return err
	}

	// // создание контекста с таймаутом в 5 секунд
	// tonApiContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// формирование структуры для ответа
	dataOut, err = myStonfiJettons.GetJettonInfoByAddres(dataIn.MasterAddress)
	if err != nil {
		if err.Error() == "Jetton was not found" {
			return JettonsErrors.InvalidJettonAddressError
		}
		return echo.NewHTTPError(500, map[string]string{"jettons": err.Error()})
	}

	return ctx.JSON(http.StatusOK, dataOut)
}
