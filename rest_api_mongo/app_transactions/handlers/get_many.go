package handlers

import (
	echo "github.com/labstack/echo/v4"

	"github.com/ej-you/HamstersShaver/rest_api_mongo/mongo"
	"github.com/ej-you/HamstersShaver/rest_api_mongo/mongo/schemas"

	coreValidator "github.com/ej-you/HamstersShaver/rest_api_mongo/core/validator"
	"github.com/ej-you/HamstersShaver/rest_api_mongo/settings"
)


// получение записей транзакций по фильтру
// @Title Get many by filter
// @Description Get transactions list from DB by filter (All parameters is optional)
// @Param ID query string false "UUID записи" "715c0b81-bf1b-46c4-bf08-5c137cc6ec4d"
// @Param Hash query string false "Хэш первой операции цепочки транзакций" "009f801c3ab128fb53e5fca0ffe47b2dcfec3f6e28a07cf992ace5297363b72f"
// @Param Type query TypesEnum false "Тип транзакции" "auto"
// @Param Finished query bool false "Завершена ли транзакция" "true"
// @Success 200 array []schemas.Transaction "Список записей транзакций, подходящих под фильтр"
// @Success 204 "Пустой ответ, если не найдено ни одной записи"
// @Tag transactions
// @Route /transactions/get-many [get]
func GetMany(ctx echo.Context) error {
	var err error
	var dataIn schemas.TransactionFilter
	var dataOut []schemas.Transaction

	// парсинг query-параметров
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = coreValidator.GetValidator().Validate(&dataIn); err != nil {
		return err
	}

	// получение списка монет
	err = mongo.NewMongoDB().GetManyByFilter(dataIn, &dataOut)
	if err != nil {
		return err
	}
	// если не найдено ни одной записи
	if len(dataOut) == 0 {
		return ctx.NoContent(204)
	}

	// настройка временной зоны для каждой записи
	for i, _ := range dataOut {
		dataOut[i].CreatedAt = dataOut[i].CreatedAt.In(settings.TimeZone)
	}
	return ctx.JSON(200, dataOut)
}
