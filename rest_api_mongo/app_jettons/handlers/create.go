package handlers

import (
	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"

	"github.com/ej-you/HamstersShaver/rest_api_mongo/mongo"
	"github.com/ej-you/HamstersShaver/rest_api_mongo/mongo/schemas"

	coreValidator "github.com/ej-you/HamstersShaver/rest_api_mongo/core/validator"
)


// создание записи монеты
// @Title Create jetton document
// @Description Create jetton document in DB
// @Param Jetton body schemas.Jetton true "Данные о монете"
// @Success 200 object schemas.Jetton "Данные созданной записи о монете"
// @Tag jettons
// @Route /jettons [post]
func Create(ctx echo.Context) error {
	var err error
	var dataIn schemas.Jetton

	// парсинг JSON-body
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}
	// генерим uuid для записи
	dataIn.ID = uuid.New()
	
	// валидация полученной структуры
	if err = coreValidator.GetValidator().Validate(&dataIn); err != nil {
		return err
	}

	// создание записи монеты
	err = mongo.NewMongoDB().Insert(dataIn)
	if err != nil {
		return err
	}
	return ctx.JSON(201, dataIn)
}
