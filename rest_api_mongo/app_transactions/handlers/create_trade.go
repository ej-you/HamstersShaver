package handlers

import (
	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"

	"github.com/ej-you/HamstersShaver/rest_api_mongo/mongo"
	"github.com/ej-you/HamstersShaver/rest_api_mongo/mongo/schemas"

	coreValidator "github.com/ej-you/HamstersShaver/rest_api_mongo/core/validator"
)


// создание записи транзакции trade функции
// @Title Create "trade" transaction document
// @Description Create transaction document in DB with info about transaction from "trade" bot function
// @Param Transaction body schemas.TransactionCreator true "Данные о транзакции trade функции"
// @Success 201 object schemas.TransactionCreator "Данные созданной записи о транзакции"
// @Tag transactions
// @Route /transactions/create-trade [post]
func CreateTrade(ctx echo.Context) error {
	var err error
	var dataIn schemas.TransactionCreator

	// парсинг JSON-body
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}
	// выставляем значение "trade" в качестве типа транзакции
	dataIn.Type = "trade"
	// генерим uuid для записи
	dataIn.ID = uuid.New()
	
	// валидация полученной структуры
	if err = coreValidator.GetValidator().Validate(&dataIn); err != nil {
		return err
	}

	// создание записи транзакции
	err = mongo.NewMongoDB().Insert(dataIn)
	if err != nil {
		return err
	}
	return ctx.JSON(201, dataIn)
}
