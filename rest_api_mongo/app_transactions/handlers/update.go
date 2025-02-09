package handlers

import (
	echo "github.com/labstack/echo/v4"

	"github.com/ej-you/HamstersShaver/rest_api_mongo/mongo"
	"github.com/ej-you/HamstersShaver/rest_api_mongo/mongo/schemas"

	coreValidator "github.com/ej-you/HamstersShaver/rest_api_mongo/core/validator"
)


// структура входных данных для обновления записей транзакций по фильтру
type UpdateIn struct {
	Filter schemas.TransactionFilter `json:"filter" description:"фильтр для выборки записей на обновление"`
	Updater schemas.TransactionUpdater `json:"updater" description:"данные для обновления"`
}

// структура выходных данных для обновления записей транзакций по фильтру
type UpdateOut struct {
	Updated int64 `json:"updated" description:"количество обновлённых записей"`
}


// обновление записей транзакций по фильтру
// @Title Update by filter
// @Description Update transactions in DB by filter (All parameters is optional)
// @Param UpdateParams body UpdateIn true "Данные для обновления"
// @Success 200 object UpdateOut "Количество обновлённых записей монет"
// @Tag transactions
// @Route /transactions [patch]
func Update(ctx echo.Context) error {
	var err error
	var dataIn UpdateIn
	var dataOut UpdateOut

	// парсинг JSON-body
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = coreValidator.GetValidator().Validate(&dataIn); err != nil {
		return err
	}
	// запрещаем менять эти данные в этом ресурсе
	dataIn.Updater.Status = nil
	dataIn.Updater.InitTrans = nil

	// обновление записей транзакций по фильтру
	dataOut.Updated, err = mongo.NewMongoDB().Update(dataIn.Filter, dataIn.Updater)
	if err != nil {
		return err
	}
	return ctx.JSON(200, dataOut)
}
