package handlers

import (
	echo "github.com/labstack/echo/v4"

	coreValidator "github.com/ej-you/HamstersShaver/rest_api_mongo/core/validator"
	// "github.com/ej-you/HamstersShaver/rest_api_mongo/settings"
)


// // структура входных данных для получения информации о последующей транзакции покупки
// type BuyPreRequestIn struct {
// 	JettonCA string `query:"jettonCA" json:"jettonCA" validate:"required"`
// 	Amount float64 `query:"amount" json:"amount" validate:"required"`
// 	Slippage int `query:"slippage" json:"slippage" validate:"required,min=1,max=100"`
// }

type Test struct {
	User 	string `json:"user" validate:"required" example:"user1" description:"user login"`
	Action 	string `json:"action" validate:"required,oneof=buy cell" example:"buy" description:"action [buy OR cell]"`
}


// эндпоинт создания новой записи транзакции
// @Title Add new transaction
// @Description Add new transaction to mongo
// @Param Test body Test true "Cтруктура входных данных для создания записи транзакции в mongo"
// @Success 200 object Test "New transaction JSON"
// @Tag transactions
// @Route /transactions/new [post]
func New(ctx echo.Context) error {
	var err error
	var dataIn Test
	// var dataOut myTongoTransactions.PreRequestBuyJetton

	// парсинг query-параметров
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = coreValidator.GetValidator().Validate(&dataIn); err != nil {
		return err
	}

	// // получение примерных данных о будующей транзакции
	// dataOut, err = myTongoTransactions.GetPreRequestBuyJetton(dataIn.JettonCA, dataIn.Amount, dataIn.Slippage)
	// if err != nil {
	// 	settings.ErrorLog.Println(err)
	// 	return err
	// }

	return ctx.JSON(201, dataIn)
}
