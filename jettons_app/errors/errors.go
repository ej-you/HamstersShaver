package errors

import (
	echo "github.com/labstack/echo/v4"
)


// ошибка при получении информации о монете по её адресу
var InvalidJettonAddressError *echo.HTTPError = echo.NewHTTPError(400, map[string]string{"jettons": "Invalid jetton master address: jetton info was not found"})
