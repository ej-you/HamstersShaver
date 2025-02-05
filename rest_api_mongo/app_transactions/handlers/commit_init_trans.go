package handlers

import (
	"fmt"
	"errors"

	echo "github.com/labstack/echo/v4"

	"github.com/ej-you/HamstersShaver/rest_api_mongo/mongo"
	"github.com/ej-you/HamstersShaver/rest_api_mongo/mongo/schemas"

	coreValidator "github.com/ej-you/HamstersShaver/rest_api_mongo/core/validator"
)


// фиксирование информации о закупочной транзакции auto конфигурации в подструктуре InitTransactionInfo
// @Title Commit init trans
// @Description Commit init transaction info into initTrans sub-object
// @Param TransactionFilter body schemas.TransactionFilter true "Фильтр для получения транзакции"
// @Success 200 object schemas.Transaction "Обновлённая запись транзакции"
// @Tag transactions
// @Route /transactions/commit-init-trans [patch]
func CommitInitTrans(ctx echo.Context) error {
	var err error
	var dataIn schemas.TransactionFilter
	var dataOut schemas.Transaction

	// парсинг JSON-body
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}	
	// валидация полученной структуры
	if err = coreValidator.GetValidator().Validate(&dataIn); err != nil {
		return err
	}
	// делаем проверку, что хотя бы один параметр фильтра был введён
	if dataIn.ID == nil && dataIn.Hash == nil {
		return echo.NewHTTPError(400, map[string]string{"validateError": "At least one filter parameter must be presented"})
	}

	// получение записи транзакции
	err = mongo.NewMongoDB().GetOneByFilter(dataIn, &dataOut)
	if err != nil {
		// если запись не найдена
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(404, map[string]string{"transactions": fmt.Sprintf("commit init trans: %v", err)})
		}
		return err
	}

	// проверка необходимых полей
	if dataOut.Status != "init" {
		return echo.NewHTTPError(400, map[string]string{"transactions": fmt.Sprintf("commit init trans: transaction status is not init")})
	}
	if dataOut.Finished != true {
		return echo.NewHTTPError(400, map[string]string{"transactions": fmt.Sprintf("commit init trans: transaction is not finished")})
	}
	if dataOut.Success != true {
		return echo.NewHTTPError(400, map[string]string{"transactions": fmt.Sprintf("commit init trans: transaction did not finished successfully")})
	}
	if dataOut.Error != false {
		return echo.NewHTTPError(400, map[string]string{"transactions": fmt.Sprintf("commit init trans: transaction was finished with error")})
	}
	if dataOut.LastTxHash == "" {
		return echo.NewHTTPError(400, map[string]string{"transactions": fmt.Sprintf("commit init trans: LastTxHash is not set")})
	}

	// данные для обновления записи
	dataOut.Status = "auto"
	dataOut.InitTrans = &schemas.InitTransactionInfo{dataOut.Hash, dataOut.LastTxHash}
	dataOut.Hash = "waiting"
	dataOut.Finished = false
	dataOut.Success = false
	dataOut.Error = false
	dataOut.LastTxHash = ""
	// структура для обновления записи
	transUpdater := schemas.TransactionUpdater{
		Status: &dataOut.Status,
		InitTrans: dataOut.InitTrans,
		Hash: &dataOut.Hash,
		Finished: &dataOut.Finished,
		Success: &dataOut.Success,
		Error: &dataOut.Error,
		LastTxHash: &dataOut.LastTxHash,
	}

	// обновление записи транзакции по фильтру
	updatedAmount, err := mongo.NewMongoDB().Update(dataIn, transUpdater)
	if err != nil {
		return err
	}
	// если не обновилась ни одна запись
	if updatedAmount == 0 {
		return fmt.Errorf("commit init trans: update: transaction was not updated")
	}

	return ctx.JSON(200, dataOut)
}
