package services

import (
	"errors"
	"fmt"
	"strings"

	tonutilsgoAddress "github.com/xssnick/tonutils-go/address"
)


// конвертация адреса из hex в base64 формат
func ConvertAddrToBase64(hexAddr string) (string, error) {
	// простейшая проверка адреса на hex-формат 
	if !strings.HasPrefix(hexAddr, "0:") {
		return "", errors.New("invalid hex address was given")
	}
	// парсинг адреса
	addr, err := tonutilsgoAddress.ParseRawAddr(hexAddr)
	if err != nil {
		return "", err
	}
	return addr.String(), nil
}


// конвертация адреса из base64 в hex формат
func ConvertAddrToHEX(base64Addr string) (string, error) {
	// простейшая проверка адреса на base64-формат 
	if strings.HasPrefix(base64Addr, "0:") {
		return "", errors.New("invalid base64 address was given")
	}
	// парсинг адреса
	addr, err := tonutilsgoAddress.ParseAddr(base64Addr)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v:%x", addr.Workchain(), addr.Data()), nil
}
