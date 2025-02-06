package handlers

import (
	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"

	"github.com/ej-you/HamstersShaver/rest_api_mongo/mongo"
	"github.com/ej-you/HamstersShaver/rest_api_mongo/mongo/schemas"

	coreValidator "github.com/ej-you/HamstersShaver/rest_api_mongo/core/validator"
)


// создание записи транзакции auto функции
// @Title Create "auto" transaction document
// @Description Create transaction document in DB with info about "auto" configuration from "auto" bot function
// @Param Transaction body schemas.TransactionAutoCreator true "Данные о транзакции auto функции"
// @Success 200 object schemas.TransactionAutoCreator "Данные созданной записи о транзакции"
// @Tag transactions
// @Route /transactions/create-auto [post]
func CreateAuto(ctx echo.Context) error {
	var err error
	var dataIn schemas.TransactionAutoCreator

	// парсинг JSON-body
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}
	// выставляем значение "auto" в качестве типа транзакции
	dataIn.Type = "auto"
	// генерим uuid для записи
	dataIn.ID = uuid.New()
	// выставляем значение "init" в качестве статуса auto
	dataIn.Status = "init"
	
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
