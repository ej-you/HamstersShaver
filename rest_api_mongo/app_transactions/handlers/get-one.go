package handlers

import (
	"errors"

	echo "github.com/labstack/echo/v4"

	"github.com/ej-you/HamstersShaver/rest_api_mongo/mongo"
	"github.com/ej-you/HamstersShaver/rest_api_mongo/mongo/schemas"

	coreValidator "github.com/ej-you/HamstersShaver/rest_api_mongo/core/validator"
)


// получение одной записи транзакции по фильтру
// @Title Get one by filter
// @Description Get one transaction from DB by filter (All parameters is optional)
// @Param ID query string false "UUID записи" "715c0b81-bf1b-46c4-bf08-5c137cc6ec4d"
// @Param Hash query string false "Hash первой операции транзакции" "009f801c3ab128fb53e5fca0ffe47b2dcfec3f6e28a07cf992ace5297363b72f"
// @Success 200 object schemas.Transaction "Запись транзакции, подходящей под фильтр"
// @Success 404 "Если с данным фильтром запись монеты не была найдена"
// @Tag transactions
// @Route /transactions [get]
func GetOne(ctx echo.Context) error {
	var err error
	var dataIn schemas.TransactionFilter
	var dataOut schemas.Transaction

	// парсинг query-параметров
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = coreValidator.GetValidator().Validate(&dataIn); err != nil {
		return err
	}

	// получение записи транзакции
	err = mongo.NewMongoDB().GetOneByFilter(dataIn, &dataOut)
	if err != nil {
		// если запись не найдена
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(404, map[string]string{"transactions": err.Error()})
		}
		return err
	}
	return ctx.JSON(200, dataOut)
}
