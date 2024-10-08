package errors

import (
	echo "github.com/labstack/echo/v4"
)


// ошибка при получении информации о монете аккаунта (у аккаунта нет такой монеты)
var AccountHasNotJettonError *echo.HTTPError = echo.NewHTTPError(404, map[string]string{"account": "Account has not given jetton"})

// ошибка при получении информации о монете аккаунта (дан неверный адрес)
var InvalidJettonAddressError *echo.HTTPError = echo.NewHTTPError(400, map[string]string{"account": "Failed to decode jetton: invalid address was given"})
