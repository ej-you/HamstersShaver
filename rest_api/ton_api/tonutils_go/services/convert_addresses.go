package services

import (
	"fmt"
	"strings"

	tonutilsgoAddress "github.com/xssnick/tonutils-go/address"

	coreErrors "github.com/ej-you/HamstersShaver/rest_api/core/errors"
)


// конвертация адреса из hex в base64 формат
func ConvertAddrToBase64(hexAddr string) (string, error) {
	var apiErr coreErrors.APIError

	// простейшая проверка адреса на hex-формат 
	if !strings.HasPrefix(hexAddr, "0:") {
		apiErr = coreErrors.New(
			fmt.Errorf("invalid hex address was given"),
			"failed to convert addr to base64",
			"rest_api",
			500,
		)
		return "", apiErr
	}
	// парсинг адреса
	addr, err := tonutilsgoAddress.ParseRawAddr(hexAddr)
	if err != nil {
		apiErr = coreErrors.New(
			fmt.Errorf("failed to parse address: %w", err),
			"failed to convert addr to base64",
			"rest_api",
			500,
		)
		return "", apiErr
	}
	return addr.String(), nil
}


// конвертация адреса из base64 в hex формат
func ConvertAddrToHEX(base64Addr string) (string, error) {
	var apiErr coreErrors.APIError

	// простейшая проверка адреса на base64-формат 
	if strings.HasPrefix(base64Addr, "0:") {
		apiErr = coreErrors.New(
			fmt.Errorf("invalid base64 address was given"),
			"failed to convert addr to hex",
			"rest_api",
			500,
		)
		return "", apiErr
	}
	// парсинг адреса
	addr, err := tonutilsgoAddress.ParseAddr(base64Addr)
	if err != nil {
		apiErr = coreErrors.New(
			fmt.Errorf("failed to parse address: %w", err),
			"failed to convert addr to hex",
			"rest_api",
			500,
		)
		return "", apiErr
	}
	return fmt.Sprintf("%v:%x", addr.Workchain(), addr.Data()), nil
}
