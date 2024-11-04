package errors

import (
	echo "github.com/labstack/echo/v4"
)


// ошибка при создании клиента для tonapi-go
var GetTonapiClientError *echo.HTTPError = echo.NewHTTPError(500, map[string]string{"tonapiClient": "Failed to get client for tonapi-go"})
// ошибка при создании клиента для tongo
var GetTongoClientError *echo.HTTPError = echo.NewHTTPError(500, map[string]string{"tongoClient": "Failed to get client for tongo"})
