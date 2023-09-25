package exchangerate_test

import (
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
