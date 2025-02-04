package handlers

import (
	echo "github.com/labstack/echo/v4"

	"github.com/ej-you/HamstersShaver/rest_api_mongo/mongo"
	"github.com/ej-you/HamstersShaver/rest_api_mongo/mongo/schemas"

	coreValidator "github.com/ej-you/HamstersShaver/rest_api_mongo/core/validator"
)


// структура выходных данных для удаления записей монет по фильтру
type DeleteOut struct {
	Deleted int64 `json:"deleted" description:"количество удалённых записей"`
}


// удаление записей монет по фильтру
// @Title Delete by filter
// @Description Delete jettons from DB by filter
// @Param JettonFilter body schemas.JettonFilter false "Фильтр для удаления записей монет"
// @Success 200 object DeleteOut "Количество удалённых записей монет"
// @Tag jettons
// @Route /jettons [delete]
func Delete(ctx echo.Context) error {
	var err error
	var dataIn schemas.JettonFilter
	var dataOut DeleteOut

	// парсинг JSON-body
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}	
	// валидация полученной структуры
	if err = coreValidator.GetValidator().Validate(&dataIn); err != nil {
		return err
	}

	// удаление записей монет по фильтру
	dataOut.Deleted, err = mongo.NewMongoDB().Delete(dataIn)
	if err != nil {
		return err
	}

	return ctx.JSON(200, dataOut)
}
