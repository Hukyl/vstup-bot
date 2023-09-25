package exchangerate_test

import (
	"testing"

	"github.com/Hukyl/vstup-bot/exchangerate"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		initialIso  string
		finalIso    string
		initialRate float64
		reverseRate float64
	}{
		{"first", "second", 5, 0.2},
		{"a", "b", 10, 0.1},
		{"x", "y", 0.4, 2.5},
	}
	for _, test := range testCases {
		r1 := exchangerate.ExchangeRate{test.initialIso, test.finalIso, test.initialRate}
		r2 := exchangerate.ExchangeRate{test.finalIso, test.initialIso, test.reverseRate}
		r1Reverse := *r1.Reverse()
		if r1Reverse != r2 {
			t.Errorf("Rate %f != %f", r2.Rate, r1Reverse.Rate)
		}
	}
}
