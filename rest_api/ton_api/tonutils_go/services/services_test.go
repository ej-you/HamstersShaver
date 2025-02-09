package services

import (
	"time"

	"testing"
	"github.com/stretchr/testify/assert"
)


// маркеры успеха и неудачи
const (
    successMarker = "\u2713"
    failedMarker  = "\u2717"
)


func logExecTime(t *testing.T, startTime *time.Time) {
	endTime := time.Now()
	t.Logf("\t\tExec time: %v", endTime.Sub(*startTime))
	*startTime = endTime
}


// convert_addresses.go
func TestConvertAddrToBase64(t *testing.T) {
	startTime := time.Now()

	// валидные данные
	t.Logf("Test convert raw addr (hex) to base64 with VALID given parameters")
	{
		// In
		var hexAddress string = "0:2f956143c461769579baef2e32cc2d7bc18283f40d20bb03e432cd603ac33ffc"
		t.Logf("\t\tInput: masterAddress=%q", hexAddress)

		// Out
		base64Address, err := ConvertAddrToBase64(hexAddress)
		if assert.NoErrorf(t, err, "\t%s\tFailed: %v", failedMarker, err) {
			t.Logf("\t%s\tOutput (got base64 addr): %q", successMarker, base64Address)
		}
	}
	logExecTime(t, &startTime)

	// адрес не в том формате
	t.Logf("Test convert addr (given in base64 format) to base64")
	{
		// In
		var hexAddress string = "EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O"
		t.Logf("\t\tInput: masterAddress=%q", hexAddress)

		// Out
		base64Address, err := ConvertAddrToBase64(hexAddress)
		if err == nil { // NOT err
			t.Logf("\t%s\tNo error but was expected one. Value: %q", failedMarker, base64Address)
		} else {
			t.Logf("\t%s\tSuccessfully got error: %v", successMarker, err)
		}
	}
	logExecTime(t, &startTime)

	// кривой адрес
	t.Logf("Test convert raw addr (hex) to base64 with INVALID given parameters")
	{
		// In
		var hexAddress string = "0:16a73dbf1b434ac651b656f8056e06463edf18d6a7b47068fee18c3905f998"
		t.Logf("\t\tInput: masterAddress=%q", hexAddress)

		// Out
		base64Address, err := ConvertAddrToBase64(hexAddress)
		if err == nil { // NOT err
			t.Logf("\t%s\tNo error but was expected one. Value: %q", failedMarker, base64Address)
		} else {
			t.Logf("\t%s\tSuccessfully got error: %v", successMarker, err)
		}
	}
	logExecTime(t, &startTime)
}


// convert_addresses.go
func TestConvertAddrToHEX(t *testing.T) {
	startTime := time.Now()

	// валидные данные
	t.Logf("Test convert addr (base64) to raw hex addr with VALID given parameters")
	{
		// In
		var base64Address string = "EQCvxJy4eG8hyHBFsZ7eePxrRsUQSFE_jpptRAYBmcG_DOGS"
		t.Logf("\t\tInput: masterAddress=%q", base64Address)

		// Out
		hexAddress, err := ConvertAddrToHEX(base64Address)
		if assert.NoErrorf(t, err, "\t%s\tFailed: %v", failedMarker, err) {
			t.Logf("\t%s\tOutput (got hex addr): %q", successMarker, hexAddress)
		}
	}
	logExecTime(t, &startTime)

	// адрес не в том формате
	t.Logf("Test convert addr (given in hex format) to hex")
	{
		// In
		var base64Address string = "0:2f956143c461769579baef2e32cc2d7bc18283f40d20bb03e432cd603ac33ffc"
		t.Logf("\t\tInput: masterAddress=%q", base64Address)

		// Out
		hexAddress, err := ConvertAddrToHEX(base64Address)
		if err == nil { // NOT err
			t.Logf("\t%s\tNo error but was expected one. Value: %q", failedMarker, hexAddress)
		} else {
			t.Logf("\t%s\tSuccessfully got error: %v", successMarker, err)
		}
	}
	logExecTime(t, &startTime)

	// кривой адрес
	t.Logf("Test convert addr (base64) to raw hex addr with INVALID given parameters")
	{
		// In
		var base64Address string = "EQCvxJy4eG8hyHBFsZ7eePxrRsUQSFE_jpptRAYBmcG"
		t.Logf("\t\tInput: masterAddress=%q", base64Address)

		// Out
		hexAddress, err := ConvertAddrToHEX(base64Address)
		if err == nil { // NOT err
			t.Logf("\t%s\tNo error but was expected one. Value: %q", failedMarker, hexAddress)
		} else {
			t.Logf("\t%s\tSuccessfully got error: %v", successMarker, err)
		}
	}
	logExecTime(t, &startTime)
}
