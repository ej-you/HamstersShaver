package services

import (
	"testing"
)


// beauty_jetton_amount.go
func TestBeautyJettonAmountFromInt64(t *testing.T) {
	t.Logf("1 test of getting beauty jetton amount from int64")
	{
		// In
		var int64Amount int64 = 2000_000_000_000
		decimals := 9
		// Out
		beauty := BeautyJettonAmountFromInt64(int64Amount, decimals)
		t.Logf("\tBeauty (amount: %d | decimals: %d): %s", int64Amount, decimals, beauty)
	}

	t.Logf("2 test of getting beauty jetton amount from int64")
	{
		// In
		var int64Amount int64 = 156_203_940_000
		decimals := 9
		// Out
		beauty := BeautyJettonAmountFromInt64(int64Amount, decimals)
		t.Logf("\tBeauty (amount: %d | decimals: %d): %s", int64Amount, decimals, beauty)
	}

	t.Logf("3 test of getting beauty jetton amount from int64")
	{
		// In
		var int64Amount int64 = 3_000_950_000
		decimals := 9
		// Out
		beauty := BeautyJettonAmountFromInt64(int64Amount, decimals)
		t.Logf("\tBeauty (amount: %d | decimals: %d): %s", int64Amount, decimals, beauty)
	}

	t.Logf("4 test of getting beauty jetton amount from int64")
	{
		// In
		var int64Amount int64 = 125_480
		decimals := 6
		// Out
		beauty := BeautyJettonAmountFromInt64(int64Amount, decimals)
		t.Logf("\tBeauty (amount: %d | decimals: %d): %s", int64Amount, decimals, beauty)
	}

	t.Logf("5 test of getting beauty jetton amount from int64")
	{
		// In
		var int64Amount int64 = 30
		decimals := 6
		// Out
		beauty := BeautyJettonAmountFromInt64(int64Amount, decimals)
		t.Logf("\tBeauty (amount: %d | decimals: %d): %s", int64Amount, decimals, beauty)
	}

	t.Logf("6 test of getting beauty jetton amount from int64")
	{
		// In
		var int64Amount int64 = 12_000_030
		decimals := 6
		// Out
		beauty := BeautyJettonAmountFromInt64(int64Amount, decimals)
		t.Logf("\tBeauty (amount: %d | decimals: %d): %s", int64Amount, decimals, beauty)
	}
}


// beauty_jetton_amount.go
func TestBeautyJettonAmountFromFloat64(t *testing.T) {
	t.Logf("1 test of getting beauty jetton amount from float64")
	{
		// In
		float64Amount := 2000.0
		// Out
		beauty := BeautyJettonAmountFromFloat64(float64Amount)
		t.Logf("\tBeauty (amount: %f): %s", float64Amount, beauty)
	}

	t.Logf("2 test of getting beauty jetton amount from float64")
	{
		// In
		float64Amount := 156.20394
		// Out
		beauty := BeautyJettonAmountFromFloat64(float64Amount)
		t.Logf("\tBeauty (amount: %f): %s", float64Amount, beauty)
	}

	t.Logf("3 test of getting beauty jetton amount from float64")
	{
		// In
		float64Amount := 3.00095
		// Out
		beauty := BeautyJettonAmountFromFloat64(float64Amount)
		t.Logf("\tBeauty (amount: %f): %s", float64Amount, beauty)
	}

	t.Logf("4 test of getting beauty jetton amount from float64")
	{
		// In
		float64Amount := 0.12548
		// Out
		beauty := BeautyJettonAmountFromFloat64(float64Amount)
		t.Logf("\tBeauty (amount: %f): %s", float64Amount, beauty)
	}

	t.Logf("5 test of getting beauty jetton amount from float64")
	{
		// In
		float64Amount := 0.00003
		// Out
		beauty := BeautyJettonAmountFromFloat64(float64Amount)
		t.Logf("\tBeauty (amount: %f): %s", float64Amount, beauty)
	}

	t.Logf("6 test of getting beauty jetton amount from float64")
	{
		// In
		float64Amount := 12.00003
		// Out
		beauty := BeautyJettonAmountFromFloat64(float64Amount)
		t.Logf("\tBeauty (amount: %f): %s", float64Amount, beauty)
	}
}
