package services

import (
	"fmt"

	tonutilsgoAddress "github.com/xssnick/tonutils-go/address"

	coreErrors "github.com/ej-you/HamstersShaver/rest_api/core/errors"
)


// конвертация адреса из hex в base64 формат
func ConvertAddrToBase64(hexAddr string) (string, error) {
	// парсинг адреса
	addr, err := tonutilsgoAddress.ParseRawAddr(hexAddr)
	if err != nil {
		return "", fmt.Errorf("convert addr to base64: failed to parse address: %v: %w", err, coreErrors.RestApiError)
	}
	return addr.String(), nil
}


// конвертация адреса из base64 в hex формат
func ConvertAddrToHEX(base64Addr string) (string, error) {
	// парсинг адреса
	addr, err := tonutilsgoAddress.ParseAddr(base64Addr)
	if err != nil {
		return "", fmt.Errorf("convert addr to hex: failed to parse address: %v: %w", err, coreErrors.RestApiError)
	}
	return fmt.Sprintf("%v:%x", addr.Workchain(), addr.Data()), nil
}
