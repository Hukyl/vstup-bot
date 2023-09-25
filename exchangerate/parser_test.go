package exchangerate_test

import (
	"slices"
	"testing"

	"github.com/Hukyl/vstup-bot/exchangerate"
)

func notIfFalse(thing bool) string {
	if thing {
		return ""
	}
	return " not"
}

func TestCurrencyExists(t *testing.T) {
	testCases := map[string]bool{
		"USD":  true,
		"UAH":  true,
		"ABC":  false,
		"ABBC": true,
		"XYD":  false,
	}
	for iso, want := range testCases {
		if exchangerate.CurrencyExists(iso) != want {
			t.Errorf("%s does%s exist", iso, notIfFalse(want))
		}
	}
}

func TestGetAvailableCurrencies(t *testing.T) {
	testCases := map[string]bool{
		"USD":  true,
		"UAH":  true,
		"ABC":  false,
		"ABBC": true,
		"XYD":  false,
	}
	availableCurrencies := exchangerate.GetAvailableCurrencies()
	for iso, contains := range testCases {
		if slices.Contains(availableCurrencies, iso) != contains {
			t.Errorf("%s does%s exist", iso, notIfFalse(contains))
		}
	}
}
