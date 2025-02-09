package handlers

import (
	"errors"

	echo "github.com/labstack/echo/v4"

	"github.com/ej-you/HamstersShaver/rest_api_mongo/mongo"
	"github.com/ej-you/HamstersShaver/rest_api_mongo/mongo/schemas"

	coreValidator "github.com/ej-you/HamstersShaver/rest_api_mongo/core/validator"
	"github.com/ej-you/HamstersShaver/rest_api_mongo/settings"
)


// получение одной записи монеты по фильтру
// @Title Get one by filter
// @Description Get one jetton from DB by filter (All parameters is optional)
// @Param ID query string false "UUID записи" "715c0b81-bf1b-46c4-bf08-5c137cc6ec4d"
// @Param Symbol query string false "Название монеты" "GRAM"
// @Param JettonCA query string false "мастер-адрес монеты (jetton_master)" "EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O"
// @Param DEX query DEXesEnum false "DEX-биржа" "Ston.fi"
// @Success 200 object schemas.Jetton "Запись монеты, подходящей под фильтр"
// @Success 404 "Если с данным фильтром запись монеты не была найдена"
// @Tag jettons
// @Route /jettons/get-one [get]
func GetOne(ctx echo.Context) error {
	var err error
	var dataIn schemas.JettonFilter
	var dataOut schemas.Jetton

	// парсинг query-параметров
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = coreValidator.GetValidator().Validate(&dataIn); err != nil {
		return err
	}

	// получение записи монеты
	err = mongo.NewMongoDB().GetOneByFilter(dataIn, &dataOut)
	if err != nil {
		// если запись не найдена
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(404, map[string]string{"jettons": err.Error()})
		}
		return err
	}
	// настройка временной зоны
	dataOut.CreatedAt = dataOut.CreatedAt.In(settings.TimeZone)

	return ctx.JSON(200, dataOut)
}
