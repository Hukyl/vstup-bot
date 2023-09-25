package exchangerate

import "fmt"

// ExchangeRate is used to represent an exchange rate from one
// currency to another. From and To are in ISO currency format.
type ExchangeRate struct {
	From string
	To   string
	Rate float64
}

func (e ExchangeRate) String() string {
	return fmt.Sprintf("1 %s - %f %s", e.From, e.Rate, e.To)
}

// Convert converts some amount of initial currency to new currency
// according to the rate.
func (e ExchangeRate) Convert(from float64) float64 {
	return e.Rate * from
}

// Reverse creates a reverse exchange rate, from new currency to initial currency.
func (e ExchangeRate) Reverse() *ExchangeRate {
	return &ExchangeRate{From: e.To, To: e.From, Rate: 1 / e.Rate}
}
