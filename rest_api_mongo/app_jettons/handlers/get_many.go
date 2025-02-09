package handlers

import (
	echo "github.com/labstack/echo/v4"

	"github.com/ej-you/HamstersShaver/rest_api_mongo/mongo"
	"github.com/ej-you/HamstersShaver/rest_api_mongo/mongo/schemas"

	coreValidator "github.com/ej-you/HamstersShaver/rest_api_mongo/core/validator"
	"github.com/ej-you/HamstersShaver/rest_api_mongo/settings"
)


// получение записей монет по фильтру
// @Title Get many by filter
// @Description Get jettons list from DB by filter (All parameters is optional)
// @Param ID query string false "UUID записи" "715c0b81-bf1b-46c4-bf08-5c137cc6ec4d"
// @Param Symbol query string false "Название монеты" "GRAM"
// @Param JettonCA query string false "Мастер-адрес монеты (jetton_master)" "EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O"
// @Param DEX query DEXesEnum false "DEX-биржа" "Ston.fi"
// @Success 200 array []schemas.Jetton "Список записей монет, подходящих под фильтр"
// @Success 204 "Пустой ответ, если не найдено ни одной записи"
// @Tag jettons
// @Route /jettons/get-many [get]
func GetMany(ctx echo.Context) error {
	var err error
	var dataIn schemas.JettonFilter
	var dataOut []schemas.Jetton

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
